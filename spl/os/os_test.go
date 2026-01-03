package os__test

import (
	"fmt"
	"os"
	"testing"
)

func TestOsCreate(t *testing.T) {

	filePath := os.TempDir() + "/aaaaaaa"
	file, err := os.Create(filePath)

	if err != nil {
		t.Error(err)
	}

	defer func() {
		file.Close()
		os.Remove(filePath)
	}()

	file.WriteString("hello,word")
	file.WriteString("hello,china")
	fileInfo, err := os.Stat(filePath)

	fmt.Println(fileInfo.Size(), fileInfo.Mode())

	fmt.Println(os.IsNotExist(err), fileInfo)
}
