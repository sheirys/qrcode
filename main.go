package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {

	content := flag.String("content", "", "content of QR code.")
	size := flag.Int("size", 200, "image size. if N is provided then image NxN will be generated.")
	file := flag.String("file", "", "where to save generated QR code. This must be and png file.")

	flag.Parse()

	if *content == "" {
		fmt.Println("qr content is not provided")
		os.Exit(1)
	}

	qrCode, err := qr.Encode(*content, qr.M, qr.Auto)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	qrCode, err = barcode.Scale(qrCode, *size, *size)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create the output file
	fd, err := os.Create(*file)
	defer fd.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	png.Encode(fd, qrCode)
}
