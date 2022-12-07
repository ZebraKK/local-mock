package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	files := []string{"64K", "2M"}
	for _, file := range files {
		size := pickSize(file)
		newFile(file, size)
	}
}

func pickSize(filename string) (size int64) {
	f := func(c rune) bool {
		return unicode.IsLetter(c)
	}
	index := strings.IndexFunc(filename, f)

	size, _ = strconv.ParseInt(filename[:index], 10, 64)
	name := filename[index:]
	switch name {
	case "K":
		size = size * 1024
	case "M":
		size = size * 1024 * 1024
	}
	fmt.Println("index:", index, " ,name:", name, " ,size:", size)
	return
}

// 文件名为 100K/ 200M
func newFile(name string, size int64) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("openfile err:", err)
		return
	}

	l := int64(1024)
	buf := make([]byte, l)
	var index int64
	for size > 0 {
		step := l
		if size > l {
			size = size - l
		} else {
			step = size
			size = 0
		}

		rand.Read(buf[:step])
		n, err := f.WriteAt(buf, index)
		if err != nil {
			fmt.Println("write file err :", err)
			return
		}
		index += int64(n)
	}

}
