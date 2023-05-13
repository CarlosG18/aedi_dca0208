package sort

func BubbleSort(v []int) []int {
	tam := len(v)
	mudou := false

	for j := 0; j < tam-1; j++ {
		for i := 0; i < tam-j-1; i++ {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				mudou = true
			}
		}
		if mudou != true {
			break
		}
	}
	return v
}