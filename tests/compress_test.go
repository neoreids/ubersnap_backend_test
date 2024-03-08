package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestCompress(t *testing.T) {
	newFileName := fmt.Sprintf("compressed%s", filepath.Ext(Image2MB))
	file, err := os.Open(Image2MB)
	if err != nil {
		t.Fatalf("error open image file test : %v", err)
	}

	defer file.Close()
	matFile, bufferByte, err := ImageLogic.CompressImage(file, 50, fmt.Sprintf("./%s", newFileName))
	if err != nil {
		t.Fatalf("error compressing file : %v", err)
	}

	defer matFile.Close()

	originalFileInfo, errGetInfo := file.Stat()
	if errGetInfo != nil {
		t.Fatalf("error get info of original file : %v", errGetInfo)
	}
	assert.NotEqual(t, bufferByte.Len(), originalFileInfo.Size())
}
