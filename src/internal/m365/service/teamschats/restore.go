package teamschats

import (
	"context"
	"errors"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ConsumeRestoreCollections will restore the specified data collections
func (h *teamsChatsHandler) ConsumeRestoreCollections(
	ctx context.Context,
	rcc inject.RestoreConsumerConfig,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.Details, *data.CollectionStats, error) {
	if len(dcs) == 0 {
		return nil, nil, clues.WrapWC(ctx, data.ErrNoData, "performing restore")
	}

	// TODO(ashmrtn): We should stop relying on the context for rate limiter stuff
	// and instead configure this when we make the handler instance. We can't
	// initialize it in the NewHandler call right now because those functions
	// aren't (and shouldn't be) returning a context along with the handler. Since
	// that call isn't directly calling into this function even if we did
	// initialize the rate limiter there it would be lost because it wouldn't get
	// stored in an ancestor of the context passed to this function.
	ctx = graph.BindRateLimiterConfig(
		ctx,
		graph.LimiterCfg{Service: path.TeamsChatsService})

	var (
		deets          = &details.Builder{}
		restoreMetrics support.CollectionMetrics
		el             = errs.Local()
	)

	// Reorder collections so that the parents directories are created
	// before the child directories; a requirement for permissions.
	data.SortRestoreCollections(dcs)

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		if el.Failure() != nil {
			break
		}

		var (
			err      error
			category = dc.FullPath().Category()
			metrics  support.CollectionMetrics
			ictx     = clues.Add(ctx,
				"category", category,
				"restore_location", clues.Hide(rcc.RestoreConfig.Location),
				"protected_resource", clues.Hide(dc.FullPath().ProtectedResource()),
				"full_path", dc.FullPath())
		)

		switch dc.FullPath().Category() {
		case path.ChatsCategory:
			// chats cannot be restored using Graph API.
			// a delegated token is required, and Corso has no
			// good way of obtaining such a token.
			logger.Ctx(ictx).Debug("Skipping restore for channel messages")
		default:
			return nil, nil, clues.NewWC(ictx, "data category not supported").
				With("category", category)
		}

		restoreMetrics = support.CombineMetrics(restoreMetrics, metrics)

		if err != nil {
			el.AddRecoverable(ictx, err)
		}

		if errors.Is(err, context.Canceled) {
			break
		}
	}

	status := support.CreateStatus(
		ctx,
		support.Restore,
		len(dcs),
		restoreMetrics,
		rcc.RestoreConfig.Location)

	return deets.Details(), status.ToCollectionStats(), el.Failure()
}
