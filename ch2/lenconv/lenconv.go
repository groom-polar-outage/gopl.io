package lenconv

import "fmt"

type Feet float64
type Meter float64

const FeetInMeter = 3.28084

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
