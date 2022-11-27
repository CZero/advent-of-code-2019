// Secure container https://adventofcode.com/2019/day/4
package main

import (
	"testing"
)

func Test_findFirst(t *testing.T) {
	type args struct {
		start int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "223450",
			args: args{
				start: 223450,
			},
			want: 223455,
		},
		{
			name: "153456",
			args: args{
				start: 153456,
			},
			want: 155555,
		},
		{
			name: "123456",
			args: args{
				start: 123456,
			},
			want: 123456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirst(tt.args.start); got != tt.want {
				t.Errorf("findFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "111111",
			args: args{
				111111,
			},
			want: true,
		},
		{
			name: "122456",
			args: args{
				122456,
			},
			want: true,
		},
		{
			name: "123789",
			args: args{
				123789,
			},
			want: false,
		},
		{
			name: "223450",
			args: args{
				223450,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.input); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
