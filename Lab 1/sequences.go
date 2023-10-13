package main

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for index, elem := range slice {
		slice[index] = f(elem)
	}
}

func mapArray(f func(a int) int, array [5]int) [5]int {
	for index, elem := range array {
		array[index] = f(elem)
	}
	// have to return because arrays are values, and slices are references to values
	return array
}

/*func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	intsArray := [5]int{1, 2, 3, 4, 5}
	intsArray = mapArray(addOne, intsArray)
	fmt.Println(intsArray)

	newSlice := intsSlice[1:3]
	newSlice = mapSlice(square, newSlice)
	fmt.Println(newSlice)

	newSlice := double(intsSlice)
	fmt.Println(newSlice)
}*/
