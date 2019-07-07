package language

import (
	"github.com/aniketnk/code-runner/language/c"
	"github.com/aniketnk/code-runner/language/cpp"
	"github.com/aniketnk/code-runner/language/java"
	"github.com/aniketnk/code-runner/language/python2"
	"github.com/aniketnk/code-runner/language/python3"
)

type runFn func([]string, []string, string) ([]string, []string, error)

var languages = map[string]runFn{
	"c":       c.Run,
	"cpp":     cpp.Run,
	"java":    java.Run,
	"python":  python3.Run,
	"python2": python2.Run,
	"python3": python3.Run,
}

// IsSupported : checks if the given language is supported
func IsSupported(lang string) bool {
	_, supported := languages[lang]
	return supported
}

// Run : chooses the correct function from languages map
func Run(lang string, files []string, testCases []string, timeout string) ([]string, []string, error) {
	return languages[lang](files, testCases, timeout)
}
