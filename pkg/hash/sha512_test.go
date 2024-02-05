package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash512(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "#1",
			args: args{
				src: "mypass",
			},
			want: "1c573dfeb388b562b55948af954a7b344dde1cc2099e978a992790429e7c01a4205506a93d9aef3bab32d6f06d75b7777a7ad8859e672fedb6a096ae369254d2",
			err:  nil,
		},
		{
			name: "#2",
			args: args{
				src: "",
			},
			want: "",
			err:  ErrEmptyPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sha512(tt.args.src)
			if err != nil {
				assert.ErrorIs(t, err, tt.err)
				return
			}

			assert.Equalf(t, tt.want, got, "want: %s actual: %s", tt.want, got)
		})
	}
}
