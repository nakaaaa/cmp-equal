package cmpequal

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Equal(t *testing.T, expected interface{}, actual interface{}, opts ...cmp.Option) bool {
	t.Helper()

	diff := cmp.Diff(expected, actual, opts...)
	if diff != "" {
		t.Errorf(diff)
		return false
	}
	return true
}
