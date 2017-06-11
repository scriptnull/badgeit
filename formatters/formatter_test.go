package formatters

import "testing"

func TestNewFormatter(t *testing.T) {
	// pass in empty string
	output, err := NewFormatter("")
	if err == nil {
		t.Error("Expected Error, but got", err)
	}
	if _, ok := output.(AllFormatter); !ok {
		t.Error("Expected output to be default all formatter")
	}

	// pass in invalid value
	output, err = NewFormatter("invalidArg")
	if err == nil {
		t.Error("Expected Error, but got", err)
	}
	if _, ok := output.(AllFormatter); !ok {
		t.Error("Expected output to be allFormatter")
	}

	// check -f="all"
	output, err = NewFormatter("all")
	if err != nil {
		t.Error("Expected Error, but got", err)
	}
	if _, ok := output.(AllFormatter); !ok {
		t.Error("Expected output to be AllFormatter")
	}

	// check -f="min"
	output, err = NewFormatter("min")
	if err != nil {
		t.Error("Expected Error, but got", err)
	}
	if _, ok := output.(MinFormatter); !ok {
		t.Error("Expected output to be MinFormatter")
	}
}
