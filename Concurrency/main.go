package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func goroutineFunction(wg *sync.WaitGroup){
	fmt.Println("Begin execution of gorountineFunction")
	for i:=0; i<5; i++{
		fmt.Println("f1:", i)
		time.Sleep(time.Second * 2)
	}
	fmt.Println("Execution completed for gorountineFunction")
	wg.Done()
}

func normal(){
	fmt.Println("Begin execution of normalFunction")
	for i:=5; i<10; i++{
		fmt.Println("f2:", i)
	}
	fmt.Println("Execution completed normalFunction")
}

func main(){
	fmt.Println("No. of CPUs:", runtime.NumCPU())
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	fmt.Println("Os:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	var wg sync.WaitGroup
	wg.Add(1)
	
	go goroutineFunction(&wg)
	normal()
	wg.Wait()
	fmt.Println("Execution of main completed")
}