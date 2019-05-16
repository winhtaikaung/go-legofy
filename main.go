package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"

	"./legofy"
)

func main() {
	fmt.Println("Lego My lovely Lego")
	//Legofy from image path
	sourceImagePath := "gopher.png"
	imgChanel := make(chan *legofy.LegoImage)
	go legofy.LegofyImagePath(sourceImagePath, 30, imgChanel)
	img := <-imgChanel
	close(imgChanel)
	fmt.Println(img.BrickCount)
	legofy.SaveAsJPEG("lego_with_path", img.Image, 100)

	// legofy from image.Image data type
	source, _ := os.Open(sourceImagePath)
	defer source.Close()
	sourceImg, _, _ := image.Decode(source) // Image Struct
	imgChanel = make(chan *legofy.LegoImage)
	go legofy.LegofyImage(sourceImg, 30, imgChanel)
	img = <-imgChanel
	close(imgChanel)
	legofy.SaveAsPNG("lego_with_img", img.Image, png.BestCompression)

}
