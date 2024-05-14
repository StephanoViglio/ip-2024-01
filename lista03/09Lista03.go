package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	var xd, yd, zd, somaDif float64

	var d [99]float64
	var x [99]float64
	var y [99]float64
	var z [99]float64

	fmt.Scanln(&n)

	if n >= 2 && n <= 1000 {
		for i := 0; i < n; i++ {
			fmt.Scanln(&x[i], &y[i], &z[i])

			if x[i] < -1000 && x[i] > 1000 {
				break
			}
			if y[i] < -1000 && y[i] > 1000 {
				break
			}
			if x[i] < -1000 && x[i] > 1000 {
				break
			}
		}

		for j := 1; j <= n-1; j++ {
			xd = (x[j] - x[j-1])
			yd = (y[j] - y[j-1])
			zd = (z[j] - z[j-1])
			somaDif = math.Pow(xd, 2) + math.Pow(yd, 2) + math.Pow(zd, 2)
			d[j] = math.Sqrt(somaDif)
		}

		for k := 1; k <= n-1; k++ {
			fmt.Printf("%.2f\n", d[k])
		}
	}

}
