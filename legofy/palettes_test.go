package legofy

import (
	"encoding/json"
	"testing"
)

func TestExtendPalette(t *testing.T) {
	numbers := make([]int, 0)
	numbers = append(numbers, 0)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	p := new(palettes)

	if len(p.extendPalette(numbers, 256, 3)) != 768 {
		t.Log("Should match with Zawgyi string")
		t.Fail()
	}

}

func TestMergePalette(t *testing.T) {
	const LEGOS string = `{"solid":{"102":[72,140,198],"106":[231,100,25],"119":[149,185,12],"124":[156,1,198],"135":[95,117,140],"138":[141,117,83],"140":[1,38,66],"141":[1,53,23],"151":[96,130,102],"154":[128,9,28],"191":[244,155,1],"192":[92,29,13],"194":[156,146,145],"199":[77,94,87],"208":[228,228,218],"212":[135,192,234],"221":[222,56,139],"222":[238,157,195],"226":[255,255,153],"268":[45,22,120],"283":[245,193,137],"308":[49,16,7],"312":[170,126,86],"024":[254,196,1],"021":[222,1,14],"023":[1,88,168],"028":[1,124,41],"018":[214,115,65],"001":[244,244,244],"026":[2,2,2],"037":[1,150,37],"005":[217,187,124],"038":[168,62,22]},"transparent":{"111":[166,145,130],"113":[238,157,195],"126":[156,149,199],"143":[206,227,246],"182":[236,118,14],"311":[153,255,102],"044":[249,239,105],"047":[231,102,72],"041":[224,42,41],"042":[182,224,234],"043":[80,177,232],"048":[99,178,110],"049":[241,237,91],"040":[238,238,238]},"effects":{"131":[141,148,150],"148":[73,63,59],"294":[254,252,213],"297":[170,127,46]},"mono":{"001":[244,244,244],"026":[2,2,2]}}`
	var result map[string]interface{}
	err := json.Unmarshal([]byte(LEGOS), &result)
	if err != nil {
		panic(err.Error())
	}
	p := new(palettes)

	if len(result) != 4 {
		t.Log("Should have initial Legos Palettes")
		t.Fail()
	}

	if len(p.mergePalettes(LEGOS)) < 5 {
		t.Log("Should Include all Palettes")
		t.Fail()
	}
}
