package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/timo-reymann/git-semver-tag/pkg/cli"
	"github.com/timo-reymann/git-semver-tag/pkg/git"
	"github.com/timo-reymann/git-semver-tag/pkg/version"
)

func getCurrentTag(tags []string) version.Version {
	currentTag := version.Version{}
	if tags == nil || len(tags) == 0 {
		currentTag = version.Empty()
	} else {
		currentTag = version.ProcessVersionTag(tags[len(tags)-1])
	}
	return currentTag
}

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
	args := cmd.Args()

	// Set current tag
	currentTag := getCurrentTag(tags)

	if args.Interactive != nil && *args.Interactive {
		err = runInteractiveAssistant(args)
		if err != nil {
			cmd.HandleError("Aborted interactive assistant")
		}
	}

	// Check suffix
	suffix := args.Suffix
	if suffix != nil {
		currentTag.Suffix = *suffix
	}

	// Increment version and print it
	if err := currentTag.IncrementBasedOnLevel(*args.Level); err != nil {
		cmd.HandleError("Unsupported level")
	}
	println(currentTag.String())

	// Create new tag
	if err := git.CreateTag(currentTag.String(), args.Message); err != nil {
		cmd.HandleError("Error creating tag: " + err.Error())
	}

	// Push tag
	push := args.Push
	if push != nil && *push {
		pushTag(currentTag, cmd)
	}

}

func selectOption(label string, options []string) (string, error) {
	prompt := promptui.Select{
		Label:        label,
		Items:        options,
		HideSelected: true,
	}

	_, result, err := prompt.Run()
	return result, err
}

func runInteractiveAssistant(args *cli.SemverGitTagCliArgs) error {
	level, err := selectOption("Select level", []string{
		"patch",
		"minor",
		"major",
	})
	if err != nil {
		return err
	}
	args.Level = &level

	push, err := selectOption("Push now?", []string{
		"yes",
		"no",
	})
	if err != nil {
		return err
	}
	if push == "yes" {
		pushYes := true
		args.Push = &pushYes
	}

	return nil
}

func pushTag(currentTag version.Version, cmd cli.SemverGitTagCliStd) {
	stdout, err := git.PushTag(currentTag.String())
	println(stdout)
	if err != nil {
		cmd.HandleError("Failed to tag: " + err.Error())
	}
}
