package sort

func MergeSort(v []int) {
	//função para dividir o vetor v recursivamente
	tam := len(v)
	v1 := make([]int, tam/2)
	v2 := make([]int, tam/2)

	if tam > 1 {
		for i := 0; i < tam/2; i++ {
			v1[i] = v[i]
		}
		for i := tam / 2; i < tam; i++ {
			v2[i-(tam/2)] = v[i]
		}

		MergeSort(v1)
		MergeSort(v2)
		Merge(v1, v2)
	}
}

func Merge(v1 []int, v2 []int) []int {
	n := len(v1)
	m := len(v2)
	new_v := make([]int, n+m)
	cont1 := 0
	cont2 := 0
	cont3 := 0

	for cont3 < len(new_v) {
		if v1[cont1] > v2[cont2] {
			new_v[cont1] = v2[cont2]
			cont3++
			cont2++
		} else {
			new_v[cont1] = v1[cont1]
			cont3++
			cont1++
		}
	}

	return new_v
}
