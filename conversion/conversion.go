package conversion

/**
 * @Time       : 2020/8/3 9:57 下午
 * @Author     : xulp
 * @Description: Provide image format conversion
 */

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"image-helper/util/constant"
	"image-helper/util/extension"
)

type Conversion struct {
	SrcFilePath string
	AimFilePath string
}

func InitConversion(srcFile, aimFile string) *Conversion {
	return &Conversion{
		SrcFilePath: srcFile,
		AimFilePath: aimFile,
	}
}

func (c *Conversion) FormatConversion() {
	srcFile, err := os.Open(c.SrcFilePath)
	if err != nil {
		log.Fatalf("file open failed: %v", err.Error())
		return
	}
	defer srcFile.Close()

	srcExtName := extension.GetFileExtName(c.SrcFilePath)
	var imageObj image.Image

	switch srcExtName {
	case constant.JPG, constant.JPEG:
		imageObj, err = jpeg.Decode(srcFile)
	case constant.PNG:
		imageObj, err = png.Decode(srcFile)
	case constant.GIF:
		imageObj, err = gif.Decode(srcFile)
	default:
		log.Fatalf("unknown format: %v", srcExtName)
		return
	}

	if err != nil {
		log.Fatalf("file decode failed: %v", err.Error())
		return
	}

	aimFile, err := os.Create(c.AimFilePath)
	if err != nil {
		log.Fatalf("file create failed: %v", err.Error())
		return
	}
	defer aimFile.Close()

	aimExtName := extension.GetFileExtName(c.AimFilePath)

	switch aimExtName {
	case constant.PNG:
		err = png.Encode(aimFile, imageObj)
	case constant.GIF:
		err = gif.Encode(aimFile, imageObj, nil)
	default:
		err = jpeg.Encode(aimFile, imageObj, &jpeg.Options{
			Quality: constant.JPEGMaxQuality,
		})
	}
	if err != nil {
		log.Fatalf("file encode failed: %v", err.Error())
		return
	}

	log.Println("the image format has been successfully converted")
}
