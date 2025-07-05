package bkmgt

import (
	"fmt"
)

const (
	_ float64 = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

func Bytes(bytes uint64) string {
	fbytes := float64(bytes)

	if fbytes >= TB {
		return fmt.Sprintf("%.2fTB", fbytes/TB)
	}

	if fbytes >= GB {
		return fmt.Sprintf("%.2fGB", fbytes/GB)
	}

	if fbytes >= MB {
		return fmt.Sprintf("%.2fMB", fbytes/MB)
	}

	if fbytes >= KB {
		return fmt.Sprintf("%.2fKB", fbytes/KB)
	}

	return fmt.Sprintf("%dB", bytes)

}

const (
	Kb = 1e3
	Mb = 1e6
	Gb = 1e9
	Tb = 1e12
)

func Bits(bits uint64) string {

	fbits := float64(bits)
	if fbits >= Tb {
		return fmt.Sprintf("%.2fTb", fbits/Tb)
	}

	if fbits >= Gb {
		return fmt.Sprintf("%.2fGb", fbits/Gb)
	}

	if fbits >= Mb {
		return fmt.Sprintf("%.2fMb", fbits/Mb)
	}

	if fbits >= Kb {
		return fmt.Sprintf("%.2fKb", fbits/Kb)
	}

	return fmt.Sprintf("%db", bits)

}
