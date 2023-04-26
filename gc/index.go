package gc

import (
	"fmt"
	"runtime"
)

func Run() {
	var mem runtime.MemStats
	printStats(mem)
	for i := 0; i < 10; i++ {
		a := make([]byte, 50000000)
		if a == nil {
			fmt.Println("make array failed!")
		}
	}
	printStats(mem)
}

func printStats(memoryStats runtime.MemStats) {
	runtime.ReadMemStats(&memoryStats)
	fmt.Println("mem.Alloc:", memoryStats.Alloc)
	fmt.Println("mem.TotalAlloc:", memoryStats.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", memoryStats.HeapAlloc)
	fmt.Println("mem.NumGC:", memoryStats.NumGC)
}
