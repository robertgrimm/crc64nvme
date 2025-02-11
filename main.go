package main

import (
	"encoding/base64"
	"fmt"
	"hash/crc64"
	"io"
	"os"
)

// The NVME polynomial
const nvme = 0x9A6C9329AC4BC9B5

func main() {
	crc := crc64.New(crc64.MakeTable(nvme))

	_, err := io.Copy(crc, os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	sum := crc.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(sum))
}
