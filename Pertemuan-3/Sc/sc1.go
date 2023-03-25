package main

	import "sort"

func MaxPleasure(n int, k int, size []int, taste []int) int {
	if n == k {
		return -1
	} else {
		arrT := make([]int, n)
		
		for i := 0; i < n; i++ { 
			arrT[i] = size[i] * taste[i]
		}
		
		sort.Slice(arrT, func(i, j int) bool { return arrT[i] > arrT[j] })


		// mp := arrT[0]
		// for i := 1; i < len(arrT); i++ {

		// 	if arrT[i] > mp {
        //         mp = arrT[i]
        //     }
        // }
		return arrT[0]
		}
	}

func main() {
	n := 9
	k := 9

	sizes := []int{4, 2, 3, 20}
	taste := []int{4, 5, 6, 20}

	result := MaxPleasure(n, k, sizes, taste)

	println(result)
}
