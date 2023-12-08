package shared

import "testing"

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("\nExpected %v\n\nGot %v", expected, actual)
		t.FailNow()
	}
}
