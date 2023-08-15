package kopia

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/logger"
)

// TODO(ashmrtn): Move this into some inject package. Here to avoid import
// cycles.
type BackupBases interface {
	RemoveMergeBaseByManifestID(manifestID manifest.ID)
	Backups() []BackupEntry
	AssistBackups() []BackupEntry
	MinBackupVersion() int
	MergeBases() []ManifestEntry
	ClearMergeBases()
	AssistBases() []ManifestEntry
	ClearAssistBases()
	MergeBackupBases(
		ctx context.Context,
		other BackupBases,
		reasonToKey func(identity.Reasoner) string,
	) BackupBases
}

type backupBases struct {
	// backups and mergeBases should be modified together as they relate similar
	// data.
	backups       []BackupEntry
	mergeBases    []ManifestEntry
	assistBackups []BackupEntry
	assistBases   []ManifestEntry
}

func (bb *backupBases) RemoveMergeBaseByManifestID(manifestID manifest.ID) {
	idx := slices.IndexFunc(
		bb.mergeBases,
		func(man ManifestEntry) bool {
			return man.ID == manifestID
		})
	if idx >= 0 {
		bb.mergeBases = slices.Delete(bb.mergeBases, idx, idx+1)
	}

	// TODO(ashmrtn): This may not be strictly necessary but is at least easier to
	// reason about.
	idx = slices.IndexFunc(
		bb.assistBases,
		func(man ManifestEntry) bool {
			return man.ID == manifestID
		})
	if idx >= 0 {
		bb.assistBases = slices.Delete(bb.assistBases, idx, idx+1)
	}

	idx = slices.IndexFunc(
		bb.backups,
		func(bup BackupEntry) bool {
			return bup.SnapshotID == string(manifestID)
		})
	if idx >= 0 {
		bb.backups = slices.Delete(bb.backups, idx, idx+1)
	}
}

func (bb backupBases) Backups() []BackupEntry {
	return slices.Clone(bb.backups)
}

func (bb backupBases) AssistBackups() []BackupEntry {
	return slices.Clone(bb.assistBackups)
}

func (bb *backupBases) MinBackupVersion() int {
	min := version.NoBackup

	if bb == nil {
		return min
	}

	for _, bup := range bb.backups {
		if min == version.NoBackup || bup.Version < min {
			min = bup.Version
		}
	}

	return min
}

func (bb backupBases) MergeBases() []ManifestEntry {
	return slices.Clone(bb.mergeBases)
}

func (bb *backupBases) ClearMergeBases() {
	bb.mergeBases = nil
	bb.backups = nil
}

func (bb backupBases) AssistBases() []ManifestEntry {
	return slices.Clone(bb.assistBases)
}

func (bb *backupBases) ClearAssistBases() {
	bb.assistBases = nil
}

// BaseKeyServiceCategory makes a backup base key using
// the reasoner's Service and Category.
func BaseKeyServiceCategory(br identity.Reasoner) string {
	return br.Service().String() + br.Category().String()
}

// MergeBackupBases reduces the two BackupBases into a single BackupBase.
// Assumes the passed in BackupBases represents a prior backup version (across
// some migration that disrupts lookup), and that the BackupBases used to call
// this function contains the current version.
//
// reasonToKey should be a function that, given a Reasoner, will produce some
// string that represents Reasoner in the context of the merge operation. For
// example, to merge BackupBases across a ProtectedResource migration, the
// Reasoner's service and category can be used as the key.
//
// Selection priority, for each reason key generated by reasonsToKey, follows
// these rules:
//  1. If the called BackupBases has an entry for a given reason, ignore the
//     other BackupBases matching that reason.
//  2. If the called BackupBases has only AssistBases, look for a matching
//     MergeBase manifest in the other BackupBases.
//  3. If the called BackupBases has no entry for a reason, look for a matching
//     MergeBase in the other BackupBases.
func (bb *backupBases) MergeBackupBases(
	ctx context.Context,
	other BackupBases,
	reasonToKey func(reason identity.Reasoner) string,
) BackupBases {
	if other == nil || (len(other.MergeBases()) == 0 && len(other.AssistBases()) == 0) {
		return bb
	}

	if bb == nil || (len(bb.MergeBases()) == 0 && len(bb.AssistBases()) == 0) {
		return other
	}

	toMerge := map[string]struct{}{}
	assist := map[string]struct{}{}

	// Track the bases in bb.
	for _, m := range bb.mergeBases {
		for _, r := range m.Reasons {
			k := reasonToKey(r)

			toMerge[k] = struct{}{}
			assist[k] = struct{}{}
		}
	}

	for _, m := range bb.assistBases {
		for _, r := range m.Reasons {
			k := reasonToKey(r)
			assist[k] = struct{}{}
		}
	}

	var toAdd []ManifestEntry

	// Calculate the set of mergeBases to pull from other into this one.
	for _, m := range other.MergeBases() {
		useReasons := []identity.Reasoner{}

		for _, r := range m.Reasons {
			k := reasonToKey(r)
			if _, ok := toMerge[k]; ok {
				// Assume other contains prior manifest versions.
				// We don't want to stack a prior version incomplete onto
				// a current version's complete snapshot.
				continue
			}

			useReasons = append(useReasons, r)
		}

		if len(useReasons) > 0 {
			m.Reasons = useReasons
			toAdd = append(toAdd, m)
		}
	}

	res := &backupBases{
		backups:     bb.Backups(),
		mergeBases:  bb.MergeBases(),
		assistBases: bb.AssistBases(),
		// Note that assistBackups are a new feature and don't exist
		// in prior versions where we were using UPN based reasons i.e.
		// other won't have any assistBackups.
		assistBackups: bb.AssistBackups(),
	}

	// Add new mergeBases and backups.
	for _, man := range toAdd {
		// Will get empty string if not found which is fine, it'll fail one of the
		// other checks.
		bID, _ := man.GetTag(TagBackupID)

		bup, ok := getBackupByID(other.Backups(), bID)
		if !ok {
			logger.Ctx(ctx).Infow(
				"not unioning snapshot missing backup",
				"other_manifest_id", man.ID,
				"other_backup_id", bID)

			continue
		}

		bup.Reasons = man.Reasons

		res.backups = append(res.backups, bup)
		res.mergeBases = append(res.mergeBases, man)
		// TODO(pandeyabs): Remove this once we remove overlap between
		// between merge and assist bases as part of #3943.
		res.assistBases = append(res.assistBases, man)
	}

	return res
}

func findNonUniqueManifests(
	ctx context.Context,
	manifests []ManifestEntry,
) map[manifest.ID]struct{} {
	// ReasonKey -> manifests with that reason.
	reasons := map[string][]ManifestEntry{}
	toDrop := map[manifest.ID]struct{}{}

	for _, man := range manifests {
		// Incomplete snapshots are used only for kopia-assisted incrementals. The
		// fact that we need this check here makes it seem like this should live in
		// the kopia code. However, keeping it here allows for better debugging as
		// the kopia code only has access to a path builder which means it cannot
		// remove the resource owner from the error/log output. That is also below
		// the point where we decide if we should do a full backup or an incremental.
		if len(man.IncompleteReason) > 0 {
			logger.Ctx(ctx).Infow(
				"dropping incomplete manifest",
				"manifest_id", man.ID)

			toDrop[man.ID] = struct{}{}

			continue
		}

		for _, reason := range man.Reasons {
			mapKey := reasonKey(reason)
			reasons[mapKey] = append(reasons[mapKey], man)
		}
	}

	for reason, mans := range reasons {
		ictx := clues.Add(ctx, "reason", reason)

		if len(mans) == 0 {
			// Not sure how this would happen but just in case...
			continue
		} else if len(mans) > 1 {
			mIDs := make([]manifest.ID, 0, len(mans))
			for _, m := range mans {
				toDrop[m.ID] = struct{}{}
				mIDs = append(mIDs, m.ID)
			}

			// TODO(ashmrtn): We should actually just remove this reason from the
			// manifests and then if they have no reasons remaining drop them from the
			// set.
			logger.Ctx(ictx).Infow(
				"dropping manifests with duplicate reason",
				"manifest_ids", mIDs)

			continue
		}
	}

	return toDrop
}

func getBackupByID(backups []BackupEntry, bID string) (BackupEntry, bool) {
	if len(bID) == 0 {
		return BackupEntry{}, false
	}

	idx := slices.IndexFunc(backups, func(b BackupEntry) bool {
		return string(b.ID) == bID
	})

	if idx < 0 || idx >= len(backups) {
		return BackupEntry{}, false
	}

	return backups[idx], true
}

// fixupAndVerify goes through the set of backups and snapshots used for merging
// and ensures:
//   - the reasons for selecting merge snapshots are distinct
//   - all bases used for merging have a backup model with item and details
//     snapshot ID
//
// Backups that have overlapping reasons or that are not complete are removed
// from the set. Dropping these is safe because it only affects how much data we
// pull. On the other hand, *not* dropping them is unsafe as it will muck up
// merging when we add stuff to kopia (possibly multiple entries for the same
// item etc).
//
// TODO(pandeyabs): Refactor common code into a helper as part of #3943.
func (bb *backupBases) fixupAndVerify(ctx context.Context) {
	toDrop := findNonUniqueManifests(ctx, bb.mergeBases)

	var (
		backupsToKeep       []BackupEntry
		assistBackupsToKeep []BackupEntry
		mergeToKeep         []ManifestEntry
		assistToKeep        []ManifestEntry
	)

	for _, man := range bb.mergeBases {
		if _, ok := toDrop[man.ID]; ok {
			continue
		}

		bID, _ := man.GetTag(TagBackupID)

		bup, ok := getBackupByID(bb.backups, bID)
		if !ok {
			toDrop[man.ID] = struct{}{}

			logger.Ctx(ctx).Info(
				"dropping merge base due to missing backup",
				"manifest_id", man.ID)

			continue
		}

		deetsID := bup.StreamStoreID
		if len(deetsID) == 0 {
			deetsID = bup.DetailsID
		}

		if len(bup.SnapshotID) == 0 || len(deetsID) == 0 {
			toDrop[man.ID] = struct{}{}

			logger.Ctx(ctx).Info(
				"dropping merge base due to invalid backup",
				"manifest_id", man.ID)

			continue
		}

		backupsToKeep = append(backupsToKeep, bup)
		mergeToKeep = append(mergeToKeep, man)
	}

	// Every merge base is also a kopia assist base.
	// TODO(pandeyabs): This should be removed as part of #3943.
	for _, man := range bb.mergeBases {
		if _, ok := toDrop[man.ID]; ok {
			continue
		}

		assistToKeep = append(assistToKeep, man)
	}

	// Drop assist snapshots with overlapping reasons.
	toDropAssists := findNonUniqueManifests(ctx, bb.assistBases)

	for _, man := range bb.assistBases {
		if _, ok := toDropAssists[man.ID]; ok {
			continue
		}

		bID, _ := man.GetTag(TagBackupID)

		bup, ok := getBackupByID(bb.assistBackups, bID)
		if !ok {
			toDrop[man.ID] = struct{}{}

			logger.Ctx(ctx).Info(
				"dropping assist base due to missing backup",
				"manifest_id", man.ID)

			continue
		}

		deetsID := bup.StreamStoreID
		if len(deetsID) == 0 {
			deetsID = bup.DetailsID
		}

		if len(bup.SnapshotID) == 0 || len(deetsID) == 0 {
			toDrop[man.ID] = struct{}{}

			logger.Ctx(ctx).Info(
				"dropping assist base due to invalid backup",
				"manifest_id", man.ID)

			continue
		}

		assistBackupsToKeep = append(assistBackupsToKeep, bup)
		assistToKeep = append(assistToKeep, man)
	}

	bb.backups = backupsToKeep
	bb.mergeBases = mergeToKeep
	bb.assistBases = assistToKeep
	bb.assistBackups = assistBackupsToKeep
}
