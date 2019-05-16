package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	"./legofy"
)

func main() {
	fmt.Println("Lego My lovely Lego")

	// imagePath := "assets/1x1.png"
	sourceImagePath := "./legofy/flower.jpg"
	// brick, _ := os.Open(imagePath)
	// defer brick.Close()
	// brickImg, _, _ := image.Decode(brick) // Image Struct

	// source, _ := os.Open(sourceImagePath)
	// defer source.Close()
	// sourceImg, _, _ := image.Decode(source) // Image Struct

	//Legofy with GoRoutine and channel
	imgChanel := make(chan *legofy.LegoImage)
	go legofy.LegofyImage(sourceImagePath, 200, "none", false, imgChanel)
	fmt.Println("Routine Async in Progress")
	img := <-imgChanel
	close(imgChanel)
	fmt.Println("Routine Async Done")
	legofy.SaveAsJPEG("graphic_lego.jpg", img.Image, 80)

}
