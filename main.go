package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"time"
	"sync"
)

type Todos struct {
	userID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	CompleteState bool `json:"completed"`
}

func main() {
	// callApi()
	// WaitGroupSync()
	ChannelSync()
}

func callApi() {
	response ,err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Print("Error resquest")
	} else {
		var TodosList []Todos
		body , errBody := ioutil.ReadAll(response.Body)
		if errBody != nil {
			fmt.Print("Error")
		} else {

		}
		defer response.Body.Close()

		json.Unmarshal(body , &TodosList)
		for _,v := range TodosList {
			fmt.Println(v)
		}
	}
}

func Todo1(arrNum []int , waitGroup *sync.WaitGroup) int {
	count := 0
	for _, v := range arrNum {
		count += v
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
	waitGroup.Done()
	return count
}

func ChannelSync() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func () {
		time.Sleep(3 * time.Second)
		c1 <- 5
	}()

	go func () {
		time.Sleep(4 * time.Second)
		c2 <- 4
	}()

	msg :=  <- c1
	msg2 := <- c2
	total := (msg + msg2)
	fmt.Println(total)
}

func WaitGroupSync() {
	var waitgroup sync.WaitGroup
	arrNum1 := []int{1,3,5,7,9}
	arrNum2 := []int{2,4,6,8,10}
	waitgroup.Add(1)
	go Todo1(arrNum1 , &waitgroup)
	waitgroup.Add(1)
	go Todo1(arrNum2 , &waitgroup)

	waitgroup.Wait()

	fmt.Println("End sync")
}