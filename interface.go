package figuras

import "fmt"

type Figura interface {
	area() float64
	perimetro() float64
}

func Medidas(f Figura) {
	fmt.Println("Medidas", f)
	fmt.Println("Area", f.area())
	fmt.Println("Perimetro", f.perimetro())
}
