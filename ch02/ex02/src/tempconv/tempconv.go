package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Meter float64
type Feet float64
type Pound float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%gC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (ft Feet) String() string      { return fmt.Sprintf("%gft", ft) }
func (lb Pound) String() string     { return fmt.Sprintf("%glb", lb) }
func (kg Kilogram) String() string  { return fmt.Sprintf("%gkg", kg) }
