package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ForwardRequest(w http.ResponseWriter, r *http.Request) {
	host := "http://0.0.0.0:9001"
	reqUrl := fmt.Sprintf("%s%s", host, r.URL)
	req, err := http.NewRequest(r.Method, reqUrl, r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error crafting the request: %v", err), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending request to API: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error making the request: %v", err)
	}

	w.Write(body)
}
