package formatters

import "testing"

func TestNewFormatter(t *testing.T) {
	// pass in empty formatter option
	fmtOpt := FormatterOption{}
	output, err := NewFormatter(fmtOpt)
	if err == nil {
		t.Error("Expected Error, but got", err)
	}
	if _, ok := output.(AllFormatter); !ok {
		t.Error("Expected output to be default all formatter")
	}

	// pass in invalid CmdArg value
	fmtOpt = FormatterOption{
		CmdArgType: "invalidArg",
	}
	output, err = NewFormatter(fmtOpt)
	if err == nil {
		t.Error("Expected Error, but got", err)
	}
	if _, ok := output.(AllFormatter); !ok {
		t.Error("Expected output to be allFormatter")
	}

	// check -f="all"
	fmtOpt = FormatterOption{
		CmdArgType: "all",
	}
	output, err = NewFormatter(fmtOpt)
	if err != nil {
		t.Error("Expected Error, but got", err)
	}
	if _, ok := output.(AllFormatter); !ok {
		t.Error("Expected output to be AllFormatter")
	}

	// check -f="min"
	fmtOpt = FormatterOption{
		CmdArgType: "min",
	}
	output, err = NewFormatter(fmtOpt)
	if err != nil {
		t.Error("Expected Error, but got", err)
	}
	if _, ok := output.(MinFormatter); !ok {
		t.Error("Expected output to be MinFormatter")
	}
}
