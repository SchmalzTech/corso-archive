package main

import (
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var (
	emailsCmd = &cobra.Command{
		Use:   "emails",
		Short: "Generate exchange emails",
		RunE:  handleExchangeEmailFactory,
	}

	eventsCmd = &cobra.Command{
		Use:   "events",
		Short: "Generate exchange calendar events",
		RunE:  handleExchangeCalendarEventFactory,
	}

	contactsCmd = &cobra.Command{
		Use:   "contacts",
		Short: "Generate exchange contacts",
		RunE:  handleExchangeContactFactory,
	}
)

func addExchangeCommands(parent *cobra.Command) {
	parent.AddCommand(emailsCmd)
	parent.AddCommand(eventsCmd)
	parent.AddCommand(contactsCmd)
}

func handleExchangeEmailFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.ExchangeService
		category = path.EmailCategory
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, tenantID, err := getGCAndVerifyUser(ctx, user)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		gc,
		service,
		category,
		selectors.NewExchangeRestore().Selector,
		tenantID, user, destination,
		count,
		func(id, now, subject, body string) []byte {
			return mockconnector.GetMockMessageWith(
				user, user, user,
				subject, body,
				now, now, now, now)
		},
	)
	if err != nil {
		return Only(ctx, err)
	}

	deets.PrintEntries(ctx)

	return nil
}

func handleExchangeCalendarEventFactory(cmd *cobra.Command, args []string) error {
	var (
		ctx      = cmd.Context()
		service  = path.ExchangeService
		category = path.EventsCategory
	)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, tenantID, err := getGCAndVerifyUser(ctx, user)
	if err != nil {
		return Only(ctx, err)
	}

	deets, err := generateAndRestoreItems(
		ctx,
		gc,
		service,
		category,
		selectors.NewExchangeRestore().Selector,
		tenantID, user, destination,
		count,
		func(id, now, subject, body string) []byte {
			return mockconnector.GetMockEventWith(
				user, subject, body, body,
				now, now)
		},
	)
	if err != nil {
		return Only(ctx, err)
	}

	deets.PrintEntries(ctx)

	return nil
}

func handleExchangeContactFactory(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), ErrNotYetImplemeted)

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	// generate mocked contacts

	return nil
}
