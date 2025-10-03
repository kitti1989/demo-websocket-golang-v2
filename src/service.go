package src

import (
	"fmt"
	"io"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pong")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello API")
}

func CallModel(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://model-api:8000/ping","application/json", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("error calling model: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// อ่าน response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error reading response: %v", err), http.StatusInternalServerError)
		return
	}

	// ส่ง response กลับไปให้ client
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}