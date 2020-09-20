package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"log"
	"os"

	"golang.org/x/image/draw"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}

	fileName := flag.Arg(0)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	img, t, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Type of image:", t)

	rct := img.Bounds()
	fmt.Println("Width:", rct.Dx())
	fmt.Println("Height:", rct.Dy())

	rate := rct.Dx() / 400

	imgDst := image.NewRGBA(image.Rect(0, 0, rct.Dx()/rate, rct.Dy()/rate))
	draw.CatmullRom.Scale(imgDst, imgDst.Bounds(), img, rct, draw.Over, nil)

	//create resized image file
	fileNameResize := "re" + fileName
	dst, err := os.Create(fileNameResize) //maybe dst file path
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer dst.Close()

	if err := jpeg.Encode(dst, imgDst, &jpeg.Options{Quality: 100}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
