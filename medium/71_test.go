package medium

import "testing"

func TestSimplifyPath(t *testing.T) {
	paths := []string{
		"/home/",
		"/../",
		"/home//foo/",
		"/a/./b/../../c/",
		"/a/../../b/../c//.//",
	}
	excepts := []string{
		"/home",
		"/",
		"/home/foo",
		"/c",
		"/c",
	}

	for i, path := range paths {
		result := SimplifyPath(path)
		if result != excepts[i] {
			t.Logf("except %s, actual %s", excepts[i], path)
			t.FailNow()
		}
	}
}
