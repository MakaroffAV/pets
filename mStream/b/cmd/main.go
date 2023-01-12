package main

import (
	"b/pkg/tnreq"
	"net/http"
	"os"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Program start point
	*/

	r, _ := tnreq.Get("https://www.youtube.com/results?search_query=g+eazy")

	f, err := os.Create("sample.file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(string(r))
	if err != nil {
		panic(err)
	}

	//s := &server{}
	//http.Handle("/", s)
	//log.Fatal(http.ListenAndServe(":8080", nil))

}
