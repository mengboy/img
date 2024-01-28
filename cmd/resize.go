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
	CompressCMD = &cobra.Command{
		Use:   "resize",
		Short: "resize ",
		Run:   resize,
	}
	sourcePath        string  // 图片路径或者图片文件夹路径
	dstPath           string  // 压缩后文件路径
	compressionFactor float64 // 压缩比例 小于1
)

func init() {
	CompressCMD.PersistentFlags().StringVarP(&sourcePath, "source path", "s", "", "源路径")
	CompressCMD.PersistentFlags().StringVarP(&dstPath, "dst path", "d", "", "压缩后路径")
	CompressCMD.PersistentFlags().Float64VarP(&compressionFactor, "compress factor", "f", 1, "压缩比例")
}

func resize(cmd *cobra.Command, args []string) {
	if compressionFactor > 1 {
		return
	}
	isDir, err := IsDir(sourcePath)
	if err != nil {
		log.Fatal("source path invalid", sourcePath)
		return
	}
	if isDir {
		err = resizeByDir(sourcePath, dstPath)
	} else {
		err = resizeFile(sourcePath, dstPath)
	}
	if err != nil {
		log.Fatal("resize failed", err)
		return
	}
	return
}

func resizeFile(sourcePath string, dstPath string) error {
	isExist, _ := IsExist(dstPath)
	if isExist {
		return errors.New("dst path is exist file")
	}
	imageManager := compress.ImageHandle{}
	err := imageManager.Resizer(sourcePath, dstPath, compressionFactor)
	if err != nil {
		return err
	}
	return nil
}

func resizeByDir(sourcePath string, dstPath string) error {
	isExist, err := IsExist(dstPath)
	if !isExist {
		// 创建目录
		err = os.Mkdir(dstPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	dstIsDir, err := IsDir(dstPath)
	if err != nil {
		return err
	}
	if !dstIsDir {
		return err
	}
	if !strings.HasSuffix(dstPath, "/") {
		dstPath += "/"
	}
	imageManager := compress.ImageHandle{}
	filepath.Walk(sourcePath, func(filePath string, info fs.FileInfo, err error) error {
		// 不处理子目录
		if info.IsDir() {
			return nil
		}
		fName := info.Name()
		dstFilePath := dstPath + fName
		err = imageManager.Resizer(filePath, dstFilePath, compressionFactor)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}
