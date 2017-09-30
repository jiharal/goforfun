package main

import (
	"fmt"
	"path"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	qr := QRCode{
		Content: "Hello Jihar",
		Size:    300,
	}

	qr.GenerateQrCodeImage("coba.png")
}

// QRCode is struct for creating / initializing qrcode.
type QRCode struct {
	Content   string // qrcode content
	Size      int    // qrcode size. Qrcode is always square
	Filename  string // for accomodate full file path qrcode we want to upload
	imgReview string // private field for accomodate link on first step
}

// Minimal size of qrcode is 300px according to Tokopedia upload tools
// Return error if size less than 300
func validateQrCodeSpec(qr *QRCode) error {
	if qr.Size < 300 {
		return fmt.Errorf("QRCode: Too small QRCode size")
	}

	return nil
}

// GenerateQrCodeImage is function to generate qrcode image.
// Filename is full filepath. Make sure path is available
// to write.
func (qr *QRCode) GenerateQrCodeImage(filename string) error {
	err := validateQrCodeSpec(qr)
	if err != nil {
		return err
	}

	// Clean filepath
	filename = path.Clean(filename)

	err = qrcode.WriteFile(qr.Content, qrcode.Medium, qr.Size, filename)
	if err != nil {
		return err
	}

	qr.Filename = filename
	return nil
}

// GenerateQrCodeImage is function to generate qrcode image resulting in byte array.
func (qr *QRCode) GenerateQrCodeImageByte() ([]byte, error) {
	err := validateQrCodeSpec(qr)
	if err != nil {
		return nil, err
	}

	png, err := qrcode.Encode(qr.Content, qrcode.Medium, qr.Size)
	if err != nil {
		return nil, err
	}

	return png, nil
}
