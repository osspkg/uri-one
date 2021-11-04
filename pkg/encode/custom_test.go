package encode_test

import (
	"testing"

	"github.com/dewep-online/uri-one/pkg/encode"
	"github.com/stretchr/testify/require"
)

func TestEncode_MarshalUnmarshal(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want string
	}{
		{name: "Case1", id: 1, want: "p"},
		{name: "Case1", id: 2, want: "L"},
		{name: "Case1", id: 3, want: "K"},
		{name: "Case1", id: 4, want: "G"},
		{name: "Case1", id: 5, want: "R"},
		{name: "Case1", id: 6, want: "S"},
		{name: "Case1", id: 7, want: "u"},
		{name: "Case1", id: 8, want: "D"},
		{name: "Case1", id: 9, want: "v"},
		{name: "Case2", id: 10, want: "o"},
		{name: "Case3", id: 100, want: "pH"},
		{name: "Case4", id: 1000, want: "PD"},
		{name: "Case5", id: 10000, want: "LIn"},
		{name: "Case6", id: 100000, want: "c0k"},
		{name: "Case7", id: 1000000000, want: "pRmUWP"},
		{name: "Case8", id: 999999, want: "Glvp"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := encode.New()
			h := v.Marshal(tt.id)
			require.Equal(t, tt.want, h)
			id := v.Unmarshal(h)
			require.Equal(t, id, tt.id)
		})
	}
}
