# Go Legofy

[![Build Status](https://travis-ci.org/winhtaikaung/go-legofy.svg?branch=master)](https://travis-ci.org/winhtaikaung/go-legofy)

[![Go Report Card](https://goreportcard.com/badge/github.com/winhtaikaung/go-legofy)](https://goreportcard.com/report/github.com/winhtaikaung/go-legofy)

[![GitHub issues](https://img.shields.io/github/issues/winhtaikaung/go-legofy.svg)](https://github.com/winhtaikaung/go-legofy/issues)
[![GitHub forks](https://img.shields.io/github/forks/winhtaikaung/go-legofy.svg)](https://github.com/winhtaikaung/go-legofy/network)
[![GitHub stars](https://img.shields.io/github/stars/winhtaikaung/go-legofy.svg)](https://github.com/winhtaikaung/go-legofy/stargazers)
[![GitHub license](https://img.shields.io/github/license/winhtaikaung/go-legofy.svg)](https://github.com/winhtaikaung/go-legofy/blob/master/LICENSE.md)

### What is it?

Go Legofy is a Go lang utility library that takes a static image and makes it so that it looks as if it was built out of LEGO.

### Before

<img alt="Before" title="Before (The inflorescence of Zoysia grass, a variety of lawn grass. Picture by Hari Krishnan)" height="500" src="gopher.png?raw=true">

### After

<img alt="After" title="After" height="500" src="lego_with_img.png?raw=true">

## Quickstart

```shell
$ go get github.com/winhtaikaung/go-legofy
```

## Usage

```go
    //Legofy from image path
import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"

	"github.com/winhtaikaung/go-legofy/legofy"
)

func main() {
	fmt.Println("Lego My lovely Lego")
	//Legofy from image path
	sourceImagePath := "gopher.png"
	imgChanel := make(chan *legofy.LegoImage)
	go legofy.LegofyImagePath(sourceImagePath, 50, imgChanel)
	img := <-imgChanel
	close(imgChanel)
	fmt.Println(img.BrickCount)
	legofy.SaveAsJPEG("lego_with_path", img.Image, 100)

	// legofy from image.Image data type
	source, _ := os.Open(sourceImagePath)
	defer source.Close()
	sourceImg, _, _ := image.Decode(source) // Image Struct
	imgChanel = make(chan *legofy.LegoImage)
	go legofy.LegofyImage(sourceImg, 50, imgChanel)
	img = <-imgChanel
	close(imgChanel)
	legofy.SaveAsPNG("lego_with_img", img.Image, png.BestCompression)

}
```
## Docs

https://godoc.org/github.com/winhtaikaung/go-legofy/legofy

## Bugs

If you find a bug:

1. Check in the [open issues](https://github.com/winhtaikaung/go-legofy/issues) if the bug already exists.
2. If the bug is not there, create a [new issue](https://github.com/winhtaikaung/go-legofy/issues/new) with clear steps on how to reproduce it.

# ToDo

- [ ] Image Palette

## Contributing

1. Fork it ( https://github.com/winhtaikaung/go-legofy )

2) Create your feature branch (`git checkout -b my-new-feature`)

3. Commit your changes (`git commit -am 'Add some feature'`)

4) Push to the branch (`git push origin my-new-feature`)

5. Create a new Pull Request

## License

MIT

[![Twitter](https://img.shields.io/twitter/url/https/github.com/winhtaikaung/go-legofy.svg?style=social)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fwinhtaikaung%2Fgo-legofy)
