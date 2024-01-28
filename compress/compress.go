package compress

import (
	"errors"
	"github.com/mengboy/img/constant"
	"log"
)

type Image interface {
	Resizer(filePath string, dstPath string, compressFactor float64) error
	QualityLower(filePath string, dstPath string, quality int) error
	Converter(filePath string, dstPath string, to string, quality int) error
}

type ImageHandle struct {
}

func (i *ImageHandle) Resizer(filePath string, dstPath string, compressFactor float64) error {
	ext := GetExtWithoutPoint(filePath)
	switch ext {
	case constant.JPGType, constant.JPEGType:
		j := &JPG{}
		return j.Resizer(filePath, dstPath, compressFactor)
	case constant.PNGType:
		p := &PNG{}
		return p.Resizer(filePath, dstPath, compressFactor)
	default:
		return errors.New("not support image type: " + ext)
	}
}

func (i *ImageHandle) QualityLower(filePath string, dstPath string, quality int) error {
	ext := GetExtWithoutPoint(filePath)
	switch ext {
	case constant.JPGType, constant.JPEGType:
		j := &JPG{}
		return j.QualityLower(filePath, dstPath, quality)
	case constant.PNGType:
		p := &PNG{}
		return p.QualityLower(filePath, dstPath, quality)
	default:
		return errors.New("not support image type: " + ext)
	}
}

func (i *ImageHandle) Converter(filePath string, dstPath string, to string, quality int) error {
	ext := GetExtWithoutPoint(filePath)
	log.Println(filePath, dstPath, to, quality, ext)
	switch ext {
	case constant.JPGType, constant.JPEGType:
		j := &JPG{}
		return j.Converter(filePath, dstPath, to, quality)
	case constant.PNGType:
		p := &PNG{}
		return p.Converter(filePath, dstPath, to, quality)
	default:
		return errors.New("not support image type: " + ext)
	}
}
