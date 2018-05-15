// Package tempconv performs Celsius and Fahrenheit computations
package tempconv

import "fmt"

// named types provide a way to separate different, potentially incompatible, uses of the underlying type
type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
)

// String methods for pretty printing
func (c Celsius) String() string    { return fmt.Sprintf("%g C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }
