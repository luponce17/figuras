package figuras

type Cuadrado struct {
	Ancho float64
	Alto  float64
}

func (cuadrado *Cuadrado) area() float64 {
	return cuadrado.Alto * cuadrado.Ancho
}

func (cuadrado *Cuadrado) perimetro() float64 {
	return (2 * cuadrado.Alto) + (2 * cuadrado.Ancho)
}
