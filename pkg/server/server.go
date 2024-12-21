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

	inp := make(map[string]any)
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil{
		send500(w)
		return
	}

	expression, ok := inp["expression"].(string)
	if !ok{
		send500(w)
		return
	}

	res, err := calc.Calc(expression)
	if err != nil{
		rErr := ErrResult{err.Error()}
		sendMessage(w, 422, rErr)
		return
	}

	rRes := OkResult{res}
	sendMessage(w, 200, rRes)
}

func sendMessage(w http.ResponseWriter, status int, message any){
	w.WriteHeader(status)
	response, err := json.MarshalIndent(message, "", "   ")
	if err != nil{
		send500(w)
		return 
	}
	if _, err := w.Write(response); err != nil{
		send500(w)
		return
	}
}

func send500(w http.ResponseWriter){
	w.WriteHeader(500)
	r, _ := json.MarshalIndent(error500,"", "    ")
	w.Write(r)
}

