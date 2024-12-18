package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/vizurth/calc_go/pkg/calc"
)

type Config struct{
	Addr string
}

func ConfigFromEnv() *Config{
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == ""{
		config.Addr = "8080"
	}
	return config
}

type Application struct{
	config *Config
}

func New() *Application{
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct{
	Expression string
}


func CalculateHandle(w http.ResponseWriter, r *http.Request){
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calc.Calc(request.Expression)
	if err != nil{
		fmt.Fprintf(w, "err: %s", err.Error())
	} else {
		fmt.Fprintf(w, "result: %f", result)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculation", CalculateHandle)
	return http.ListenAndServe(":" + a.config.Addr, nil)
}