// package main
// import (
// "fmt"
// "os"
// )
// func main() {
// file, err := os.Open("fileread.go")
// if err != nil {
// return
// }
// defer file.Close()
// // getting size of file
// stat, err := file.Stat()
// if err != nil {
// return
// }
// // reading file
// bs := make([]byte, stat.Size())
// _, err = file.Read(bs)
// if err != nil {
// return
// }
// fmt.Println(string(bs))
// }

import (
	"bytes"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open("fileread.go")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var n int64 = bytes.MinRead

	if fi, err := f.Stat(); err == nil {
		if size := fi.Size() + bytes.MinRead; size > n {
			n = size
		}
	}
	return readAll(file, n)
}