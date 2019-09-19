package cmd

import (
	"github.com/timo-reymann/git-semver-tag/pkg/cli"
	"github.com/timo-reymann/git-semver-tag/pkg/git"
	"github.com/timo-reymann/git-semver-tag/pkg/version"
)

func Execute() {
	cmd := cli.SemverGitTagCliStd{}

	if !git.IsGitInstalled() {
		cmd.HandleError("Git is not installed, but required")
	}

	cmd.ParseArgs()

	tags, err := git.GetTags()
	if err != nil {
		cmd.HandleError("Failed to get tags: " + err.Error())
	}

	currentTag := version.Version{}

	if tags == nil || len(tags) == 0 {
		currentTag = version.Version{
			Prefix: "",
			Major:  0,
			Minor:  0,
			Patch:  0,
			Suffix: "",
		}
	} else {
		currentTag = version.ProcessVersionTag(tags[len(tags)-1])
	}

	suffix := cmd.Args().Suffix
	if suffix != nil {
		currentTag.Suffix = *suffix
	}

	switch *cmd.Args().Level {
	case "major":
		currentTag.NextMajor()

	case "minor":
		currentTag.NextMinor()

	case "patch":
		currentTag.NextPatch()

	case "suffix-only":
		// do nothing with version itself

	default:
		cmd.HandleError("Unsupported level")
	}

	print(currentTag.String())

	if err := git.CreateTag(currentTag.String(), cmd.Args().Message); err != nil {
		cmd.HandleError("Error creating tag: " + err.Error())
	}

}
