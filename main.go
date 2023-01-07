package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://sukebei.nyaa.si/?f=0&c=0_0&q=&s=leechers&o=desc")
		if err != nil {
			fmt.Printf("Error getting sukebei %s", err.Error())
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading sukebei %s", err.Error())
		}
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	})
	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
