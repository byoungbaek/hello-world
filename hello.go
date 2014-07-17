package main

import "fmt"
//import "errors"

// My error handling -------------------------------------------------------

type myError struct {
	code int32
	s string
}

func (e *myError) Error() string {
	return fmt.Sprintf("[%d] %s", e.code, e.s)
}

func myNewError(code int32, text string) error {
    return &myError{code, text}
}

//--------------------------------------------------------------------------------

func m(a int) (float64, error) {
	if a < 0 {
		return 0, myNewError(-10, "hahaha")
	}
	var a0 int8
	a0 = 10
	b := 2.4
	c := 0.0
	c = float64(int(a0) + a)+b
	return c, nil
}

func f() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("Hi:", i)
	}
	fmt.Println("Hi-sum:", sum)
}

func main() {
	//var x float64
	//var eee error
	/*x, eee = */ m(-20)
	v, err := m(-50)
	if err != nil {
		fmt.Println("Err:", err)
	} else {
		fmt.Println("Hello Go", v)
	}
	go f()
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}
