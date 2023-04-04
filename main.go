package main

import (
  "fmt"
	"main/packs/src/search"
)

func main() {
	values := make([]int, 20)

	for i := 0; i < 20; i++ {
		values[i] = i
	}

  for i := 0; i < 20; i++ {
    fmt.Println(search.Binsearch_rec(values, 0, len(values)-1, i))
  }
}


