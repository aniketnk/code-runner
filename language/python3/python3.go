package python3

import (
	"path/filepath"

	"github.com/aniketnk/code-runner/cmd"
)

// Run : Specific function to compile and run code
func Run(files []string, testCases []string, timeout string) ([]string, []string, error) {
	workDir := filepath.Dir(files[0])
	stdoutList, stderrList := make([]string, len(testCases)), make([]string, len(testCases))
	attempts := 0
	var err error

	for i := range testCases {
		stdoutList[i], stderrList[i], err = cmd.RunStdin(workDir, testCases[i], "timeout", timeout, "python3", files[0])
		attempts++
		if err != nil {
			break
		}
	}
	return stdoutList[:attempts], stderrList[:attempts], err
}
