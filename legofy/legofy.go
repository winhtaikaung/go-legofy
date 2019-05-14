package legofy

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"math"
	"os"

	"github.com/BurntSushi/graphics-go/graphics"
)

type legofy struct {
}

func (l *legofy) applyColorOverlay() {

}

func (l *legofy) overLayeffect(color int, overlay int) int {
	if color < 33 {
		return overlay - 100
	} else if color > 233 {
		return overlay + 100
	} else {
		return overlay - 133 + color
	}
}

func (l *legofy) makeLegoImage(baseImg image.Image, brickImg image.Image) {
	//To implement legofy process
	baseW, baseH := baseImg.Bounds().Max.X, baseImg.Bounds().Max.Y
	brickW, brickH := brickImg.Bounds().Max.X, brickImg.Bounds().Max.Y

	upLeft := image.Point{0, 0}
	lowRight := image.Point{baseW * brickW, baseH * brickH}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	//filling white
	for y := 0; y < lowRight.Y; y++ {
		for x := 0; x < lowRight.X; x++ {
			img.Set(x, y, color.RGBA{180, 180, 250, 255})
		}
	}

	for brickX := 0; brickX < baseW; brickX++ {
		for brickY := 0; brickY < baseH; brickY++ {
			baseImg.At(brickX, brickY).RGBA()
			// apply color overlay Here
		}
	}

	l.generateSample("graphic_lego.png", baseImg)

}

func (l *legofy) generateSample(name string, img image.Image) {

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

func (l *legofy) getNewFileName() {

}

func (l *legofy) getLegoPalette(paletteMode string) []float64 {
	p := new(palettes)
	legos := p.legos()
	palette := legos[paletteMode]
	data := palette.([]float64)

	return p.extendPalette(data, 0, 0)

}

func (l *legofy) applyThumbNailEffect(baseImage image.Image, palettes []float64, dither bool) {

	paletteImage := image.NewRGBA(image.Rect(0, 0, 1, 1))
	fmt.Println(paletteImage)
}

func LegofyImage(sourceImg image.Image, brickImg image.Image, brickSize int, palette string, dither bool) {
	l := new(legofy)
	newsizeX, newSizeY := l.getNewSize(sourceImg, brickImg, brickSize)
	fmt.Println(newsizeX, newSizeY)
	thumbImg := image.NewRGBA(image.Rect(0, 0, newsizeX, newSizeY))
	graphics.Thumbnail(thumbImg, sourceImg)

	// Check Palette mode in Future
	// l.makeLegoImage(sourceImg, brickImg)
	l.makeLegoImage(thumbImg, brickImg)

}

func (l *legofy) readImage(path string) image.Image {
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

func (l *legofy) getNewSize(baseImage image.Image, brickImg image.Image, size int) (int, int) {
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
