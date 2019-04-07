package legofy

import (
	"fmt"
	"testing"
)

func TestGetLegoPallette(t *testing.T) {
	fmt.Println("TestGel Logo Pallette")
}

func TestApplyThumbnNailEffects(t *testing.T) {
	fmt.Println("Test ApplyThumbnail Pallette")
}

func TestGetNewSize(t *testing.T) {
	fmt.Println("Test GetNewSize")
	l := new(legofy)
	brickImg := l.readImage("../assets/1x1.png")
	baseImg := l.readImage("graphic.jpg")
	x, y := l.getNewSize(baseImg, brickImg, 30)
	if x != 30 {
		t.Log("Should match with 30")
		t.Fail()
	}
	if y != 22 {
		t.Log("Y Should match with 22")
		t.Fail()
	}

	x, y = l.getNewSize(baseImg, brickImg, 200)
	if x != 200 {
		t.Log("Should match with 200")
		t.Fail()
	}
	if y != 150 {
		t.Log("Y Should match with 150")
		t.Fail()
	}
}
func TestLegofyImage(*testing.T) {
	fmt.Println("TestLegofyImage")
}
