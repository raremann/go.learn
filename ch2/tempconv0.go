// Package tempconv performs Celsius and Fahrenheit computations
package tempconv

import "fmt"

// named types provide a way to separate different, potentially incompatible, uses of the underlying type
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

// Explicit type conversion are required from the underlying type (or any other compatible named type)
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(c Fahrenheit) Celsius { return Celsius(5/9*(f - 32) }