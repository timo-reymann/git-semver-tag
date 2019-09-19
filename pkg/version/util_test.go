package version

import "testing"

func TestProcessVersionTag(t *testing.T) {
	var testCases = []struct {
		input           string
		expectedVersion Version
	}{
		{
			"1.0.0",
			Version{
				Prefix: "",
				Major:  1,
				Minor:  0,
				Patch:  0,
				Suffix: "",
			},
		},
		{
			"1.1.0",
			Version{
				Prefix: "",
				Major:  1,
				Minor:  1,
				Patch:  0,
				Suffix: "",
			},
		},
		{
			"1.1.1",
			Version{
				Prefix: "",
				Major:  1,
				Minor:  1,
				Patch:  1,
				Suffix: "",
			},
		},
		{
			"v1.1.1",
			Version{
				Prefix: "v",
				Major:  1,
				Minor:  1,
				Patch:  1,
				Suffix: "",
			},
		},
		{
			"1.1.1-beta1",
			Version{
				Prefix: "",
				Major:  1,
				Minor:  1,
				Patch:  1,
				Suffix: "beta1",
			},
		},
	}

	for _, tc := range testCases {
		version := ProcessVersionTag(tc.input)

		if version.Prefix != tc.expectedVersion.Prefix {
			t.Error("Prefix doesnt match")
		}

		if version.Minor != tc.expectedVersion.Minor {
			t.Error("Minor version doesnt match")
		}

		if version.Major != tc.expectedVersion.Major {
			t.Errorf("Major version doesnt match")
		}

		if version.Patch != tc.expectedVersion.Patch {
			t.Error("Patch version doesnt match")
		}

		if version.Suffix != tc.expectedVersion.Suffix {
			t.Error("Suffix doesnt match")
		}
	}
}
