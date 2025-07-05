package byter

import (
	"fmt"
)

const (
	_ float64 = 1 << (10 * iota)
	kB
	mB
	gB
	tB
)

func Bytes(bytes uint64) string {
	fbytes := float64(bytes)

	if fbytes >= tB {
		return fmt.Sprintf("%.2fTB", fbytes/tB)
	}

	if fbytes >= gB {
		return fmt.Sprintf("%.2fGB", fbytes/gB)
	}

	if fbytes >= mB {
		return fmt.Sprintf("%.2fMB", fbytes/mB)
	}

	if fbytes >= kB {
		return fmt.Sprintf("%.2fKB", fbytes/kB)
	}

	return fmt.Sprintf("%dB", bytes)

}

const (
	kb = 1e3
	mb = 1e6
	gb = 1e9
	tb = 1e12
)

func Bits(bits uint64) string {

	fbits := float64(bits)
	if fbits >= tb {
		return fmt.Sprintf("%.2fTb", fbits/tb)
	}

	if fbits >= gb {
		return fmt.Sprintf("%.2fGb", fbits/gb)
	}

	if fbits >= mb {
		return fmt.Sprintf("%.2fMb", fbits/mb)
	}

	if fbits >= kb {
		return fmt.Sprintf("%.2fKb", fbits/kb)
	}

	return fmt.Sprintf("%db", bits)

}
