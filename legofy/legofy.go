package legofy

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
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

func (l *legofy) makeLegoImage() {

}

func (l *legofy) getNewFileName() {

}

func (l *legofy) getLegoPalette() {

}

func (l *legofy) applyThumbNailEffect() {

}

func (l *legofy) legofyImage(sourceImg image.Image, brickImg image.Image, brickSize int, palette string, dither bool) {
	// newScanner(sourceImg)
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

		return int(math.Round(float64(newImageSize.Max.X) / float64(scale))), int(math.Round(float64(newImageSize.Max.Y) / float64(scale)))
	}
	return newImageSize.Max.X, newImageSize.Max.Y
}
