package main 

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func main() {
	invs := 0
	text, err := ioutil.ReadFile("IntegerArray.txt")
	if err != nil{
		panic(err)
	}
	string_array := strings.Split(string(text), "\r\n")
	numbers_array := []int{}
	string_array = string_array[:len(string_array)-1]
	for _, i := range string_array {
		num, err := strconv.Atoi(string(i))
		if err != nil{
			panic(err)
		}
		numbers_array = append(numbers_array, num)
	}
	invs, _ = MergeSort(numbers_array)
	fmt.Println("Inersions %q", invs)
}

func MergeSort(slice []int) (int, []int) {
	invs, temp, left_invs, right_invs := 0, 0, 0, 0
	res_merge, left_slice, right_slice := []int{}, []int{}, []int{}

	if len(slice) < 2 {
		return invs, slice
	}
	mid := (len(slice)) / 2
	left_invs, left_slice = MergeSort(slice[:mid])
	invs = left_invs
	right_invs, right_slice = MergeSort(slice[mid:])
	invs += right_invs
	res_merge, temp = Merge(left_slice, right_slice, mid)
	invs += temp

	return invs, res_merge
}

func Merge(left, right []int, mid int) ([]int, int) {

	size, i, j, invs := len(left)+len(right), 0, 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
			invs = invs + mid - i
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
			invs = invs + mid - i
		}
	}
	return slice, invs
}
