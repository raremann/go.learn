package tempconv

// Explicit type conversion are required from the underlying type (or any other compatible named type)

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius(5.0 / 9.0 * (f - 32)) }

func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }
