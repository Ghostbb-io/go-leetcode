package _1377

import "sort"

type soldiers struct {
	total int
	row   int
}

func kWeakestRows(mat [][]int, k int) []int {
	eachRow := make([]soldiers, len(mat))
	for i, m := range mat {
		temps := soldiers{total: 0, row: i}
		for _, s := range m {
			if s == 0 {
				break
			}
			temps.total++
		}
		eachRow[i] = temps
	}
	sort.SliceStable(eachRow, func(i, j int) bool {
		if eachRow[i].total == eachRow[j].total {
			return eachRow[i].row < eachRow[j].row
		}
		return eachRow[i].total < eachRow[j].total
	})

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, eachRow[i].row)
	}
	return result
}
