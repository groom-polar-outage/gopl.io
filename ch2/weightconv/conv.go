package weightconv

func PToK(p Pound) Kilogram { return Kilogram(p / PoundInKilogram) }

func KToP(k Kilogram) Pound { return Pound(k * PoundInKilogram) }
