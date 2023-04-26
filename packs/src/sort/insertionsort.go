package sort

func InsertionSort(v [5]int) [5]int {
	tam := len(v)

	for i := 1; i < tam; i++ {
		j := i
		local := i
		for k := i; k > 0; k-- {
			if v[j] < v[k-1] {
				local = k - 1
				//colocar uma variavel temporaria para guardar o valor
			}
		}
		v[local] = v[j]
	}

	return v
}
