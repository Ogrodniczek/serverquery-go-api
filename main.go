package main
import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

    r.Handle("/status", StatusHandler).Methods("GET")

    http.ListenAndServe(":8009", handlers.LoggingHandler(os.Stdout, r))

}
var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello API"))
})
