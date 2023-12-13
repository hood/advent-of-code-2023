package shared

import (
	"fmt"
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, actual interface{}, notes ...string) {
	if expected != actual {
		t.Errorf("\nExpected:\n%v\n\nGot:\n%v", expected, actual)

		if len(notes) > 0 {
			fmt.Println(notes)
		}

		t.FailNow()
	}
}
