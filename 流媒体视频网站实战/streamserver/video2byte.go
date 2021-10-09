package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

// 读取文件到[]byte中
func file2Bytes(filename string) error {

	// File
	file, err := os.Open(VIDEO_DIR + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return err
	}

	// []byte
	data := make([]byte, stats.Size())
	count, err := file.Read(data)
	if err != nil {
		return err
	}
	fmt.Printf("read file %s len: %d \n", filename, count)

	newfilename := strings.Split(filename, ".")[0]
	fp, err := os.Create(VIDEO_DIR + newfilename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fp.Close()

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)
	fp.Write(buf.Bytes())

	return err
}
