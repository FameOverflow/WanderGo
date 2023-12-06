package Config

import (
	"log"
	"strconv"
	"strings"
)

func parseInput(input string) []Place {
	lines := strings.Split(input, "\n")
	places := make([]Place, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.FieldsFunc(line, func(c rune) bool {
			return c == ' ' || c == ',' || c == '-' || c == ','
		})

		name := parts[0]
		topLeftLat, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Println(err)
			continue
		}
		topLeftLon, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			log.Println(err)
			continue
		}
		bottomRightLat, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			log.Println(err)
			continue
		}
		bottomRightLon, err := strconv.ParseFloat(parts[4], 64)
		if err != nil {
			log.Println(err)
			continue
		}
		location := Place{
			PlaceName: name,
			TopLeftPoint: Address{
				X: topLeftLat,
				Y: topLeftLon,
			},
			BottomRightPoint: Address{
				X: bottomRightLat,
				Y: bottomRightLon,
			},
			CenterPoint: Address{
				X: (topLeftLat + bottomRightLat) / 2,
				Y: (topLeftLon + bottomRightLon) / 2,
			},
		}
		places = append(places, location)
	}
	for i := range places {
		err := GLOBAL_DB.Model(&Place{}).Create(&places[i]).Error
		if err != nil {
			log.Println(err)
			return []Place{}
		}
	}
	return places
}
