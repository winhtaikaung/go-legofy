package legofy

import (
	"encoding/json"
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
const LEGOS string = `{"solid": {"024": [254, 196, 1], "106": [231, 100, 25], "021": [222, 1, 14], "221": [222, 56, 139], "023": [1, 88, 168], "028": [1, 124, 41], "119": [149, 185, 12], "192": [92, 29, 13], "018": [214, 115, 65], "001": [244, 244, 244], "026": [2, 2, 2], "226": [255, 255, 153], "222": [238, 157, 195], "212": [135, 192, 234], "037": [1, 150, 37], "005": [217, 187, 124], "283": [245, 193, 137], "208": [228, 228, 218], "191": [244, 155, 1], "124": [156, 1, 198], "102": [72, 140, 198], "135": [95, 117, 140], "151": [96, 130, 102], "138": [141, 117, 83], "038": [168, 62, 22], "194": [156, 146, 145], "154": [128, 9, 28], "268": [45, 22, 120], "140": [1, 38, 66], "141": [1, 53, 23], "312": [170, 126, 86], "199": [77, 94, 87], "308": [49, 16, 7]}, "transparent": {"044": [249, 239, 105], "182": [236, 118, 14], "047": [231, 102, 72], "041": [224, 42, 41], "113": [238, 157, 195], "126": [156, 149, 199], "042": [182, 224, 234], "043": [80, 177, 232], "143": [206, 227, 246], "048": [99, 178, 110], "311": [153, 255, 102], "049": [241, 237, 91], "111": [166, 145, 130], "040": [238, 238, 238]}, "effects": {"131": [141, 148, 150], "297": [170, 127, 46], "148": [73, 63, 59], "294": [254, 252, 213]}, "mono": {"001": [244, 244, 244], "026": [2, 2, 2]}}`

// Primate member functions
func (p *palettes) extendPalette(palette []float64, colors int, rgb int) []float64 {
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

func (p *palettes) flattenPalettes(palettes string) map[string]interface{} {
	mergedPalettes := p.mergePalettes(palettes)
	unifiedPalettes := make(map[string]interface{})
	for key, val := range mergedPalettes {
		colList := make([]float64, 0)
		if pal, ok := val.(map[string]interface{}); ok {
			for _, colors := range pal {
				if colorL, ok := colors.([]interface{}); ok {
					for _, item := range colorL {
						colList = append(colList, item.(float64))
					}
				} else {
					log.Printf("record not a map[string]interface{} %s", colorL)
				}

			}
		} else {
			log.Printf("record not a map[string]interface{} %s", pal)
		}
		unifiedPalettes[key] = colList
	}
	return unifiedPalettes
}

func (p *palettes) legos() map[string]interface{} {
	return p.flattenPalettes(LEGOS)
}
