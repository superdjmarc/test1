package main
import (
	"github.com/gorilla/mux"
	"github.com/superdjmarc/test1.git/package1"
	"log"
	"net/http"
	"time"
)

//https://blog.jetbrains.com/go/2020/02/26/working-with-go-modules-getting-started/

func main() {

	package1.PrintHello()

	mtx := mux.NewRouter()
	mtx.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World!"))
	})
	srv := &http.Server{
		Handler:      mtx,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
