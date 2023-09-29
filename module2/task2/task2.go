package task2

import (
	"awesomeProject/utility"
	"fmt"
	"math"
)

func Task() {
	utility.AnnounceNewTask(2, 2)
	playground()
	radiusAndSquare()
}

func playground() {
	var A *int
	var B = 25
	A = &B
	fmt.Printf("The pointer on B (%d) has the adress 0x%x and the value %d\n", B, A, *A)

	*A = 50
	fmt.Printf("The pointer on B (%d) has the adress 0x%x and the value %d\n\n", B, A, *A)
}

func radiusAndSquare() {
	length := 35.
	R := new(float64)
	*R = length / (2 * math.Pi)
	fmt.Printf("Round length %.2f = 2 * pi * radius %.2f\n", math.Round(length*100)/100, math.Round(*R*100)/100)
	rSquared := math.Pow(*R, 2)
	square := math.Pi * rSquared
	fmt.Printf("Round square %.2f = pi * r^2 %.2f\n", math.Round(square*100)/100, math.Round(rSquared*100)/100)
}
