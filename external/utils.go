package external

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
)

func userHomeDir() string {
	// get user home directory for windows: %USERPROFILE%
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}

	// get user home directory for *nix: $HOME
	return os.Getenv("HOME")
}

func loadJSONFile(path string, p interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(c, p)
	if err != nil {
		return err
	}

	return nil
}
