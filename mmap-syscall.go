package main

import (
	"os"
	"fmt"
	"syscall"
	//"unsafe"
	//"time"
)

func main() {
	fmt.Printf("---------- read and write memory mapped file ----------\n")
	f, _ := os.OpenFile(os.Args[1], os.O_RDWR, 755)
	fi, _ := f.Stat()
	defer f.Close()
	data, err := syscall.Mmap(int(f.Fd()), 0, int(fi.Size()), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	fmt.Printf("err: %v\n", err)
	defer syscall.Munmap(data)
	copy(data,[]byte("HELLO, WORLD!\n"))
}
