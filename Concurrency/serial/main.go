package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)


func checkAndSave(url string, wg *sync.WaitGroup){
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN", url)
	}else{
		defer resp.Body.Close()

		fmt.Printf("%s status code => %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			file := strings.Split(url, "//")[1]
			file += ".txt"
			data,err := ioutil.ReadAll(resp.Body)

			if err!=nil{
				fmt.Println(err)
			}else{
				err := ioutil.WriteFile(file, data, 0664)
				if err != nil{
					fmt.Println(err)
				}
			}
		}else{
			log.Fatal(resp.StatusCode)
		}

	}
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	
	urls := []string{"https://www.goal.com", "https://golang.org"}
	wg.Add(len(urls))
	for _, url := range urls{
		go checkAndSave(url,&wg)
		println(strings.Repeat("#",20))
	}
	wg.Wait()
}