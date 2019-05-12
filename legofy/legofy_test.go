package legofy

import (
	"fmt"
	"testing"
)

func TestGetLegoPallette(t *testing.T) {
	l := new(legofy)
	l.getLegoPalette("all")
	
}

func TestApplyThumbnNailEffects(t *testing.T) {
	fmt.Println("Test ApplyThumbnail Pallette")
}

func TestOverLayEffect(t *testing.T) {
	fmt.Println("Test OverLayEffect")
	l := new(legofy)
	effect := l.overLayeffect(0, 84)
	if effect != -16 {
		t.Log("Should equal with -16")
		t.Fail()
	}

	effect = l.overLayeffect(234, 84)
	if effect != 184 {
		t.Log("Should equal with 184")
		t.Fail()
	}

	effect = l.overLayeffect(52, 84)

	if effect != 3 {
		t.Log("Should equal with 3")
		t.Fail()
	}
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
