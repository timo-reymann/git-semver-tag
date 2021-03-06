package git

import "strings"

// GetTags returns all tags for the current git repository sorted by refname
func GetTags() (tags []string, err error) {
	stdout, err := ExecGitCommand("tag", "--list", "--sort=v:refname")

	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(stdout), "\n"), nil
}

// CreateTag creates a new tag
func CreateTag(tag string, message *string) (err error) {
	args := []string{
		"tag",
		tag,
	}

	if message != nil {
		args = append(args, "-m")
		args = append(args, *message)
	}

	_, err = ExecGitCommand(args...)
	return err
}

// PushTag pushes the specified tag to the origin 'origin' that might fail if the origin is not defined.
func PushTag(tag string) (stdout string, err error) {
	return ExecGitCommand("push", "origin", "refs/tags/"+tag)
}
