package driver

import "os"

const ACC_ENV_KEY = "USDKACC"
const ACC_SKIP_REASON = "skip test for non-acc environment"

// check test env is acceptance testing or not
func IsAcc() bool {
	return len(os.Getenv(ACC_ENV_KEY)) > 0
}
