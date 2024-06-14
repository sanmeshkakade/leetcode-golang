package addbinary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
		{
			name: "two",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := addBinary(test.args.a, test.args.b)

			assert.Equal(t, test.want, got)
		})
	}
}
