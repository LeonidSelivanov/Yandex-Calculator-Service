package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/LeonidSelivanov/Yandex-Calculator-Service/pkg/calculation"
)

type Config struct {
	Addres string
}

type Application struct {
	config *Config
}

type Request struct {
	Expession string `json:"expression"`
}

func New() *Application {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Application{
		config: &Config{
			Addres: ":" + port,
		},
	}
}

func (app *Application) CalcHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	result, err := calculation.Calc(req.Expession)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "result: %f", result)
	w.WriteHeader(http.StatusOK)
}

func (app *Application) RunServer() {
	http.HandleFunc("/api/v1/calculate", app.CalcHandler)
	http.ListenAndServe(":8080", nil)
}
