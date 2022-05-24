package libs

import (
	"fmt"
	"os"
)

func Greet() {
	println("hello world")
}

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

/*
var (
		reader libs.FooReader
		writer libs.FooWriter
	)
	// input := make([]byte, 4096)
	// s, err := reader.Read(input)
	// if err != nil {
	// 	log.Fatalln("Unable to read data")
	// }
	// fmt.Printf("Read %d bytes from stdin\n", s)
	// s, err = writer.Write(input)
	// if err != nil {
	// 	log.Fatalln("Unable to write data")
	// }
	// fmt.Printf("Wrote %d bytes to stdout\n", s)

	/* using copy */
// 	if s, err := io.Copy(&writer, &reader); err != nil {
// 		log.Fatalln("Unable to read/write data")
// 		}
// fmt.Printf("Read %d bytes from stdin\n", s)
