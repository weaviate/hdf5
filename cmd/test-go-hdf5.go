package main

import (
	// stdlib
	"fmt"

	// local
	"hdf5"
)

func main() {
	fmt.Println("=== go-hdf5 ===")
	m,n,r,err := hdf5.GetLibVersion()
	if err != nil {
		fmt.Printf("** error ** %s\n", err)
		return
	}
	fmt.Printf("=== version: %d.%d.%d\n", m, n, r)
	fmt.Println("=== bye.")
}