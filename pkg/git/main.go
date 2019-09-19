package git

import (
	"os/exec"
)

// ExecGitCommand runs a git command with the given args returning stdout, and if a error occurs the error
func ExecGitCommand(args ...string) (stdout string, err error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.CombinedOutput()

	if out != nil {
		stdout = string(out)
	}

	return stdout, err
}

// IsGitInstalled returns if the user has git installed or not by trying to getting the git version
func IsGitInstalled() bool {
	_, err := ExecGitCommand("--version")
	return err == nil
}
