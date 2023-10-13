package main

func calculateNextState(p golParams, worlds [][]byte) [][]byte {
	world := make([][]byte, len(worlds))
	for i := range world {
		world[i] = make([]byte, len(worlds[i]))
	}
	//copy(world, worlds)
	for index, element := range worlds {
		var prevElement, nextElement []byte

		prevElement = worlds[((index + len(world) - 1) % len(world))]
		nextElement = worlds[(index+len(world)+1)%len(world)]

		for index2, element2 := range element {
			// makes a list of the neighbours of the element2
			values := []byte{prevElement[index2], nextElement[index2]}
			if index2 > 0 {
				values = append(values, element[(index2-1)%len(worlds)], prevElement[(index2-1)%len(element)], nextElement[(index2-1)%len(element)])
			} else {
				values = append(values, element[15], prevElement[15], nextElement[15])
			}
			if index2 < 15 {
				values = append(values, element[(index2+1)%len(element)], prevElement[(index2+1)%len(element)], nextElement[(index2+1)%len(element)])
			} else {
				values = append(values, element[0], prevElement[0], nextElement[0])
			}
			num := 0
			for _, value := range values {
				if value > 0 {
					num += 1
				}
			}
			// if the element is dead, then run through those checks
			if element2 == 0 {
				if num > 0 {
					//fmt.Println(num, ", world", index2, index)
				}
				if num == 3 {
					world[index][index2] = 255
				} else {
					world[index][index2] = 0
				}

			} else {
				if num < 2 || num > 3 {
					world[index][index2] = 0
					//fmt.Println(world[index])
					//fmt.Println(worlds[index])
				} else {
					world[index][index2] = 255
				}

			}
		}
	}
	return world

}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	coordinates := []cell{}
	for index, row := range world {
		for index2, _ := range row {
			if world[index][index2] > 0 {
				coordinates = append(coordinates, cell{index2, index})
			}
		}
	}
	return coordinates
}
