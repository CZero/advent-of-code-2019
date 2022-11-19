// https://adventofcode.com/2019/day/1
// --- Day 1: The Tyranny of the Rocket Equation ---

package main

import "testing"

func TestCalcModuleFuel(t *testing.T) {
	type args struct {
		mass string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				mass: "12",
			},
			want: 2,
		},
		{
			args: args{
				mass: "14",
			},
			want: 2,
		},
		{
			args: args{
				mass: "1969",
			},
			want: 654,
		},
		{
			args: args{
				mass: "100756",
			},
			want: 33583,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcModuleFuel(tt.args.mass); got != tt.want {
				t.Errorf("CalcModuleFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
