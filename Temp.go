package main

import (
	"fmt"
	"time"
)

type MyCar struct {
	innerFloat  float64
	innerString string
}

func (myCar MyCar) Horn() {
	fmt.Println("Horn")
}

func (myCar MyCar) Brake() error {
	fmt.Println("Brake")
	return CustomError{errorTime: time.Now(), errorVal: "Our cool new Error message"}
}

func (myCar MyCar) LooksGood() {
	fmt.Println("Brake")
}

type ABC struct {
}

func (abc *ABC) Horn() {
	fmt.Print("ABC Horn")
}

type CustomError struct {
	errorTime time.Time
	errorVal  string
}

func (customError CustomError) Error() string {
	return fmt.Sprintf("%v", customError.errorTime) + ", " + customError.errorVal
}

func (abc *ABC) Brake() error {
	fmt.Print("ABC Horn")
	return nil
}

type Car interface {
	Horn()
	Brake() error
}

func print(intVal Car) {
	fmt.Printf("%+v\n", intVal)
}

func main() {
	var car Car = MyCar{}
	//describe(car)
	err := car.Brake()
	fmt.Println(err)

	//i := 100
	//var i interface{}
	//i = "ABC"
	//j, ok := i.(int)
	//fmt.Printf("Type: %T, j: %v, j AssertedOrNot %v", j, j, ok)
}

func describe(i Car) {
	switch v := i.(type) {
	case MyCar:
		fmt.Println("MyCar", v)
	case *ABC:
		fmt.Println("*ABC")
	}
}
