package lenconv

func FToM(f Feet) Meter { return Meter(f / FeetInMeter) }

func MToF(m Meter) Feet { return Feet(m * FeetInMeter) }
