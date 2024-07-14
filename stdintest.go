package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var count int
var f *os.File

func rotate() *os.File {
	if f != nil {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}

	fname := fmt.Sprintf("./dump-%d.bin", count)
	count++
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("rotated:", fname)

	return f
}

func main() {
	var n, written int
	reader := os.Stdin

	for {
		if written == 0 || written >= 4096*10 {
			f = rotate()
			written = 0
		}

		n, err := io.CopyN(f, reader, 4096)
		if err != nil {
			log.Fatal(err)
		}
		written += n
		log.Println("written:", written)
		time.Sleep(time.Millisecond * 500)
	}
}
