package bits

import (
	"testing"
)

func TestSetBit(t *testing.T) {
	type args struct {
		bitset Bits
		flag   Bits
	}
	tests := []struct {
		name string
		args args
		want Bits
	}{
		{
			name: "1 to 0",
			args: args{
				bitset: 0,
				flag:   1,
			},
			want: 1,
		},
		{
			name: "2 to 1",
			args: args{
				bitset: 1,
				flag:   2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetBit(tt.args.bitset, tt.args.flag); got != tt.want {
				t.Errorf("SetBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearBit(t *testing.T) {
	type args struct {
		bitset Bits
		flag   Bits
	}
	tests := []struct {
		name string
		args args
		want Bits
	}{
		{
			name: "1 to 3",
			args: args{
				bitset: 3,
				flag:   1,
			},
			want: 2,
		},
		{
			name: "1 to 1",
			args: args{
				bitset: 1,
				flag:   1,
			},
			want: 0,
		},
		{
			name: "1 to 2",
			args: args{
				bitset: 2,
				flag:   1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearBit(tt.args.bitset, tt.args.flag); got != tt.want {
				t.Errorf("ClearBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToggleBit(t *testing.T) {
	type args struct {
		bitset Bits
		flag   Bits
	}
	tests := []struct {
		name string
		args args
		want Bits
	}{
		{
			name: "1 to 2",
			args: args{
				bitset: 2,
				flag:   1,
			},
			want: 3,
		},
		{
			name: "1 to 3",
			args: args{
				bitset: 3,
				flag:   1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToggleBit(tt.args.bitset, tt.args.flag); got != tt.want {
				t.Errorf("ToggleBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasBit(t *testing.T) {
	type args struct {
		bitset Bits
		flag   Bits
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1 from 2",
			args: args{
				bitset: 2,
				flag:   1,
			},
			want: false,
		},
		{
			name: "1 from 3",
			args: args{
				bitset: 3,
				flag:   1,
			},
			want: true,
		},
		{
			name: "2 from 2",
			args: args{
				bitset: 2,
				flag:   2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasBit(tt.args.bitset, tt.args.flag); got != tt.want {
				t.Errorf("HasBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
