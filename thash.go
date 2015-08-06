package thash

import (
	"math"
)

// generates tile-hash by ZXY
func ZXYtoHash(z, x, y int) int64 {
	var hash int64
	positions := make(map[int][2]int)

	for currentZoom := 1; currentZoom <= z; currentZoom++ {
		tileSize := int(math.Pow(2, float64(z-currentZoom)))

		centerX, centerY := tileSize-1, tileSize-1
		for i := (currentZoom - 1); i > 0; i-- {
			n := int(math.Pow(2, float64(z-i)))

			centerX += positions[i][0] * n
			centerY += positions[i][1] * n
		}

		positionX, positionY := 0, 0
		if x > centerX {
			positionX = 1
		}
		if y > centerY {
			positionY = 1
		}

		positions[currentZoom] = [2]int{positionX, positionY}

		p := positions[currentZoom]
		hash += int64(int(math.Pow10(z-currentZoom)) * (2*p[1] + p[0] + 1))
	}

	return hash
}

// calculates ZXY by tile-hash
func HashtoZXY(hash int64) (z, x, y int) {
	z = MaxZoom(hash)

	for currentZoom := 1; currentZoom <= z; currentZoom++ {
		tile := getDigit(hash, currentZoom)
		tileSize := int(math.Pow(2, float64(z-currentZoom)))

		positionX := (tile - 1) % 2
		positionY := (tile - 1) / 2

		x += tileSize * positionX
		y += tileSize * positionY
	}

	return
}

// returns coordinates of central point of tile by tile-hash
func CentralPoint(hash int64) [2]float32 {
	z, x, y := HashtoZXY(hash)

	n := math.Pow(2, float64(z))

	lon_deg := ((float64(x)+0.5)/n)*360.0 - 180.0
	lat_rad := math.Atan(math.Sinh(math.Pi * (1 - (2*(float64(y)+0.5))/n)))
	lat_deg := lat_rad * (180.0 / math.Pi)

	return [2]float32{float32(lat_deg), float32(lon_deg)}
}

// returns max zoom in the tile-hash
func MaxZoom(hash int64) int {
	maxZoom := 0
	for hash != 0 {
		hash = hash / 10
		maxZoom++
	}

	return maxZoom
}

func getDigit(h int64, r int) int {
	max := MaxZoom(h)

	p1 := int64(math.Pow10(max - r))
	p2 := int64(math.Pow10(max - r + 1))

	r1 := int(h / p1)
	r2 := int(h/p2) * 10

	if r2 == 0 {
		return r1
	}

	return r1 - r2
}
