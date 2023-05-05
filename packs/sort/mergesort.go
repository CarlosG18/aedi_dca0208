package sort

func MergeSort(v []int, tam int) []int{
	if tam > 1 {
		meio := tam / 2
		tam_v1 := meio
		tam_v2 := tam - meio
		v1 := make([]int, tam_v1)
		v2 := make([]int, tam_v2)

		for i := 0; i < meio; i++ {
			v1[i] = v[i]
		}

		for i := meio; i < tam; i++ {
			v2[i-meio] = v[i]
		}

		v1 = MergeSort(v1, len(v1))
		v2 = MergeSort(v2, len(v2))
		return Merge(v1, v2)
	}
	return v
}

func Merge(v1 []int, v2 []int) []int{
	n := len(v1)
	m := len(v2)
	new_v := make([]int, n+m)
	cont1 := 0
	cont2 := 0
	cont3 := 0

	for cont1 < n && cont2 < m {
		if v1[cont1] > v2[cont2] {
			new_v[cont3] = v2[cont2]
			cont3++
			cont2++
		} else {
			new_v[cont3] = v1[cont1]
			cont3++
			cont1++
		}
	}
	for cont1 < n {
		new_v[cont3] = v1[cont1]
		cont1++
		cont3++
	}
	for cont2 < m {
		new_v[cont3] = v2[cont2]
		cont2++
		cont3++
	}
	return new_v
}
