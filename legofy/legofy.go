package legofy

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"math"
	"os"

	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/anthonynsimon/bild/blend"
)

type Legofy struct {
}

type LegoImage struct {
	Image      image.Image
	BrickCount int
}

func (l *Legofy) applyColorOverlay(brickImg image.Image, brickColor color.Color, pixel int) *image.RGBA {

	overlayR, overlayG, overlayB, overlayA := brickColor.RGBA()
	brickY := brickImg.Bounds().Max.Y
	brickX := brickImg.Bounds().Max.X

	cimg := image.NewRGBA(brickImg.Bounds())
	draw.Draw(cimg, brickImg.Bounds(), brickImg, image.Point{}, draw.Src)
	for y := 0; y < brickY; y++ {
		for x := 0; x < brickX; x++ {
			cimg.Set(x, y, color.RGBA{l.overLayeffect(uint8(overlayR)), l.overLayeffect(uint8(overlayG)), l.overLayeffect(uint8(overlayB)), uint8(overlayA)})
		}
	}

	return blend.Overlay(cimg, brickImg)
}

func (l *Legofy) overLayeffect(color uint8) uint8 {
	if color <= 33 {
		return 33
	} else if color >= 233 {
		return 233
	} else {
		return color
	}
}

func (l *Legofy) makeLegoImage(baseImg image.Image, brickImg image.Image, legChanel chan *LegoImage) {
	//To implement legofy process
	baseW, baseH := baseImg.Bounds().Max.X, baseImg.Bounds().Max.Y
	brickW, brickH := brickImg.Bounds().Max.X, brickImg.Bounds().Max.Y

	upLeft := image.Point{0, 0}
	lowRight := image.Point{baseW * brickW, baseH * brickH}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	// filling white
	// for y := 0; y < lowRight.Y; y++ {
	// 	for x := 0; x < lowRight.X; x++ {
	// 		img.Set(x, y, color.RGBA{255, 255, 255, 255})
	// 	}
	// }
	// cimg := image.NewRGBA(brickImg.Bounds())
	var i = 0
	for brickX := 0; brickX < baseW; brickX++ {
		for brickY := 0; brickY < baseH; brickY++ {
			color := baseImg.At(brickX, brickY)

			// draw.Draw(img, brickImg.Bounds(), l.applyColorOverlay(brickImg, color), image.Point{, brickY * brickH}, draw.Over)
			draw.Draw(img, img.Bounds(), l.applyColorOverlay(brickImg, color, brickY), image.Point{-(brickX * brickW), -(brickY * brickH)}, draw.Src)
			// apply color overlay Here
			// fmt.Println(brickX*brickW, brickY*brickH)
			i++
		}
	}
	legChanel <- &LegoImage{img, i}

}

func SaveAsJPEG(name string, img image.Image, quality int) {

	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var opt jpeg.Options
	if opt.Quality = 100; quality <= 100 {
		opt.Quality = quality
	}
	// ok, write out the data into the new JPEG file
	// TODO file write should be able to selec

	err = jpeg.Encode(f, img, &opt) // put quality to 80%
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
func SaveAsPNG(name string, img image.Image) {

	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

}

func (l *Legofy) getNewFileName() {

}

func (l *Legofy) getLegoPalette(paletteMode string) []float64 {
	p := new(palettes)
	legos := p.legos()
	palette := legos[paletteMode]
	data := palette.([]float64)

	return p.extendPalette(data, 0, 0)

}

func (l *Legofy) applyThumbNailEffect(baseImage image.Image, palettes []float64, dither bool) {

	paletteImage := image.NewRGBA(image.Rect(0, 0, 1, 1))
	fmt.Println(paletteImage)
}

func LegofyImage(imgsrc string, brickImg image.Image, brickSize int, palette string, dither bool, legoChan chan *LegoImage) {

	l := new(Legofy)
	sourceImg := l.readImage(imgsrc)
	newsizeX, newSizeY := l.getNewSize(sourceImg, brickImg, brickSize)
	fmt.Println(newsizeX, newSizeY)
	thumbImg := image.NewRGBA(image.Rect(0, 0, newsizeX, newSizeY))
	graphics.Thumbnail(thumbImg, sourceImg)

	// Check Palette mode in Future
	makerChan := make(chan *LegoImage)
	go l.makeLegoImage(thumbImg, brickImg, makerChan)
	legolizedImg := <-makerChan
	close(makerChan)
	legoChan <- legolizedImg

}

func (l *Legofy) readImage(path string) image.Image {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	img, _, err := image.Decode(file) // Image Struct
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", path, err)
	}

	return img
}

func (l *Legofy) getNewSize(baseImage image.Image, brickImg image.Image, size int) (int, int) {
	newImageSize := baseImage.Bounds()
	brickSize := brickImg.Bounds()
	scaleX, scaleY := 0, 0
	if size != 0 {
		scaleX, scaleY = size, size
	} else {
		scaleX, scaleY = brickSize.Max.X, brickSize.Max.Y
	}

	if newImageSize.Max.X > scaleX || newImageSize.Max.Y > scaleY {
		scale := 0.00
		if newImageSize.Max.X < newImageSize.Max.Y {
			scale = float64(newImageSize.Max.Y) / float64(scaleY)
		} else {
			scale = float64(newImageSize.Max.X) / float64(scaleX)
		}

		return int(math.Floor(float64(newImageSize.Max.X) / float64(scale))), int(math.Floor(float64(newImageSize.Max.Y) / float64(scale)))
	}
	return newImageSize.Max.X, newImageSize.Max.Y
}
