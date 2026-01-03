package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type IO struct {
	input  string
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewIO() *IO {
	return &IO{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

// 读取输入
func (v *IO) Read() string {

	readString, _ := v.reader.ReadString('\n')
	return strings.Trim(readString, "\n")
}

// Writeln
// 写入终端
func (v *IO) Writeln(contents ...string) {
	for _, content := range contents {
		v.writer.WriteString(content + "\n")
	}
	v.writer.Flush()
}

func (v *IO) ReadInt() (num int) {
	num, _ = strconv.Atoi(v.Read())
	return

}
