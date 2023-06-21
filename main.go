package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	// fmt.Println("hello peeps")
	name := "Alhaji"
	stat := fmt.Sprintf("%v's age is 20", name)
	fmt.Println(stat)

	//Array
	countries := [5]string{"Ghana", "Nigeria", "Niger", "Mali", "South Africa"}

	//Slice
	companies := [] string{"4Sure", "Turntabl", "Amalitech"}

	// Map
	currencies := map[string]float64{
		"USD" : 45.9,
		"AUD" : 78.8,
	}

	// Struct which is similar to class in OOP
	type Student struct {
		id int
		name string
		course string
	}

	student := Student{1, "Jon", "Math"}

	// Pointer stores the memory address location of another variable
	age := 10
	ptr := &age

	// function type
	method := f

	fmt.Printf("countries is of type: %T\n", countries)
	fmt.Printf("companies is of type: %T\n", companies)
	fmt.Printf("currencies is of type: %T\n", currencies)
	fmt.Printf("student is of type: %T\n", student)
	fmt.Printf("ptr is of type: %T and value:%v\n", ptr, ptr)
	fmt.Printf("method is of type: %T\n", method)

	// If condition
	if age >100 {
		fmt.Println("You too old")
	}else{
		fmt.Println("you are young")
	}

	fmt.Println("Command Line args:", os.Args)

	if ans, err := strconv.ParseInt("23", 8, 32); err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%T", ans)
	}

	// For Loops 
	// for i := 0; i<10; i++{
	// 	fmt.Println(i)
	// }

	for i, a := range countries{
		fmt.Printf("%v => %v\t", i, a)
	}
	fmt.Println()

	// Files Handling 
	// Create a File
	file, err := os.Create("log.txt")
	if err !=nil {
		log.Fatal(err)
	}
	file.Close()
	
	// Open a File
	newFile, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	nf, err := os.OpenFile("legit.txt", os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	nf.Close()

	// Get File Information
	fileInfo, err := os.Stat("log.txt")

	fmt.Println("Filename: ", fileInfo.Name())
	fmt.Println("Size: ", fileInfo.Size())
	fmt.Println("Last Modified: ", fileInfo.ModTime())
	fmt.Println("Permissions: ", fileInfo.Mode())

	// Write to File using os package
	fileWrite, err := os.OpenFile("demo.txt", os.O_CREATE | os.O_APPEND | os.O_TRUNC | os.O_WRONLY, 0644)
	
	if err != nil {
		log.Fatal(err)
	}
	defer fileWrite.Close()

	data := [] byte ("Hello Go programming language")

	fileWritten, err := fileWrite.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Size of file written in byte is:", fileWritten)


	// Write to File using ioutil package
	ioData := [] byte ("Message written using ioutil package")
	ioFileError := ioutil.WriteFile("io.txt", ioData, 0644)

	if ioFileError != nil {
		log.Fatal(err)
	}
	// Write to File using bufio package
	buFile, err := os.OpenFile("bufio.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}	
	defer buFile.Close()

	bufWriter := bufio.NewWriter(buFile)

	firstData := "Welcome to Chelsea fc Manuel Ugarte"

	bufWriter.WriteString(firstData)

	fmt.Println("Available buffer size:", bufWriter.Available())

	secondData := "\nWe hope you enjoy the stay"

	bufWriter.Write([]byte(secondData))

	fmt.Println("Size of buffered data: ", bufWriter.Buffered())
	bufWriter.Flush()

	// Read a file using ioutil package
	records, err := ioutil.ReadFile("RestApi/api.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read:%s", records)
	fmt.Printf("Size of data:%d\n", len(records))

	// Read data from stdIn 
	// scanner := bufio.NewScanner(os.Stdin)

	// fmt.Println("\nEnter data here:")

	// success := scanner.Scan()

	// if success {
	// 	fmt.Println(scanner.Text())
	// }

	// Variadic functions
	varidic(34,5,67)

	// Anonymous functions
	func(msg string){
		fmt.Println(msg)
	}("Here's an example of an anonymous function")

	// Function that returns a function(anonymous)
	result := increment(10)
	fmt.Printf("type:%T\n", result)
	result()
	fmt.Println(result())
	
	// Receiver function demo
	car := car{name: "Audi", price: 45000.}
	newCar := car.changeCar("Honda", 5600.)
	(&car).change("Benz", 78909.)
	fmt.Printf("first car:%v\n", car)
	fmt.Printf("second car:%v\n", newCar)
	fmt.Printf("third car: %v\n", car)

	// Inteface demo
	dog := dog{name: "Lucky", age: 5}
	cat := cat{name: "Dodo", breed: "Asian"}
	print(dog)
	print(cat)

	// Embedded interface demo
	box := cube{edge: 5, color: "red"}
	fmt.Println(measure(box))

	// Empty interface can accept every type as a value
	var empty interface{}
	empty = "foobar"
	fmt.Printf("Empty is:%v\n", empty)
	empty = []int{1, 2, 3, 4}
	fmt.Printf("Empty is:%v\n", empty)
	fmt.Printf("The length of empty is:%v\n", len(empty.([]int))) // Type asserted empty interface to int slice

}


func f(){}

func varidic (nums ...int){
	fmt.Println(nums)
}

func increment(num int) func() int {
	return func() int {
		num++
		return num
	}
}

// Receiver functions or methods

type car struct {
	name string
	price float64
}

func (c car) changeCar(name string, price float64) car {
	c.name = name
	c.price = price
	return c
}

func (c *car) change(name string, price float64){
	c.name = name
	c.price = price
}

// interface
type animal interface{
	move() string
	eat() string
}

type dog struct{
	name string
	age int
}

type cat struct{
	name string
	breed string
}

func (d dog) move() string{
	return fmt.Sprintf("Dog by name %s is moving", d.name)
}

func (d dog) eat() string{
	return fmt.Sprintf("Dog %s is eating its meal", d.name)
}

func (c cat) move() string{
	return fmt.Sprintf("%s cat breed is moving and hissing", c.breed)
}

func (c cat) eat() string{
	return fmt.Sprintf("%s cat breed with name %s is eating big meal", c.breed, c.name)
}

func print(a animal){
	fmt.Printf("Animal type is: %#v\n", a)
	fmt.Println(a.move())
	fmt.Println(a.eat())
}

// Embedded interface

type shape interface{
	area() float64
}

type object interface{
	volume() float64
}

type geometry interface{
	shape
	object
	getColor() string
}

type cube struct{
	edge float64
	color string
}

func (c cube) area() float64{
	return 2 * c.edge
}

func (c cube) volume() float64{
	return 6 * c.edge
}

func (c cube) getColor() string{
	return c.color
}

func measure(g geometry) (float64, float64) {
	return g.volume(), g.area()
}