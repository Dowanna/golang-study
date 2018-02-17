package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit((c - 32) * 5 / 9)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f * 9 / 5) + 32)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
