package shell

import (
	"fmt"
)

type exportTransformer func(string, string) string

var transformers map[string]exportTransformer = map[string]exportTransformer{
	"fish": fishTransformer,
	"bash": bashTransformer,
}

func getTransformer(shell string) exportTransformer {
	transformer, ok := transformers[shell]

	if !ok {
		return transformers["bash"]
	}
	return transformer
}

func GenerateScript(shell string, envVars map[string]string) string {
	var script string
	transformer := getTransformer(shell)

	for key, val := range envVars {
		script += transformer(key, val) + "\n"
	}

	return script
}

func fishTransformer(key, value string) string {
	return fmt.Sprintf("set -x %s %s", key, value)
}

func bashTransformer(key, value string) string {
	return fmt.Sprintf("export %s=%s", key, value)
}
