package main
import "fmt"
func PairSum(arr []int, target int) []int {
    kiri, kanan := 0, len(arr)-1
    for kiri < kanan {
        jumlah := arr[kiri] + arr[kanan]
        if jumlah == target {
            return []int{kiri, kanan}
        } else if jumlah < target {
            kiri++
        } else {
            kanan--
        }
    }
    return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1,2,3,4,6},6))
	fmt.Println(PairSum([]int{2,5,9,11},11))
	fmt.Println(PairSum([]int{1,3,5,7},12))
	fmt.Println(PairSum([]int{1,4,6,8},10))
	fmt.Println(PairSum([]int{1,5,6,7},6))
}