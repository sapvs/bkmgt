package byter

import (
	"fmt"
	"testing"
)

func ExampleBits() {
	fmt.Println(Bits(16_000_000_000))
	// Output: 16.00Gb
}

func ExampleBytes() {
	fmt.Println(Bytes(16_000_000_000))
	// Output: 14.90GB
}

func TestBytes(t *testing.T) {
	t.Parallel()
	type args struct {
		bytes uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "GB_1", args: args{bytes: 16_000_000_000}, want: "14.90GB"},
		{name: "GB_2", args: args{bytes: 1_000_000_000}, want: "953.67MB"},
		{name: "MB", args: args{bytes: 16_000_000}, want: "15.26MB"},
		{name: "KB", args: args{bytes: 16_000}, want: "15.62KB"},
		{name: "K_2", args: args{bytes: 1024}, want: "1.00KB"},
		{name: "B_1", args: args{bytes: 16}, want: "16B"}, {name: "B_2", args: args{bytes: 999}, want: "999B"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes(tt.args.bytes); got != tt.want {
				t.Errorf("Bytes() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestBits(t *testing.T) {
	t.Parallel()
	type args struct {
		bits uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Gb_1", args: args{bits: 16_000_000_000}, want: "16.00Gb"},
		{name: "Gb_2", args: args{bits: 1_000_000_000}, want: "1.00Gb"},
		{name: "Mb", args: args{bits: 16_000_000}, want: "16.00Mb"},
		{name: "Kb", args: args{bits: 16_000}, want: "16.00Kb"},
		{name: "K_2", args: args{bits: 1000}, want: "1.00Kb"},
		{name: "b_1", args: args{bits: 16}, want: "16b"},
		{name: "b_2", args: args{bits: 999}, want: "999b"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bits(tt.args.bits); got != tt.want {
				t.Errorf("Bits() = %s, want %s", got, tt.want)
			}
		})
	}
}
