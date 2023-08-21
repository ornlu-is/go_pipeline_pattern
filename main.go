package main

import "fmt"

func numberStream() <-chan float64 {
	ch := make(chan float64)
	numberStrings := []float64{1., 2., 3., 4., 5., 6., 7., 8., 9., 10.}

	go func() {
		for _, numberString := range numberStrings {
			ch <- numberString
		}

		close(ch)
		return
	}()

	return ch
}

func power(data <-chan float64) <-chan float64 {
	ch := make(chan float64)

	go func() {
		defer close(ch)
		for value := range data {
			ch <- value * value
		}
	}()

	return ch
}

func duplicate(data <-chan float64) <-chan float64 {
	ch := make(chan float64)

	go func() {
		defer close(ch)
		for value := range data {
			ch <- 2. * value
		}
	}()

	return ch
}

func main() {
	dataStream := numberStream()

	for value := range duplicate(power(dataStream)) {
		fmt.Println(value)
	}

	fmt.Println("bye")
}
