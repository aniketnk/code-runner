package cpp

import (
	"path/filepath"

	"github.com/aniketnk/code-runner/cmd"
	"github.com/aniketnk/code-runner/util"
)

// Run : Specific function to compile and run code
func Run(files []string, testCases []string, timeout string) ([]string, []string, error) {
	workDir := filepath.Dir(files[0])
	binName := "a.out"

	sourceFiles := util.FilterByExtension(files, "cpp")
	args := append([]string{"g++", "-std=c++11", "-o", binName}, sourceFiles...)

	stdout, stderr, err := cmd.Run(workDir, args...)
	if err != nil {
		return []string{stdout}, []string{stderr}, err
	}

	binPath := filepath.Join(workDir, binName)
	stdoutList, stderrList := make([]string, len(testCases)), make([]string, len(testCases))
	attempts := 0

	for i := range testCases {
		stdoutList[i], stderrList[i], err = cmd.RunStdin(workDir, testCases[i], "timeout", timeout, binPath)
		attempts++
		if err != nil {
			break
		}
	}
	return stdoutList[:attempts], stderrList[:attempts], err
}
