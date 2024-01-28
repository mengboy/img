package compress

import (
	"github.com/mengboy/img/constant"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"log"
	"os"
)

type PNG struct {
	ic ImageConvert
}

func (p *PNG) Resizer(filePath string, dstPath string, compressFactor float64) error {
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
	err = png.Encode(outputFile, resizedImg)
	if err != nil {
		return err
	}
	return nil
}

// QualityLower png 忽略quality，直接使用最佳压缩算法
func (p *PNG) QualityLower(filePath string, dstPath string, quality int) error {
	//if quality < 0 {
	//	return nil
	//}
	//if quality < 10 {
	//	quality = quality * 10
	//}
	//if quality > 100 {
	//	return nil
	//}
	// 打开PNG图片文件
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 解码PNG图片
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

	encoder := png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	// 保存为PNG格式
	err = encoder.Encode(outputFile, img)
	if err != nil {
		return err
	}

	return nil
}

func (p *PNG) Converter(filePath string, dstPath string, to string, quality int) error {
	switch to {
	case constant.JPEGType, constant.JPGType:
		log.Println(filePath, dstPath, to, quality)
		return p.ic.ToJPEG(filePath, dstPath, quality)
	default:
		return nil
	}
}
