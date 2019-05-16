# Go Legofy

### What is it?

Go Legofy is a Go lang utility library that takes a static image and makes it so that it looks as if it was built out of LEGO.

### Before

<img alt="Before" title="Before (The inflorescence of Zoysia grass, a variety of lawn grass. Picture by Hari Krishnan)" height="500" src="gopher.png?raw=true">
### After 
<img alt="After" title="After" height="500" src="lego_with_img.png?raw=true">

### Bugs

If you find a bug:

1. Check in the [open issues](https://github.com/winhtaikaung/go-legofy/issues) if the bug already exists.
2. If the bug is not there, create a [new issue](https://github.com/winhtaikaung/go-legofy/issues/new) with clear steps on how to reproduce it.

### Quickstart

```shell
$ go get github.com/winhtaikaung/go-legofy
```

```go
    //Legofy from image path
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
