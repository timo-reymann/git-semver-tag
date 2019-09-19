package version

import "fmt"

type Version struct {
	Prefix string
	Major  int
	Minor  int
	Patch  int
	Suffix string
}

// NextMajor increases the major version and resets everything except the prefix
func (v *Version) NextMajor() {
	v.Major++
	v.Minor = 0
	v.Patch = 0
	v.Suffix = ""
}

// NextMinor increases the minor version and resets patch and Suffix
func (v *Version) NextMinor() {
	v.Minor++
	v.Patch = 0
	v.Suffix = ""
}

// NextPatch increases the patch version and rests the Suffix
func (v *Version) NextPatch() {
	v.Patch++
	v.Suffix = ""
}

func (v *Version) String() string {
	suffix := ""
	if v.Suffix != "" {
		suffix = "-" + v.Suffix
	}
	return fmt.Sprintf("%s%d.%d.%d%s", v.Prefix, v.Major, v.Minor, v.Patch, suffix)
}
