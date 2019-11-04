package media

import (
	"encoding/base64"
	"errors"
	"image"
	"image/png"
	"os"
	"strings"
)

var (
	ErrWrongBase64Format = errors.New("Wrong Image Data Format ")
)

// ImageData represent the uploaded data
type ImageData struct {
	Title   string
	Picture string
}

// GetPicture returns the Image and the error
func (im *ImageData) GetPicture() (image.Image, error) {
	parts := strings.Split(im.Picture, ",")
	if len(parts) != 2 {
		return nil, ErrWrongBase64Format
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(parts[1]))
	m, _, err := image.Decode(reader)
	return m, err
}

// SavePng Saves the image as a png, returns an error
func SavePng(img image.Image, fileName string) error {

	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	return png.Encode(out, img)
}

// SaveMp4
func SaveBts(bts []byte, fileName string) error {

	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(bts)

	return err
}
