package tempconv

const (
	meterToFt = 3.28084
	kgToPound = 2.20462
)

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FToK(f Fahrenheit) Kelvin {
	return CToK(Celsius((f * 9 / 5) + 32))
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

func MToFt(m Meter) Feet {
	return Feet(m * meterToFt)
}

func FtToM(f Feet) Meter {
	return Meter(f / meterToFt)
}

func KgToLb(kg Kilogram) Pound {
	return Pound(kg * kgToPound)
}

func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb / kgToPound)
}
