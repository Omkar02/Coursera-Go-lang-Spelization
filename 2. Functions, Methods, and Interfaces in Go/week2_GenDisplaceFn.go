package main

import (
	"fmt"
)

// s =Â½ a t2 + vot + so

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {

	return func(t float64) float64 {
		return ((0.5 * a * (t * t)) + (vo * t) + so)
	}
}

func main() {

	param := [4]float64{}

	fmt.Printf(" %s\n", "Please enter acceleration ...  ")
	fmt.Scan(&param[0])
	fmt.Printf(" %s\n", "Please enter initial velocity ...  ")
	fmt.Scan(&param[1])
	fmt.Printf(" %s\n", "Please enter initial displacement ...  ")
	fmt.Scan(&param[2])
	fmt.Printf(" %s\n", "Please enter time (seconds) ...  ")
	fmt.Scan(&param[3])

	fn := GenDisplaceFn(param[0], param[1], param[2])

	fmt.Printf("\n%s%v%s  ...   %v\n\n", "Displacement after ", param[3], " seconds", fn(param[3]))

}