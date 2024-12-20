package server

import (
	"encoding/json"
	"net/http"
	"example.com/m/pkg/calc"
)

type Server struct{}

func NewServer() *Server{
	return &Server{}
}

func(c *Server) Run(){
	http.HandleFunc("/api/v1/calculate", calcHandler)
	http.ListenAndServe("", nil)
}

func calcHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var inp Input
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil{
		send500(w)
		return
	}

	expression := inp.Expression

	res, err := calc.Calc(expression)
	if err != nil{
		rErr := ErrResult{err.Error()}
		sendMessage(w, 402, rErr)
		return
	}

	rRes := OkResult{res}
	sendMessage(w, 200, rRes)
}

func sendMessage(w http.ResponseWriter, status int, message any){
	w.WriteHeader(status)
    if encodeErr := json.NewEncoder(w).Encode(message); encodeErr != nil {
        send500(w)
    }
}

func send500(w http.ResponseWriter){
	w.WriteHeader(500)
	r, _ := json.Marshal(error500)
	w.Write(r)
}

