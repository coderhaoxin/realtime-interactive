package lib

import (
	"testing"
)

func TestParseFileName01(t *testing.T) {
	filename := "a.png_200x100.png"

	filename, formate, width, height, err := ParseFileName(filename)

	if filename != "a.png" || formate != "png" || width != 200 || height != 100 || err != nil {
		t.Error(filename, formate, width, height, err)
	}
}

func TestParseFileName02(t *testing.T) {
	filename := "b.jpg_200x100.jpg"

	filename, formate, width, height, err := ParseFileName(filename)

	if filename != "b.jpg" || formate != "jpg" || width != 200 || height != 100 || err != nil {
		t.Error(filename, formate, width, height, err)
	}
}

func TestParseFileName03(t *testing.T) {
	filename := "c.jpeg_200x100.jpeg"

	filename, formate, width, height, err := ParseFileName(filename)

	if filename != "c.jpeg" || formate != "jpeg" || width != 200 || height != 100 || err != nil {
		t.Error(filename, formate, width, height, err)
	}
}

func TestParseSize(t *testing.T) {
	width, height := ParseSize("600x600")

	if width != 600 || height != 600 {
		t.Error("parse size error", width, height)
	}
}
