package main

import "fmt"

func passValueToChannel(num int, c chan int){
	c <- num
}

func factorial(n int, c chan int ){
	f := 1
	for i:=2; i<=n; i++{
		f*=i
	}
	c <- f
}

func main() {
	// Declaration of Bidirectional channel
	var c chan int

	// Short declaration and assignment of bidirectional channel
	c1 := make(chan int)
	defer close(c1)

	// Only receiving channel
	c2 := make(<-chan string)

	// Only sending channel
	c3 := make(chan<- string)
	defer close(c3)
	fmt.Printf("%T, %T, %T, %T\n", c, c1, c2, c3)

	go passValueToChannel(20, c1)
	received := <- c1

	fmt.Println("Data received from channel is:", received)
	go factorial(5, c1)
	f := <- c1
	fmt.Printf("factorial of %d is %d\n", 5, f)

	// Goroutines, channels and anonymous functions
	var result int

	for i:=1; i<=20; i++{
		go func (){
			f:=1
			for j:=2; j<=i; j++{
				f*=j
			}
			c1 <- f
		}()
		result = <- c1
		fmt.Printf("Anonymous func factorial of %d is: %d\n", i, result)
	}
}
