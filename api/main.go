package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func requestBodyHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディを読み取る
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// リクエストボディを出力
	fmt.Printf("Received request body: %s\n", body)

	// レスポンスを返す
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func main() {
	http.HandleFunc("/echo", requestBodyHandler)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
