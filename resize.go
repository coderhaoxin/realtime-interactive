package resize

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

func Resize(oriImgPath, destImgPath, format string, width, height uint, force bool) (string, error) {
	oriImgFullPath, err := filepath.Abs(oriImgPath)

	if err != nil {
		return "", err
	}

	destImgFullPath, err := filepath.Abs(destImgPath)

	if err != nil {
		return "", err
	}

	imgFile, err := os.Open(oriImgFullPath)

	if err != nil {
		return "", err
	}

	var img image.Image

	if format == "jpeg" || format == "jpg" {
		img, err = jpeg.Decode(imgFile)
	} else if format == "png" {
		img, err = png.Decode(imgFile)
	} else {
		return "", errors.New("unknown format")
	}

	imgFile.Close()
	if err != nil {
		return "", err
	}

	if force {
		img = resize.Resize(width, height, img, resize.Lanczos3)
	} else {
		img = resize.Thumbnail(width, height, img, resize.Lanczos3)
	}

	out, err := os.Create(destImgFullPath)
	defer out.Close()

	if err != nil {
		return "", err
	}

	if format == "jpeg" || format == "jpg" {
		jpeg.Encode(out, img, nil)
	} else {
		// png
		png.Encode(out, img)
	}

	return destImgFullPath, nil
}
