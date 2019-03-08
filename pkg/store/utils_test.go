package store

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_notFound(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Error = nil",
			args: args{err: nil},
			want: false,
		},
		{
			name: "Error = gorm.ErrRecordNotFound",
			args: args{err: gorm.ErrRecordNotFound},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := notFound(tt.args.err); got != tt.want {
				t.Errorf("notFound() = %v, want %v", got, tt.want)
			}
		})
	}

	assert.Panics(t, func() { notFound(errors.New("custom error")) }, "must panic on error")
}

func Test_found(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Error = nil",
			args: args{err: nil},
			want: true,
		},
		{
			name: "Error = gorm.ErrRecordNotFound",
			args: args{err: gorm.ErrRecordNotFound},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := found(tt.args.err); got != tt.want {
				t.Errorf("notFound() = %v, want %v", got, tt.want)
			}
		})
	}

	assert.Panics(t, func() { found(errors.New("custom error")) }, "must panic on error")
}
