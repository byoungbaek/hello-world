package main

import "fmt"
import "github.com/byoungbaek/my-error" // package path for package my_error

func main() {
	//var x float64
	//var eee error
	/*x, eee = */ calc(-20)
	v, err := calc(-50)
	if err != nil {
		fmt.Println("Err:", err)
	} else {
		fmt.Println("Hello Go", v)
	}

	go sum(0, 9)

	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}

func calc(a int) (float64, error) {
	if a < 0 {
		return 0, my_error.New(-10, "hahaha")
	}
	var a0 int8
	a0 = 10
	b := 2.4
	c := 0.0
	c = float64(int(a0)+a) + b
	return c, nil
}

func sum(start int, end int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
		fmt.Println("Hi:", i)
	}
	fmt.Println("Hi-sum:", sum)
}
