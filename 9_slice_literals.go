package main

import "fmt"

func main() {
	// В Go slice literal — это удобный способ сразу создать и инициализировать срез (slice).
	// По сути, это синтаксис для объявления среза с начальными значениями,
	// похожий на массивный литерал, только без указания длины.
	d := []int{0, 1, 5, 7, 11, 13}
	fmt.Println(d)

	b := []bool{true, false, true, true, false, true}
	fmt.Println(b)

	result := []struct {
		i int
		b bool
	}{
		{d[0], b[0]},
		{d[1], b[1]},
		{d[2], b[2]},
		{d[3], b[3]},
		{d[4], b[4]},
		{d[5], b[5]},
	}

	fmt.Println(result)

}
