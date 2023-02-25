package app

import (
	"fmt"
	"net/http"
)

func demoHttp() {
	resp, err := http.Get("https://www.google.com/")
	//client := &http.Client{
	//	Timeout: 1 * time.Second,
	//}
	//req, err := http.NewRequest("GET", "https://www.google.com/", nil)
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	//req = req.WithContext(ctx)
	//resp, err := client.Do(req)
	fmt.Println(err)
	fmt.Println(resp)
	//cancel()
}
