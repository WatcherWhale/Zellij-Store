package shell

import (
	"os"
	"strings"
)

func DetectShell() string {
	envVars := os.Environ()

	for _, eVar := range envVars {
		key := strings.ToLower(strings.Split(eVar, "=")[0])

		if strings.Contains(key, "fish_pid") {
			return "fish"
		}

		if strings.Contains(key, "bash_version") {
			return "bash"
		}
	}

	return "fish"
}
