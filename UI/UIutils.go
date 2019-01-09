package UI

import (
	"fmt"
	"runtime"
)

func PrintMemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	res := fmt.Sprintf("Alloc = %v MiB\n", bToMb(m.Alloc))
	res += fmt.Sprintf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	res += fmt.Sprintf("Sys = %v MiB\n", bToMb(m.Sys))
	res += fmt.Sprintf("NumGC = %v\n\n", m.NumGC)
	return res
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
