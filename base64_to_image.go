package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"
)


func Base64ToImage(data string) error {
	input := data
	b64data := input[strings.IndexByte(input, ',')+1:]
	gettype := input[11:15]
	fmt.Println(gettype)
	unbased, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		panic("Cannot decode b64")
	}
	r := bytes.NewReader(unbased)

	if gettype == "png;" {
		im, err := png.Decode(r)
		if err != nil {
			panic("Bad png")
		}

		f, err := os.OpenFile("example.png", os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}

		png.Encode(f, im)
	} else if gettype == "jpeg" {
		im, err := jpeg.Decode(r)
		if err != nil {
			panic("Bad jpeg")
		}
		f, err := os.OpenFile("test.jpg", os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}
		jpeg.Encode(f, im, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}
	return nil
}

func main() {
	data := "Code base64 here"
	Base64ToImage(data)
}
