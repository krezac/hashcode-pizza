package main

func (s *slice) isValid(pizza []string, minItems, maxSize int) bool {
	if (s.r1-s.r0+1)*(s.c1-s.c0+1) > maxSize {
		return false
	}
	mushrooms, tomatoes := 0, 0
	for i := s.r0; i <= s.r1; i++ {
		for j := s.c0; j <= s.c1; j++ {
			if pizza[i][j] == Tomato {
				tomatoes++
			} else if pizza[i][j] == Mushroom {
				mushrooms++
			} else {
				panic("Invalid char in pizza")
			}
		}
		if tomatoes < minItems || mushrooms < minItems {
			return false
		}
	}
	return true
}

func process(pizza []string, minItems, maxSize int) []slice {
	slices := []slice{}

	for r, row := range pizza {
		start := 0
		for start < len(row) {
			found := false
			for i := 0; i < maxSize && start+i < len(row); i++ {
				//fmt.Printf("row %d, start: %d, len: %d\n", r, start, i)
				slice := slice{
					r0: r,
					r1: r,
					c0: start,
					c1: start + i,
				}
				if slice.isValid(pizza, minItems, maxSize) {
					slices = append(slices, slice) // test line
					start = slice.c1 + 1
					//fmt.Printf("found, new start: %d\n", start)
					found = true
					break
				}
			}
			if !found {
				start++
			}
		}
	}
	return slices
}
