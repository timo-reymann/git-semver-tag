package cmd

import (
	"github.com/timo-reymann/git-semver-tag/pkg/cli"
	"github.com/timo-reymann/git-semver-tag/pkg/git"
	"github.com/timo-reymann/git-semver-tag/pkg/version"
)

// Execute the command line
func Execute() {
	cmd := cli.SemverGitTagCliStd{}

	// Check if git is installed first
	if !git.IsGitInstalled() {
		cmd.HandleError("Git is not installed, but required")
	}

	// Try to get tags for current folder
	tags, err := git.GetTags()
	if err != nil {
		cmd.HandleError("Failed to get tags: " + err.Error())
	}

	cmd.ParseArgs()

	// Set current tag
	currentTag := version.Version{}
	if tags == nil || len(tags) == 0 {
		currentTag = version.Empty()
	} else {
		currentTag = version.ProcessVersionTag(tags[len(tags)-1])
	}

	// Check suffix
	suffix := cmd.Args().Suffix
	if suffix != nil {
		currentTag.Suffix = *suffix
	}

	// Increment version and print it
	if err := currentTag.IncrementBasedOnLevel(*cmd.Args().Level); err != nil {
		cmd.HandleError("Unsupported level")
	}
	println(currentTag.String())

	// Create new tag
	if err := git.CreateTag(currentTag.String(), cmd.Args().Message); err != nil {
		cmd.HandleError("Error creating tag: " + err.Error())
	}

	// Push tag
	push := cmd.Args().Push
	if push != nil && *push {
		stdout, err := git.PushTag(currentTag.String())
		println(stdout)
		if err != nil {
			cmd.HandleError("Failed to tag: " + err.Error())
		}
	}

}
