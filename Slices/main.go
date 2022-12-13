package main

import (
	"fmt"
)

func main() {
	var a []string
	// Nối thêm các phần tử
	a = append(a, "Hoa", "Thang", "Loc")
	a = append(a, "Duy", "Huy")
	fmt.Println("All: ", a)
	// Xóa tên của Hòa
	remove := remove(a, "Hoa")
	fmt.Println("Sau khi xoa: ", remove)
}

// Xóa phần tử trong slices
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
