package main

import (
	"testing"
)

func TestParseDate(t *testing.T) {

	_, _, err := convertDate("invalid")
	if err == nil {
		t.Errorf("expected error")
	}
	_, _, err = convertDate("1/1/1/1")
	if err == nil {
		t.Errorf("expected error")
	}

	begin, end, err := convertDate("12/10/2019")
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}

	got := begin.String()
	expected := "2019-10-12 00:00:00 +0000 UTC"
	if got != expected {
		t.Errorf("error: got %q; expected %q\n", got, expected)
	}

	got = end.String()
	expected = "2019-10-13 00:00:00 +0000 UTC"
	if got != expected {
		t.Errorf("error: got %q; expected %q\n", got, expected)
	}
}
