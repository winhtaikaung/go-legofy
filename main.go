package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"./legofy"
)

func main() {
	fmt.Println("Lego My lovely Lego")

	imagePath := "assets/1x1.png"
	sourceImagePath := "profile.jpg"
	brick, _ := os.Open(imagePath)
	defer brick.Close()
	brickImg, _, _ := image.Decode(brick) // Image Struct

	source, _ := os.Open(sourceImagePath)
	defer source.Close()
	sourceImg, _, _ := image.Decode(source) // Image Struct

	//Legofy with GoRoutine and channel
	imgChanel := make(chan *legofy.LegoImage)
	go legofy.AsyncLegofyImage(sourceImg, brickImg, 1000, "none", false, imgChanel)
	fmt.Println("Routine Async in Progress")
	img := <-imgChanel
	close(imgChanel)
	fmt.Println("Routine Async Done")
	legofy.SaveAsJPEG("graphic_lego.jpg", img.Image, 20)

	blkImg := legofy.LegofyImage(sourceImg, brickImg, 1000, "none", false)
	fmt.Println("Routine Sync in Progress")
	legofy.SaveAsJPEG("graphic_lego_sync.jpg", blkImg.Image, 20)
	fmt.Println("Routine Sync in Done")

	// image1, err := os.Open(sourceImagePath)
	// if err != nil {
	// 	log.Fatalf("failed to open: %s", err)
	// }

	// first, _, _ := image.Decode(image1)

	// defer image1.Close()

	// image2, err := os.Open(imagePath)
	// if err != nil {
	// 	log.Fatalf("failed to open: %s", err)
	// }
	// second, _, _ := image.Decode(image2)

	// defer image2.Close()

	// offset := image.Pt(300, 200)
	// b := first.Bounds()
	// image3 := image.NewRGBA(b)
	// draw.Draw(image3, b, first, image.ZP, draw.Src)
	// draw.Draw(image3, second.Bounds().Add(offset), second, image.ZP, draw.Over)

	// third, err := os.Create("result.png")
	// if err != nil {
	// 	log.Fatalf("failed to create: %s", err)
	// }
	// // jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})

	// defer third.Close()
	// err = png.Encode(third, image3)
	// if err != nil {
	// 	panic(err)
	// }
}
