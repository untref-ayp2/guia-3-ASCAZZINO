package main

import "fmt"

type slicelist struct {
	items []int
}

func (s *slicelist) Append(value int) {
	s.items = append(s.items, value)
}

func (s *slicelist) InsertAt(posicion, value int) {
	if posicion < 0 || posicion > len(s.items) {
		panic("Error")
	}
	s.items = append(s.items[:posicion+1], s.items[posicion:]...)
}

func (s *slicelist) Remove(posicion int) {
	if posicion < 0 || posicion > len(s.items) {
		panic("Error")
	}
	s.items = append(s.items[posicion:], s.items[:posicion+1]...)

}

func (s *slicelist) Get(posicion int) int {
	if posicion < 0 || posicion > len(s.items) {
		panic("Error")
	}
	return s.items[posicion]
}

func (s *slicelist) Len() int {
	return len(s.items)
}

func (s *slicelist) String() string {
	str := "["
	for i, val := range s.items {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprintf("%v", val)
	}
	str += "]"
	return str
}

func (s slicelist) Concat(s2 slicelist) slicelist {
	// Crear una nueva slicelist con la capacidad de contener los elementos de ambas slicelists
	result := slicelist{
		items: make([]int, len(s.items), len(s.items)+len(s2.items)),
	}

	// Copiar los elementos de la primera slicelist a la nueva slicelist
	copy(result.items, s.items)

	// Agregar los elementos de la segunda slicelist a la nueva slicelist
	result.items = append(result.items, s2.items...)

	return result
}

func main() {
	// Example usage
	list := slicelist{}
	list2 := slicelist{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list2.Append(0)
	list2.Append(6)
	list2.Append(7)
	fmt.Println(list)
	list.InsertAt(1, 5)
	fmt.Println(list.Get(2))
	list.Remove(0)
	fmt.Println(list.Len())
	fmt.Println(list)
	fmt.Println(list.Concat(list2))
}
