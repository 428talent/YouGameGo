package util

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func Base64toJpg(data string, filename string) (string, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		return "", err
	}
	jpgFilename := fmt.Sprintf("static/upload/user/avatar/%s.jpg", filename)
	f, err := os.OpenFile(jpgFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}
	defer f.Close()
	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		return "", err
	}
	return jpgFilename, nil
}

func Base64toPng(data string, filename string) (string, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		return "", err
	}

	//Encode from image format to writer
	pngFilename := fmt.Sprintf("static/upload/user/avatar/%s.png", filename)
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}
	defer f.Close()

	err = png.Encode(f, m)
	if err != nil {
		return "", err
	}
	return pngFilename, nil
}
