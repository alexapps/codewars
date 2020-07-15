package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func main() {
	fmt.Println("Distance between 3 point by plane")
	coordinates := []Point{{1, 1}, {3, 2}, {4, 4}}
	var totalDistance float64
	for i := range coordinates {

		if i+1 < len(coordinates) {
			totalDistance = totalDistance + distanceBetween2Points(coordinates[i], coordinates[i+1])
		}
	}
	fmt.Println(fmt.Sprintf("Distance between pints %v is %v", coordinates, totalDistance))
}

// Simple 2D distance between 2 points (x1, y1) and (x2, y2)
func distanceBetween2Points(a, b Point) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}
