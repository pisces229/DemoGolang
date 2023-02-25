package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func demoHttp() {
	root, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	run := func(ctx context.Context) {
		//resource := "https://www.google.com/"
		resource := "https://localhost:9110/api/default/run"
		//resp, err := http.Get(resource)
		client := &http.Client{
			//Timeout: 1 * time.Second,
		}
		req, err := http.NewRequest("GET", resource, nil)
		req = req.WithContext(ctx)
		resp, err := client.Do(req)
		fmt.Println("done")
		fmt.Println(err)
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(err)
		fmt.Println(string(body))
	}
	go run(root)
	time.Sleep(1 * time.Second)
	cancel()
	fmt.Println("cancel")
}
