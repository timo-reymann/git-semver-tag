package version

import (
	"strconv"
	"strings"
)

// ProcessVersionTag parses the given version into a version struct
func ProcessVersionTag(tag string) Version {
	v := Version{}

	if strings.HasPrefix(tag, "v") {
		v.Prefix = "v"
		tag = tag[1:]
	}

	if strings.Contains(tag, "-") {
		suffixBegin := strings.Index(tag, "-")
		v.Suffix = tag[suffixBegin+1:]
		tag = tag[:suffixBegin]
	}

	versionParts := strings.Split(tag, ".")
	if len(versionParts) == 3 {
		v.Major, _ = strconv.Atoi(versionParts[0])
		v.Minor, _ = strconv.Atoi(versionParts[1])
		v.Patch, _ = strconv.Atoi(versionParts[2])
	}

	return v
}
