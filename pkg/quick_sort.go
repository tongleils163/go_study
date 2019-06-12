package pkg

import "fmt"

func QuickSort(data []int, left, right int) []int {
	if left >= right {
		return data
	}
	tmp := data[left]
	fmt.Println(data)
	i := left
	j := right
	for i != j {
		for data[j] >= tmp && i < j {
			j--
		}
		for data[i] <= tmp && i < j {
			i++
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[left] = data[i]
	data[j] = tmp
	QuickSort(data, left, i-1)
	QuickSort(data, i+1, right)
	return data
}

