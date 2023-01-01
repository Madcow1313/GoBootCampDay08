package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	bound := len(arr)
	if bound == 0 {
		return 0, errors.New("error! empty array")
	}
	if idx > bound-1 {
		return 0, errors.New("error! index out of bounds")
	}
	if idx < 0 {
		return 0, errors.New("error! index can only be positive integer")
	}
	pointer := unsafe.Pointer(&arr[0])
	size := unsafe.Sizeof(&arr[0])
	return *(*int)(unsafe.Pointer((uintptr)(pointer) + size*uintptr(idx))), nil
}

func catchOrPrint(res int, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	res, err := getElement(arr, 1)
	catchOrPrint(res, err)

	arr = []int{1, 2, 3, 4, 5}
	res, err = getElement(arr, -1)
	catchOrPrint(res, err)

	arr = []int{1, 2, 3, 4, 5}
	res, err = getElement(arr, 10)
	catchOrPrint(res, err)

	arr = []int{}
	res, err = getElement(arr, 0)
	catchOrPrint(res, err)
}
