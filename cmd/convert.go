package cmd

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mengboy/img/compress"
	"github.com/spf13/cobra"
)

var (
	ConvertCMD = &cobra.Command{
		Use:   "convert",
		Short: "convert ",
		Run:   convert,
	}
	cSourcePath   string // 图片路径或者图片文件夹路径
	cDstPath      string // 压缩后文件路径
	convertTo     string
	cQualityValue int // 压缩比例 小于1
)

func init() {
	ConvertCMD.PersistentFlags().StringVarP(&cSourcePath, "source path", "s", "", "源路径")
	ConvertCMD.PersistentFlags().StringVarP(&cDstPath, "dst path", "d", "", "压缩后路径")
	ConvertCMD.PersistentFlags().StringVarP(&convertTo, "convert type", "c", "", "转换为jpg、png")
	ConvertCMD.PersistentFlags().IntVarP(&cQualityValue, "quality value factor", "q", 80, "压缩质量")
}

func convert(cmd *cobra.Command, args []string) {
	isDir, err := IsDir(cSourcePath)
	if err != nil {
		log.Fatal("source path invalid", cSourcePath)
		return
	}
	if isDir {
		err = convertByDir(cSourcePath, cDstPath)
	} else {
		err = convertFile(cSourcePath, cDstPath)
	}
	if err != nil {
		log.Fatal("convert failed", err)
		return
	}
	return
}

func convertFile(sourcePath string, dstPath string) error {
	isExist, _ := IsExist(dstPath)
	if isExist {
		return errors.New("dst path is exist file")
	}
	imageManager := compress.ImageHandle{}
	log.Println(sourcePath, dstPath, convertTo, cQualityValue)
	err := imageManager.Converter(sourcePath, dstPath, convertTo, cQualityValue)
	if err != nil {
		return err
	}
	return nil
}

func convertByDir(sourcePath string, dstDir string) error {
	isExist, err := IsExist(dstDir)
	if !isExist {
		// 创建目录
		err = os.Mkdir(dstDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	dstIsDir, err := IsDir(dstDir)
	if err != nil {
		return err
	}
	if !dstIsDir {
		return err
	}
	if !strings.HasSuffix(dstDir, "/") {
		dstDir += "/"
	}
	imageManager := compress.ImageHandle{}
	filepath.Walk(sourcePath, func(filePath string, info fs.FileInfo, err error) error {
		// 不处理子目录
		if info.IsDir() {
			return nil
		}
		info.Name()
		fName := GetFileNameWithoutExt(filePath)
		dstFilePath := dstDir + fName + "." + convertTo
		err = imageManager.Converter(filePath, dstFilePath, convertTo, cQualityValue)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}
