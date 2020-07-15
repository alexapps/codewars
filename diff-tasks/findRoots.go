package main
import (
	"fmt"
	"math"
)

func findRoots(a, b, c float64) (float64, float64) {
	d := b*b - 4*a*c
	if d < 0.0 {
		return 0,0
	} else if d == 0.0 {
		return -1 * b / 2 * a, 0
	} else if d > 0.0 {
		x1 := (-1 * b - math.Sqrt(d))/ (2 * a)
		x2 := (-1 * b + math.Sqrt(d))/ (2 * a)
		return x1, x2
	}
    return 0,0
}
    
func main() {
    x1, x2 := findRoots(2, 10, 8)
    fmt.Printf("Roots: %f, %f", x1, x2)
}