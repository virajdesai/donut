package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var A, B float64 = 0, 0
	var i, j float64
	var k int

	var z [1760]float64
	var b [1760]string

	fmt.Print("\033[H\033[2J")

	for {
		for i := range b {
			b[i] = " "
		}

		for i := range z {
			z[i] = 0
		}

		for j = 0; j < 6.28; j += 0.07 {
			for i = 0; i < 6.28; i += 0.02 {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(A)
				f := math.Sin(j)
				g := math.Cos(A)
				h := d + 2
				D := 1 / (c*h*e + f*g + 5)
				l := math.Cos(i)
				m := math.Cos(B)
				n := math.Sin(B)
				t := c*h*g - f*e

				x := int(40 + 30*D*(l*h*m-t*n))
				y := int(12 + 15*D*(l*h*n+t*m))

				o := int(x + 80*y)

				N := int(8 * ((f*e-c*d*g)*m - c*d*e - f*g - l*d*n))

				if 22 > y && y > 0 && x > 0 && 80 > x && D > z[o] {
					z[o] = D
					if N > 0 {
						b[o] = string(".,-~:;=!*#$@"[N])
					} else {
						b[o] = "."
					}
				}
			}
		}

		print("\x1b[H")
		for k = 0; k <= 1760; k++ {
			if k%80 > 0 {
				fmt.Printf(b[k])
			} else {
				fmt.Printf("\n")
			}

			A += 0.00004
			B += 0.00002
		}

		time.Sleep(16 * time.Millisecond)
	}
}
