package version

import "testing"

func TestVersion_NextMajor(t *testing.T) {
	v := Version{
		Prefix: "",
		Major:  1,
		Minor:  2,
		Patch:  3,
		Suffix: "abc",
	}

	v.NextMajor()

	if v.Major != 2 {
		t.Fatal("Major not increased")
	}

	if v.Minor != 0 {
		t.Fatal("Minor not reseted")
	}

	if v.Patch != 0 {
		t.Fatal("Patch not reseted")
	}

	if v.Suffix != "" {
		t.Fatal("Suffix not reseted")
	}
}

func TestVersion_NextMinor(t *testing.T) {
	v := Version{
		Prefix: "",
		Major:  1,
		Minor:  2,
		Patch:  3,
		Suffix: "abc",
	}

	v.NextMinor()

	if v.Major != 1 {
		t.Fatal("Major should not be increased")
	}

	if v.Minor != 3 {
		t.Fatal("Minor not reseted")
	}

	if v.Patch != 0 {
		t.Fatal("Patch not reseted")
	}

	if v.Suffix != "" {
		t.Fatal("Suffix not reseted")
	}
}

func TestVersion_NextPatch(t *testing.T) {
	v := Version{
		Prefix: "",
		Major:  1,
		Minor:  2,
		Patch:  3,
		Suffix: "abc",
	}

	v.NextPatch()

	if v.Major != 1 {
		t.Fatal("Major should not be increased")
	}

	if v.Minor != 2 {
		t.Fatal("Minor should not be reseted")
	}

	if v.Patch != 4 {
		t.Fatal("Patch should not be reseted")
	}

	if v.Suffix != "" {
		t.Fatal("Suffix not reseted")
	}
}
