package legofy

import (
	"fmt"
	"testing"
)

func TestOverLayEffect(t *testing.T) {
	fmt.Println("Test OverLayEffect")
	l := new(Legofy)
	effect := l.overLayeffect(255)
	if effect != 233 {
		t.Log("Should equal with -16")
		t.Fail()
	}
	effect = l.overLayeffect(0)
	if effect != 33 {
		t.Log("Should equal with 33")
		t.Fail()
	}

}

func TestGetNewSize(t *testing.T) {
	fmt.Println("Test GetNewSize")
	l := new(Legofy)
	brickImg := l.readImage("../assets/1x1.png")
	baseImg := l.readImage("../assets/flower.jpg")
	x, y := l.getNewSize(baseImg, brickImg, 30)

	if x != 19 {
		t.Log("X Should match with 19")
		t.Fail()
	}
	if y != 29 {
		t.Log("Y Should match with 29")
		t.Fail()
	}

	// x, y = l.getNewSize(baseImg, brickImg, 200)
	// if x != 200 {
	// 	t.Log("Should match with 200")
	// 	t.Fail()
	// }
	// if y != 150 {
	// 	t.Log("Y Should match with 150")
	// 	t.Fail()
	// }
}
