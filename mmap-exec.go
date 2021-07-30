package main

import (
	//"os"
	"fmt"
	"syscall"
	"unsafe"
	//"time"
)

func main() {
	fmt.Printf("---------- code casting ----------\n")
	code := []byte{
		0x48,0xc7,0x44,0x24,0x10,0x00,0x00,0x00,0x00,
		0x48,0x8b,0x44,0x24,0x08,
		0x48,0xff,0xc0,
		0x48,0x89,0x44,0x24,0x10,
		0xc3,
	}
	fmt.Printf("len(code): %v [byte]\n",len(code))
	fmt.Printf("cap(code): %v [byte]\n",cap(code))
	fmt.Printf("code: %#v\n", code)
	fmt.Printf("&code: %p\n", &code)
	i := 0
	for i < len(code) {
		fmt.Printf("&code[%v]: %p, code[%v]: %#v\n", i, &code[i], i, code[i])
		i++
	} 
	fmt.Printf("-----\n")

	/*
	func inc(n int) int {
  	0x1e9			48c744241000000000	MOVQ $0x0, 0x10(SP)	
		return n + 1
  	0x1f2			488b442408		MOVQ 0x8(SP), AX	
  	0x1f7			48ffc0			INCQ AX			
  	0x1fa			4889442410		MOVQ AX, 0x10(SP)	
  	0x1ff			c3			RET
	*/

	prog, _ := syscall.Mmap(0, 0, len(code), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_ANON|syscall.MAP_SHARED)
	//fmt.Printf("err: %v\n", err)
	defer syscall.Munmap(prog)

	copy(prog,code)
	ptr := &prog
	fmt.Printf("len(prog): %v [byte]\n",len(prog))
	fmt.Printf("cap(prog): %v [byte]\n",cap(prog))
	fmt.Printf("prog: %#v\n", prog)
	fmt.Printf("&prog: %p\n", &prog)

	i = 0
	for i < len(prog) {
		fmt.Printf("&prog[%v]: %p, prog[%v]: %#v\n", i, &prog[i], i, prog[i])
		i++
	} 
	fmt.Printf("-----\n")

	xptr := (*[]byte)(unsafe.Pointer(ptr))
	fmt.Printf("xptr: %#v\n",xptr)
	fmt.Printf("len(*xptr): %v [byte]\n",len(*xptr))
	fmt.Printf("cap(*xptr): %v [byte]\n",cap(*xptr))
	fmt.Printf("-----\n")

	fmt.Printf("&ptr(=&&prog): %p\n", &ptr)
	fmt.Printf("unsafe.Pointer(&ptr): %#v\n", unsafe.Pointer(&ptr)) //this is the target for casting to func.
	fmt.Printf("unsafe.Pointer(ptr): %#v\n", unsafe.Pointer(ptr)) //unable to be casted because it's already casted by []byte.
	inc := *(*func(int) int)(unsafe.Pointer(&ptr)) //casting to the address of the prog's pointer, not prog's pointer directory.
	// inc := *(*func(int) int)(unsafe.Pointer(ptr)) //unable to be casted because it's already casted by []byte.
	// inc := *(*func(int) int)(&ptr) // cannot convert &ptr (type **[]byte) to type *func(int) int.
	fmt.Printf("inc: %#v\n", inc)
	fmt.Printf("&inc: %#v\n", &inc)
	fmt.Printf("&inc: %p\n", &inc)
	fmt.Printf("-----\n")
	fmt.Printf("\n")

	fmt.Println(inc(100)) // 101 (=100+1)
}
