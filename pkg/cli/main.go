package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// SemverGitTagCliArgs represents the comand line arguments
type SemverGitTagCliArgs struct {
	Level       *string
	Suffix      *string
	Message     *string
	Push        *bool
	Interactive *bool
}

// SemverGitTagCli is the interface for a semver git tag cli implementation
type SemverGitTagCli interface {
	HandleError(msg string)
	ParseArgs()
	Args() *SemverGitTagCliArgs
}

// SemverGitTagCliStd implements SemverGitTagCli using stdout
type SemverGitTagCliStd struct {
	args *SemverGitTagCliArgs
}

// HandleError prints the error messages and terminates the process
func (cli *SemverGitTagCliStd) HandleError(msg string) {
	_, _ = fmt.Fprintf(os.Stderr, "Error:  %s\n", msg)
	os.Exit(1)
}

// ParseArgs declares and parses the command line args using flag and validates them,
// if a invalid parameter was passed HandleError is called
func (cli *SemverGitTagCliStd) ParseArgs() {
	cli.args = &SemverGitTagCliArgs{
		Level:   flag.String("level", "", "major|minor|patch|suffix-only - Version to increase"),
		Suffix:  flag.String("suffix", "", "Suffix to append to the version (optional)"),
		Message: flag.String("message", "", "Message for tag (optional)"),
		Push:    flag.Bool("push", false, "Push git tag (if origin is set)"),
	}
	flag.Parse()
	cli.verifyArgIsPresent("Level", cli.args.Level)
}

// Args returns a ref to the cli's current arguments
func (cli *SemverGitTagCliStd) Args() *SemverGitTagCliArgs {
	return cli.args
}

func (cli *SemverGitTagCliStd) verifyArgIsPresent(name string, val *string) {
	if val == nil || strings.TrimSpace(*val) == "" {
		flag.PrintDefaults()
		println("")
		cli.HandleError(fmt.Sprintf("Required parameter '%s' is not present", strings.ToLower(name)))
	}
}
