package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	list := []string{"Larry", "Curly", "Moe", "Curly"}
	set := SliceToSet[string](list)
	assert.True(t, set.Contains("Larry"))
	assert.True(t, set.Contains("Curly"))
	assert.False(t, set.Contains("Curly Joe"))
}

func TestSliceToSet(t *testing.T) {
	tests := []struct {
		name string
		list []string
		want Set[string]
	}{
		{"No dups",
			[]string{
				"Larry",
				"Curly",
				"Moe",
			},
			SliceToSet[string](
				[]string{
					"Larry",
					"Curly",
					"Moe",
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := SliceToSet[string](tt.list)
			assert.True(t, want.Equal(have))
		})
	}
}
