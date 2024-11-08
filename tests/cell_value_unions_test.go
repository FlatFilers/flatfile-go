package tests

import (
	"encoding/json"
	"testing"

	flatfile "github.com/FlatFilers/flatfile-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCellValueUnion(t *testing.T) {
	testCases := []struct {
		value []byte
	}{
		{value: []byte("0")},
		{value: []byte(`""`)},
		{value: []byte("false")},
	}
	for _, testCase := range testCases {
		t.Run(string(testCase.value), func(t *testing.T) {
			cellValue := &flatfile.CellValueUnion{}
			require.NoError(t, json.Unmarshal(testCase.value, cellValue))

			bytes, err := json.Marshal(cellValue)
			require.NoError(t, err)
			assert.Equal(t, string(testCase.value), string(bytes))
		})
	}
}
