package parsers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDate(t *testing.T) {

	_, _, err := ConvertDate("invalid")
	require.Error(t, err)

	_, _, err = ConvertDate("1/1/1/1")
	require.Error(t, err)

	begin, end, err := ConvertDate("12/10/2019")
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	require.NoError(t, err)

	got := begin.String()
	expected := "2019-10-12 00:00:00 +0000 UTC"
	require.Equal(t, got, expected)

	got = end.String()
	expected = "2019-10-13 00:00:00 +0000 UTC"
	require.Equal(t, got, expected)
}
