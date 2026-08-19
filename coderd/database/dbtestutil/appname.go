package dbtestutil

import (
	"github.com/coder/coder/v2/codersdk"
)

// AllFamilyApps returns the app names of every attributed family, sorted,
// for tests that construct the session count query parameters.
func AllFamilyApps() (vscode, ssh, jetbrains, reconnectingPty []string) {
	return codersdk.AppNamesInFamily(codersdk.AppFamilyVSCode),
		codersdk.AppNamesInFamily(codersdk.AppFamilySSH),
		codersdk.AppNamesInFamily(codersdk.AppFamilyJetBrains),
		codersdk.AppNamesInFamily(codersdk.AppFamilyReconnectingPTY)
}
