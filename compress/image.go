package compress

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

type ImageConvert struct {
}

//func (i *ImageConvert) ToWebp(filePath string, dstPath string, quality int) error {
//	// 打开PNG图片文件
//	file, err := os.Open(filePath)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//	// 解码PNG图片
//	img, _, err := image.Decode(file)
//	if err != nil {
//		return err
//	}
//	// 创建输出文件
//	outputFile, err := os.Create(dstPath)
//	if err != nil {
//		return err
//	}
//	defer outputFile.Close()
//	// 设置WebP编码参数
//	options := webp.Options{
//		Lossless: false, // 是否使用无损压缩
//		Quality:  float32(quality),
//	}
//	// 保存为WebP格式
//	err = webp.Encode(outputFile, img, &options)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (i *ImageConvert) ToJPEG(filePath string, dstPath string, quality int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

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
	jpegOptions := &jpeg.Options{
		Quality: quality,
	}

	// 将PNG转为JPEG并进行压缩
	err = jpeg.Encode(outputFile, img, jpegOptions)
	if err != nil {
		return err
	}

	return nil
}

func (i *ImageConvert) ToPNG(filePath string, dstPath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

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

	pngEncoder := &png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	err = pngEncoder.Encode(outputFile, img)
	if err != nil {
		return err
	}

	return nil
}
