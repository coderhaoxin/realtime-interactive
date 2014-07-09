package lib

import (
	"path/filepath"
	"testing"
)

func TestResize01(t *testing.T) {
	oriImgPath, err := filepath.Abs("../fixture/2014-07-06.jpg")

	if err != nil {
		t.Error(err)
	}

	destImgPath, err := filepath.Abs("../fixture/2014-07-06.jpg.300x300.jpg")

	if err != nil {
		t.Error(err)
	}

	Resize(oriImgPath, destImgPath, "jpg", 300, 300, false)
}

func TestResize02(t *testing.T) {
	oriImgPath, err := filepath.Abs("../fixture/2014-07-06.jpg")

	if err != nil {
		t.Error(err)
	}

	destImgPath, err := filepath.Abs("../fixture/2014-07-06.jpg.300x300xforce.jpg")

	if err != nil {
		t.Error(err)
	}

	Resize(oriImgPath, destImgPath, "jpg", 300, 300, true)
}
