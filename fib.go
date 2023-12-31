package gopark

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	cursors := make([]int, 2)
	return func() int {
		if cursors[0] == 0 {
			cursors[0] = 1
			cursors[1] = 1
			return 0
		}
		retVal := cursors[0]
		temp := cursors[1]
		cursors[1] = cursors[0] + cursors[1]
		cursors[0] = temp
		return retVal
	}
}
