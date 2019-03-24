package legofy

import (
	"encoding/json"
	"fmt"
	"log"
)

// Salutation : palettes
// Printer : palettes struct
// Greet : This module contains the `lego` palette mappings.
// CreateMessage : Color mapping source;
// - http://www.brickjournal.com/files/PDFs/2010LEGOcolorpalette.pdf
type palettes struct {
}

// Color Mapping Source
const LEGOS string = `{"solid":{"102":[72,140,198],"106":[231,100,25],"119":[149,185,12],"124":[156,1,198],"135":[95,117,140],"138":[141,117,83],"140":[1,38,66],"141":[1,53,23],"151":[96,130,102],"154":[128,9,28],"191":[244,155,1],"192":[92,29,13],"194":[156,146,145],"199":[77,94,87],"208":[228,228,218],"212":[135,192,234],"221":[222,56,139],"222":[238,157,195],"226":[255,255,153],"268":[45,22,120],"283":[245,193,137],"308":[49,16,7],"312":[170,126,86],"024":[254,196,1],"021":[222,1,14],"023":[1,88,168],"028":[1,124,41],"018":[214,115,65],"001":[244,244,244],"026":[2,2,2],"037":[1,150,37],"005":[217,187,124],"038":[168,62,22]},"transparent":{"111":[166,145,130],"113":[238,157,195],"126":[156,149,199],"143":[206,227,246],"182":[236,118,14],"311":[153,255,102],"044":[249,239,105],"047":[231,102,72],"041":[224,42,41],"042":[182,224,234],"043":[80,177,232],"048":[99,178,110],"049":[241,237,91],"040":[238,238,238]},"effects":{"131":[141,148,150],"148":[73,63,59],"294":[254,252,213],"297":[170,127,46]},"mono":{"001":[244,244,244],"026":[2,2,2]}}`

// Primate member functions
func (p *palettes) extendPalette(palette []int, colors int, rgb int) []int {
	if rgb == 0 {
		rgb = 3
	}

	if colors == 0 {
		colors = 256
	}

	missingColors := 255
	palette = append(palette)
	if missingColors > 0 {
		firstColor := palette[:rgb]
		for i := 0; i < missingColors; i++ {
			for j := 0; j < len(firstColor); j++ {
				palette = append(palette, firstColor[j])
			}
		}
	}
	return palette[:colors*rgb]

}

func (p *palettes) mergePalettes(palettes string) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(palettes), &result)
	if err != nil {
		panic(err.Error())
	}

	unified := make(map[string]interface{})
	for _, palettes := range result {
		if pal, ok := palettes.(map[string]interface{}); ok {
			for key, val := range pal {
				unified[key] = val
			}
		} else {
			log.Printf("record not a map[string]interface{} %s", pal)
		}
	}
	result["all"] = unified
	return result
}

func Palettes() {
	fmt.Println("hello")
}
