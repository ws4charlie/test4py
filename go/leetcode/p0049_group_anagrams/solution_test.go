package p0049

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
		{nil, nil},
	}
	for _, tt := range tests {
		if got := GroupAnagrams(tt.strs); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("GroupAnagrams(%q) = %q, want %q", tt.strs, got, tt.want)
		}
	}
}
