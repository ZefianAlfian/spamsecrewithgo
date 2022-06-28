package main

import (
	"fmt"
	"strconv"
	"bytes"
	// "io/ioutil"
	"time"
	"net/http"
	"encoding/json"
)

// type SecretoComment struct {
	// Id string `json:"id"`
	// MsgId string `json:"msgid"`
	// Comment string `json:"comment"`
// }

type SecretoMessage struct {
	Id string `json:"id"`
	Message string `json:"message"`
}

func input(text string) string {
    fmt.Print(text)

    // var then variable name then variable type
    var first string

    // Taking input from user
    fmt.Scanln(&first)

    return first
}


func startReq (url string, jsonData SecretoMessage) {
	body, _ := json.Marshal(jsonData)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")
	
	res, _ := client.Do(req)
	fmt.Println(res.StatusCode)
}

func main(){
	fmt.Printf("Hello, this tools spam secreto with thread\n\n")

	url := input("Url : ")
	id := input("ID : ")
	// msgId := input("Message ID : ")
	comment := input("Comment : ")
	thread := input("Thread : ")

	thrd, _ := strconv.Atoi(thread)

	// jsonData := SecretoComment{
		// Id: id,
		// MsgId: msgId,
		// Comment: comment,
	// }

	jsonData := SecretoMessage{
		Id: id,
		Message: comment,
	}

	for i := 0; i < thrd; i++ {
		go func() {
			for {
				startReq(url, jsonData)
				time.Sleep(30 * time.Second)
			}
		}()
	}
	for {
		startReq(url, jsonData)
		time.Sleep(30 * time.Second)
	}
	
}
