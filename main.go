package main

import (
	"net/http"
	"os"

	controllers "github.com/danielsilveiradevbr/conciliacaoXLS/src/controllers/conciliacao"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	r.Post("/conciliacao", controllers.ConciliaXLXs)

	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "3000"
	}
	porta = ":" + porta
	println(porta)

	http.ListenAndServe(porta, r)
}
