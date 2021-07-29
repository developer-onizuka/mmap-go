# Memory Mapped File
# 1.
```
$ go get github.com/edsrzf/mmap-go
```

# 2.
```
$ echo "hello, world!" > test.txt
```

# 3. 
```
$ go run mmap.go test.txt
or
$ go run mmap-syscall.go test.txt
```

# 4.
```
$ cat test.txt 
HELLO, WORLD!
```

# Code Casting
# 1.
```
$ go tool compile -S -N inc.go
"".inc STEXT nosplit size=23 args=0x10 locals=0x0
	0x0000 00000 (inc.go:3)	TEXT	"".inc(SB), NOSPLIT|ABIInternal, $0-16
	0x0000 00000 (inc.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (inc.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (inc.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (inc.go:3)	PCDATA	$0, $0
	0x0000 00000 (inc.go:3)	PCDATA	$1, $0
	0x0000 00000 (inc.go:3)	MOVQ	$0, "".~r1+16(SP)
	0x0009 00009 (inc.go:4)	MOVQ	"".n+8(SP), AX
	0x000e 00014 (inc.go:4)	INCQ	AX
	0x0011 00017 (inc.go:4)	MOVQ	AX, "".~r1+16(SP)
	0x0016 00022 (inc.go:4)	RET
	0x0000 48 c7 44 24 10 00 00 00 00 48 8b 44 24 08 48 ff  H.D$.....H.D$.H.
	0x0010 c0 48 89 44 24 10 c3                             .H.D$..
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 69 6e 63                                         inc
go.loc."".inc SDWARFLOC size=0
go.info."".inc SDWARFINFO size=57
	0x0000 03 22 22 2e 69 6e 63 00 00 00 00 00 00 00 00 00  ."".inc.........
	0x0010 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0f  ................
	0x0020 6e 00 00 03 00 00 00 00 01 9c 0f 7e 72 31 00 01  n..........~r1..
	0x0030 03 00 00 00 00 02 91 08 00                       .........
	rel 8+8 t=1 "".inc+0
	rel 16+8 t=1 "".inc+23
	rel 26+4 t=29 gofile../mnt/go/src/mmap-go/inc.go+0
	rel 36+4 t=28 go.info.int+0
	rel 49+4 t=28 go.info.int+0
go.range."".inc SDWARFRANGE size=0
go.isstmt."".inc SDWARFMISC size=0
	0x0000 04 0e 01 09 00                                   .....
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
```

# 2.
```
$ go tool objdump -S inc.o
TEXT %22%22.inc(SB) gofile../mnt/go/src/mmap-go/inc.go
func inc(n int) int {
  0x1e9			48c744241000000000	MOVQ $0x0, 0x10(SP)	
	return n + 1
  0x1f2			488b442408		MOVQ 0x8(SP), AX	
  0x1f7			48ffc0			INCQ AX			
  0x1fa			4889442410		MOVQ AX, 0x10(SP)	
  0x1ff			c3			RET
```

# 3.
You can cast the code above to "func(int) int". 
But Please note the address to cast is not silce's pointer it self but the address of slice's pointer.
See also https://www.techscore.com/tech/Go/Lang/Basic14/

```
----- output example of "go run mmap-exec.go -----
&ptr(=&&prog): 0xc000084020
unsafe.Pointer(&ptr): (unsafe.Pointer)(0xc000084020)
unsafe.Pointer(ptr): (unsafe.Pointer)(0xc000078060)
inc: (func(int) int)(0x7fbadc4bc000)
&inc: (*func(int) int)(0xc000084028)
--------------------------------------------------

                                                                               inc (type (func(int) int))
                                        prog (type []byte)                     +--------------+ 0xc000084028
type []byte                             +----------------+ 0xc000078060 <-----<| 0xc000078060 | (&ptr) 
+---------------+ 0x7fbadc4bc000 <-----<| 0x7fbadc4bc000 | (&prog = ptr)       +--------------+
| prog[0]=0x48  |                       +----------------+                     | 0xc000078060 |
+---------------+ 0x7fbadc4bc001        | len=23         |                     +--------------+
| prog[1]=0xc7  |                       +----------------+
+---------------+ 0x7fbadc4bc002        | cap=23         |
| prog[2]=0x44  |                       +----------------+
+---------------+ 0x7fbadc4bc003
| prog[3]=0x24  |
+---------------+ 0x7fbadc4bc004
| prog[4]=0x10  |
+---------------+ 
       ...
+---------------+ 
| prog[22]=0xc3 |
+---------------+ 
```
```
        copy(prog,code)
        ptr := &prog
	
        fmt.Printf("unsafe.Pointer(&ptr): %#v\n", unsafe.Pointer(&ptr)) //this is the target for casting to func.
        fmt.Printf("unsafe.Pointer(ptr): %#v\n", unsafe.Pointer(ptr)) //unable to be casted because it's already casted by []byte.
        inc := *(*func(int) int)(unsafe.Pointer(&ptr)) //casting to the address of the prog's pointer, not prog's pointer directory.
        // inc := *(*func(int) int)(unsafe.Pointer(ptr)) //unable to be casted because it's already casted by []byte.
        // inc := *(*func(int) int)(&ptr) // cannot convert &ptr (type **[]byte) to type *func(int) int.
```
