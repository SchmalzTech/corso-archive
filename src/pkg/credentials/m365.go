package credentials

import (
	"os"

	"github.com/alcionai/clues"
)

// envvar consts
const (
	AzureClientID     = "AZURE_CLIENT_ID"
	AzureClientSecret = "AZURE_CLIENT_SECRET"
	AzureUsername     = "AZURE_USERNAME"
	AzureUserPassword = "AZURE_USER_PASSWORD"
)

// M365 aggregates m365 credentials from flag and env_var values.
type M365 struct {
	AzureClientID     string
	AzureClientSecret string
	AzureUsername     string
	AzureUserPassword string
}

// M365 is a helper for aggregating m365 secrets and credentials.
func GetM365() M365 {
	// check env and overide is flags found
	// var AzureClientID, AzureClientSecret string
	AzureClientID := os.Getenv(AzureClientID)
	AzureClientSecret := os.Getenv(AzureClientSecret)
	AzureUserPassword := os.Getenv(AzureUserPassword)
	AzureUsername := os.Getenv(AzureUsername)

	return M365{
		AzureClientID:     AzureClientID,
		AzureClientSecret: AzureClientSecret,
		AzureUsername:     AzureUsername,
		AzureUserPassword: AzureUserPassword,
	}
}

func (c M365) Validate() error {
	check := map[string]string{
		AzureClientID:     c.AzureClientID,
		AzureClientSecret: c.AzureClientSecret,
	}

	for k, v := range check {
		if len(v) == 0 {
			return clues.Stack(errMissingRequired, clues.New(k))
		}
	}

	return nil
}
