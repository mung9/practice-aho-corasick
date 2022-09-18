package aho

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	tests := []struct {
		words []string
		input string
		want  []string
	}{
		{
			words: []string{"사과", "배", "복숭아"},
			input: "사과를 먹었더니 복숭아도 먹고싶네?",
			want:  []string{"사과", "복숭아"},
		},
		{
			words: []string{"치킨", "통닭", "후라이드"},
			input: "맛있는 치킨 강남역점",
			want:  []string{"치킨"},
		},
		{
			words: []string{},
			input: "스타벅스 교보타워점",
			want:  []string{},
		},
		{
			words: []string{"AGW", "AGWPV", "OWK", "APO", "XIC", "ZZK", "COAG", "XI", "ICO"},
			input: "AGXICOAGWP",
			want:  []string{"XI", "XIC", "ICO", "COAG", "AGW"},
		},
		{
			words: []string{"A", "AA", "AAA", "AAAA", "AAAAA"},
			input: "BBABAAAB",
			want:  []string{"A", "AA", "AAA"},
		},
	}

	for _, test := range tests {
		aho := New()
		for _, w := range test.words {
			aho.AddWord(w)
		}
		aho.Build()

		actual := aho.FindAll(test.input)

		sort.Strings(actual)
		sort.Strings(test.want)
		assert.Equal(t, test.want, actual)
	}
}
