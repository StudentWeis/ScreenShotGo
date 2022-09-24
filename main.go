package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"

	"github.com/kbinani/screenshot"
)

var fileName = flag.String("f", "NoFileName", "File name.")
var dirName = flag.String("d", "./", "Directory name.")

func main() {
	flag.Parse()

	// 获取当前所有的活动屏幕，并编号
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		// 判断路径是否存在，若路径不存在则创建之
		if _, err := os.Stat(*dirName); err != nil {
			os.MkdirAll(*dirName, 0777)
		}
		fileName := fmt.Sprintf("%s%s_%d.jpg", *dirName, *fileName, i)
		file, _ := os.Create(fileName)
		defer file.Close()
		jpeg.Encode(file, img, &jpeg.Options{Quality: 90})
	}
}
