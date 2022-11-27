// Day 2: A
package main

import (
	"reflect"
	"testing"
)

func Test_resolveIntcode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name        string
		args        args
		wantRescode []int
	}{
		{
			name: "Res-example1",
			args: args{
				code: "1,9,10,3,2,3,11,0,99,30,40,50",
			},
			wantRescode: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRescode := resolveIntcode(tt.args.code); !reflect.DeepEqual(gotRescode, tt.wantRescode) {
				t.Errorf("resolveIntcode() = %v, want %v", gotRescode, tt.wantRescode)
			}
		})
	}
}
