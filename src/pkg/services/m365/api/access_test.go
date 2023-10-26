package api

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/tester/tsetup"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
)

type AccessAPIIntgSuite struct {
	tester.Suite
	its tsetup.M365
}

func TestAccessAPIIntgSuite(t *testing.T) {
	suite.Run(t, &AccessAPIIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *AccessAPIIntgSuite) SetupSuite() {
	suite.its = tsetup.NewM365IntegrationTester(suite.T())
}

func (suite *AccessAPIIntgSuite) TestGetToken() {
	tests := []struct {
		name      string
		creds     func() account.M365Config
		expectErr require.ErrorAssertionFunc
	}{
		{
			name:      "good",
			creds:     func() account.M365Config { return suite.its.AC.Credentials },
			expectErr: require.NoError,
		},
		{
			name: "bad tenant ID",
			creds: func() account.M365Config {
				creds := suite.its.AC.Credentials
				creds.AzureTenantID = "ZIM"

				return creds
			},
			expectErr: require.Error,
		},
		{
			name: "missing tenant ID",
			creds: func() account.M365Config {
				creds := suite.its.AC.Credentials
				creds.AzureTenantID = ""

				return creds
			},
			expectErr: require.Error,
		},
		{
			name: "bad client ID",
			creds: func() account.M365Config {
				creds := suite.its.AC.Credentials
				creds.AzureClientID = "GIR"

				return creds
			},
			expectErr: require.Error,
		},
		{
			name: "missing client ID",
			creds: func() account.M365Config {
				creds := suite.its.AC.Credentials
				creds.AzureClientID = ""

				return creds
			},
			expectErr: require.Error,
		},
		{
			name: "bad client secret",
			creds: func() account.M365Config {
				creds := suite.its.AC.Credentials
				creds.AzureClientSecret = "MY TALLEST"

				return creds
			},
			expectErr: require.Error,
		},
		{
			name: "missing client secret",
			creds: func() account.M365Config {
				creds := suite.its.AC.Credentials
				creds.AzureClientSecret = ""

				return creds
			},
			expectErr: require.Error,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ac, err := NewClient(suite.its.AC.Credentials, control.DefaultOptions())
			require.NoError(t, err, clues.ToCore(err))

			ac.Credentials = test.creds()

			err = ac.Access().GetToken(ctx)
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
