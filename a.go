package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	// trainLabel()
	trainImage()
}

func createPng(buf []byte) {
	const width, height = 28, 28
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	i := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			i = r*width + c
			img.Set(c, r, color.NRGBA{
				R: buf[i],
				G: buf[i],
				B: buf[i],
				A: 255,
			})
		}
	}
	f, err := os.Create("image.png")
	ck(err)
	defer f.Close()

	err = png.Encode(f, img)
	ck(err)
}

func trainImage() {
	fmt.Println("\nTrain Image")
	fileName := "data/train-images-idx3-ubyte"
	buf, err := ioutil.ReadFile(fileName)
	ck(err)

	p("magic number", buf[:4])
	p("number of images", buf[4:8])
	p("number of rows", buf[8:12])
	p("number of columns", buf[12:16])
	/* for _, c := range buf[16:44] {
		p("pixel", []byte{c})
	} */

	imageIndex := 0

	const SIZE = 28 * 28
	start := 16 + imageIndex*SIZE
	end := start + SIZE
	createPng(buf[start:end])

	trainLabel(imageIndex)

}

func trainLabel(index int) {
	// fmt.Println("\nTrain Label")
	fileName := "data/train-labels-idx1-ubyte"
	buf, err := ioutil.ReadFile(fileName)
	ck(err)

	/* p("magic number", buf[:4])
		p("number of items", buf[4:8])
	  for i, c := range buf[8:20] {
	    p("label "+strconv.Itoa(i), []byte{c})
	  } */

	p("\nlabel "+strconv.Itoa(index), buf[8+index:9+index])

}

func p(tag string, buf []byte) {
	fmt.Printf("%20s: % x\n", tag, buf)
}

func ck(err error) {
	if err != nil {
		panic(err)
	}
}
