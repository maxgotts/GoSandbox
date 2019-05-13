package main

import(
	"fmt"
	"sort"
)

func main() {
	arrays := [][2]int{{0,1},{4,3},{3,2}}
	sort.SliceStable(arrays, func (i, j int) bool {
		fmt.Println(i,j)
		return arrays[i][0] < arrays[j][0]
	})
	fmt.Println(arrays)
}

