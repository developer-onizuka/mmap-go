package main

import(
	"github.com/edsrzf/mmap-go" // go get github.com/edsrzf/mmap-go
	"os"
	"fmt"
)

func main() {
	f,_ := os.OpenFile(os.Args[1],os.O_RDWR,0755)
	defer f.Close()
	m,_ := mmap.Map(f,mmap.RDWR,0)
	defer m.Unmap()
	fmt.Println(m)

	msg := "HELLO, WORLD!"
	copy(m,[]byte(msg))
	m.Flush()
}
