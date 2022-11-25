package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Printf("Hãy nhập tên của bạn: ")
	fmt.Scanln(&name)
	fmt.Printf("Hãy nhập số tuổi :")
	fmt.Scanln(&age)
	if age > 100 {
		fmt.Printf("WTF BRO...")
	} else if age < 0 {
		fmt.Printf("LOLLLL KIDS")
	} else {
		fmt.Println("Xin Chào :", name, "\nTuổi: ", age, "\nChúc bạn 1 ngày mới vui vẻ")
	}
}
