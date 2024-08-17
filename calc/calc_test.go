package calc

import (
	"testing"
)

func TestIsEven(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "偶数の場合",
			args: args{n: 2},
			want: true,
		},
		{
			name: "奇数の場合",
			args: args{n: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.n); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "偶数の場合",
			args: args{n: 2},
			want: false,
		},
		{
			name: "奇数の場合",
			args: args{n: 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOdd(tt.args.n); got != tt.want {
				t.Errorf("IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
