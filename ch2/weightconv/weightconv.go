package weightconv

import "fmt"

const PoundInKilogram = 2.20462

type Pound float64
type Kilogram float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

func (p Pound) String() string { return fmt.Sprintf("%gp", p) }
