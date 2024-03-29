package defaults

import (
	"github.com/shztki/nifcloud-sdk-go/internal/shareddefaults"
)

// SharedCredentialsFilename returns the SDK's default file path
// for the shared credentials file.
//
// Builds the shared config file path based on the OS's platform.
//
//   - Linux/Unix: $HOME/.nifcloud/credentials
//   - Windows: %USERPROFILE%\.aws\credentials
func SharedCredentialsFilename() string {
	return shareddefaults.SharedCredentialsFilename()
}

// SharedConfigFilename returns the SDK's default file path for
// the shared config file.
//
// Builds the shared config file path based on the OS's platform.
//
//   - Linux/Unix: $HOME/.nifcloud/config
//   - Windows: %USERPROFILE%\.aws\config
func SharedConfigFilename() string {
	return shareddefaults.SharedConfigFilename()
}
