package compress

import (
	"github.com/mengboy/img/constant"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"os"
)

type JPG struct {
	ic ImageConvert
}

func (j *JPG) Resizer(filePath string, dstPath string, compressFactor float64) error {
	// 打开图片文件
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// 计算新的宽高
	width := uint(float64(img.Bounds().Dx()) * compressFactor)
	height := uint(float64(img.Bounds().Dy()) * compressFactor)

	// 缩放图片
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)
	// 创建输出文件
	outputFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 保存压缩后的图片
	err = jpeg.Encode(outputFile, resizedImg, nil)
	if err != nil {
		return err
	}
	return nil
}

func (j *JPG) QualityLower(filePath string, dstPath string, quality int) error {
	if quality < 0 {
		return nil
	}
	if quality < 10 {
		quality = quality * 10
	}
	if quality > 100 {
		return nil
	}
	// 打开图片文件
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// 创建输出文件
	outputFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 设置JPEG编码参数
	options := &jpeg.Options{
		Quality: int(quality),
	}
	// 保存压缩后的图片
	err = jpeg.Encode(outputFile, img, options)
	if err != nil {
		return err
	}

	return nil
}

func (j *JPG) Converter(filePath string, dstPath string, to string, quality int) error {
	switch to {
	case constant.PNGType:
		return j.ic.ToPNG(filePath, dstPath)
	default:
		return nil

	}

}
