package main

import {
	"fmt"
}

func main() {

	/*
		slice := []int{1, 2, 3}
		for i := 0; i < len(slice); i++ {
			println(slice[i])
		}
	*/

	/*
		slice1 := []int{10, 20, 30}
		for i, v := range slice1 {
			println(i, v)
		}
	*/

	/*
		wellPorts := map[string]int{
			"http":  90,
			"https": 900}
		for k, v := range wellPorts {
			println(k, v)
		}
	*/

	wellPorts := map[string]int{
		"http":  90,
		"https": 900}
	for _, v := range wellPorts {
		println(v)
	}
}
