package application

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
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
	Expression string `json:"expression"`
}
type Result struct{
	ResultString string `json:"result"`
}

func CalculateHandle(w http.ResponseWriter, r *http.Request){
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err1 := calc.Calc(request.Expression)
	resCalcStr := strconv.FormatFloat(result, 'g', -1, 64)
	resultJson := Result{ResultString: resCalcStr}
	bytesJson, err2 := json.Marshal(resultJson)
	if err2 != nil{

	}
	if err1 != nil{
		w.WriteHeader(422)
		w.Write([]byte(`{"error": "Expression is not valid"}`))
	} else if err2 != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"error": "Internal server error"}`))
	}else {
		w.WriteHeader(200)
		w.Write(bytesJson)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalculateHandle)
	return http.ListenAndServe(":" + a.config.Addr, nil)
}