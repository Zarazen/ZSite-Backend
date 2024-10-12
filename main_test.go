package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		// TODO: Add test cases.
		{name: "Basic test", args: args{a: 1, b: 2}, wantC: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := add(tt.args.a, tt.args.b); gotC != tt.wantC {
				t.Errorf("add() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
