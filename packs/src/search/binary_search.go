package search

// forma iterativa
func Binsearch_ite(values []int, value int) int {
	start := 0
	end := len(values) - 1

	for start <= end {
		mid := (int) (start + end) / 2
		if value == values[mid] {
			return mid
		} else if value < values[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
  return -1
}

//forma recursiva

func Binsearch_rec(values []int, start int, end int, value int) int {
	if start <= end {
		mid := (int)(start+end) / 2
		if value == values[mid] {
			return mid
		} else if value < values[mid] {
			Binsearch_rec(values, start, mid-1, value)
		} else {
			Binsearch_rec(values, mid+1, end, value)
		}
	}else{
    return -1
  }
}
