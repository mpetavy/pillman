package main

import (
	"github.com/faiface/pixel"
	"github.com/mpetavy/common"
	"image"
	"os"
)

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if common.Error(err) {
		return nil, err
	}
	defer func() {
		common.Error(file.Close())
	}()

	img, _, err := image.Decode(file)
	if common.Error(err) {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
