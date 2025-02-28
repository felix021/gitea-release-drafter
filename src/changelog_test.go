package src

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_readLabelFromTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no label",
			args: args{
				title: "This is a title",
			},
			want: "(uncategorized)",
		},
		{
			name: "label in brackets",
			args: args{
				title: "[label] This is a title",
			},
			want: "label",
		},
		{
			name: "label with colon",
			args: args{
				title: "label: This is a title",
			},
			want: "label",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, readLabelFromTitle(tt.args.title), "readLabelFromTitle(%v)", tt.args.title)
		})
	}
}
