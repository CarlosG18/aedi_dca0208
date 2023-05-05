package sort

func InsertionSort(v [5]int) [5]int {
	tam := len(v)

	for i := 1; i < tam; i++ {
		valor := v[i]
		local := i

		for k := i; k > 0; k-- {
			if valor < v[k-1] {
				v[k] = v[k-1]
				local = k - 1
			}
		}
		v[local] = valor
	}

	return v
}
