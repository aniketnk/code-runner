package java

import (
	"path/filepath"

	"github.com/aniketnk/code-runner/cmd"
)

// Run : Specific function to compile and run code
func Run(files []string, testCases []string, timeout string) ([]string, []string, error) {
	workDir := filepath.Dir(files[0])
	fname := filepath.Base(files[0])

	stdout, stderr, err := cmd.Run(workDir, "javac", fname)
	if err != nil {
		return []string{stdout}, []string{stderr}, err
	}

	stdoutList, stderrList := make([]string, len(testCases)), make([]string, len(testCases))

	for i := range testCases {
		stdoutList[i], stderrList[i], err = cmd.RunStdin(workDir, testCases[i], "timeout", timeout, "java", className(fname))
		if err != nil {
			break
		}
	}
	return stdoutList, stderrList, err
}

func className(fname string) string {
	ext := filepath.Ext(fname)
	return fname[0 : len(fname)-len(ext)]
}
