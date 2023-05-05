package sort

func SelectionSort(v [7]int) [7]int {
	tam := len(v)

	for j := 0; j < tam-1; j++ {
		index_menor := j
		for i := 1 + j; i < tam; i++ {
			if v[index_menor] > v[i] {
				index_menor = i
			}
		}
		v[j], v[index_menor] = v[index_menor], v[j]
	}

	return v
}
