package slicelist

import (
	"fmt"
)

// Implementar la lista sobre slices

type slicelist []interface {
}

func (s *slicelist) Add(item interface{}) {
	*s = append(*s, item)
}

func (s *slicelist) Remove(valor int) {
	*s = append((*s)[:valor], (*s)[:valor+1]...)
}

func (s slicelist) Get(valor, posicion int) interface{} {
	if posicion < 0 {
		panic("la posicion no puede ser 0")
	}

	return s[valor]
}

func (s slicelist) Len() int {
	return len(s)
}

func main() {
	// Crear una nueva lista
	myList := slicelist{}

	// Agregar elementos a la lista
	myList.Add(1)
	myList.Add("hello")
	myList.Add(3.14)

	// Acceder a un elemento de la lista
	fmt.Println(myList.Get(1)) // Output: hello

	// Eliminar un elemento de la lista
	myList.Remove(0)

	// Obtener la longitud de la lista
	fmt.Println(myList.Len()) // Output: 2
}
