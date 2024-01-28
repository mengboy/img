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
	QualityCMD = &cobra.Command{
		Use:   "quality",
		Short: "quality ",
		Run:   quality,
	}
	qSourcePath  string // 图片路径或者图片文件夹路径
	qDstPath     string // 压缩后文件路径
	qualityValue int    // 压缩比例 小于1
)

func init() {
	QualityCMD.PersistentFlags().StringVarP(&qSourcePath, "source path", "s", "", "源路径")
	QualityCMD.PersistentFlags().StringVarP(&qDstPath, "dst path", "d", "", "压缩后路径")
	QualityCMD.PersistentFlags().IntVarP(&qualityValue, "quality value factor", "q", 80, "压缩质量")
}

func quality(cmd *cobra.Command, args []string) {
	if qualityValue == 100 {
		return
	}
	isDir, err := IsDir(qSourcePath)
	if err != nil {
		log.Fatal("source path invalid", qSourcePath)
		return
	}
	if isDir {
		err = qualityByDir(qSourcePath, qDstPath)
	} else {
		err = qualityFile(qSourcePath, qDstPath)
	}
	if err != nil {
		log.Fatal("resize failed", err)
		return
	}
	return
}

func qualityFile(sourcePath string, dstPath string) error {
	isExist, _ := IsExist(dstPath)
	if isExist {
		return errors.New("dst path is exist file")
	}
	imageManager := compress.ImageHandle{}
	err := imageManager.QualityLower(sourcePath, dstPath, qualityValue)
	if err != nil {
		return err
	}
	return nil
}

func qualityByDir(sourcePath string, dstPath string) error {
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
		err = imageManager.QualityLower(filePath, dstFilePath, qualityValue)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}
