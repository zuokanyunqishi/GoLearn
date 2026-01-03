package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	dial, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriterSize(dial, 512)
	reader := bufio.NewReaderSize(os.Stdin, 512)
	connerReader := bufio.NewReader(dial)

	for {
		str, _ := reader.ReadString('\n')
		fmt.Println("read", str)

		writer.Write([]byte("hello word"))
		writer.Flush()
		b, _ := connerReader.ReadByte()
		fmt.Println(string(b))
	}

}
