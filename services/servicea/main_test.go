package main

import "testing"

func Test_power2(t *testing.T) {

	type args struct {
		num int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				num: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		if got := power2(tt.args.num); got != tt.want {
			t.Errorf("power2() = %v, want %v", got, tt.want)
		}
	}
}