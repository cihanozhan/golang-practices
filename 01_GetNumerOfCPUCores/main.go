package main

// http://www.cihanozhan.com/go-ile-cpu-cekirdek-sayisini-ogrenme/

import (
	"fmt"
	"runtime"
)

func main() {
	cores := runtime.NumCPU()
	fmt.Printf("This machine has %d CPU cores. \n", cores)
	runtime.GOMAXPROCS(cores)
}
