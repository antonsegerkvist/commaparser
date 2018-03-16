package commaparser

import "testing"

func TestAll(t *testing.T) {
	str := "*"
	expectedOutput := []string{"id", "color", "time", "value"}

	hash, err := ParseString(str, expectedOutput)

	if err != nil {
		t.Error(`==> ParseString: Got error from string:`, str)
	}

	for _, v := range expectedOutput {
		if hash[v] != true {
			t.Error(`==> ParseString: String`, v, `not found in output`)
		}
	}
}

func TestSimple(t *testing.T) {
	str := "id,color,time"
	fields := []string{"id", "color", "time", "value"}

	hash, err := ParseString(str, fields)

	if err != nil {
		t.Error(`==> ParseString: Got error from string:`, str)
	}

	if hash[fields[0]] != true {
		t.Error(`==> ParseString: String`, fields[0], `not found in output`)
	}

	if hash[fields[1]] != true {
		t.Error(`==> ParseString: String`, fields[1], `not found in output`)
	}

	if hash[fields[2]] != true {
		t.Error(`==> ParseString: String`, fields[2], `not found in output`)
	}

	if hash[fields[3]] != false {
		t.Error(`==> ParseString: String`, fields[3], `found in output`)
	}

}
