package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambdaurl"
	"github.com/bmoffatt/must.go"
	"github.com/bmoffatt/smithy-lambda-url-example/model"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	http.HandleFunc("/wave", func(w http.ResponseWriter, r *http.Request) {
		spew.Print(must.Be(lambdaurl.RequestFromContext(r.Context())))
		var input model.WaveInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		spew.Print(input)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Text string
			Time time.Time
		}{
			Text: "hello " + *input.Name,
			Time: time.Now(),
		})
	})
	lambdaurl.Start(http.DefaultServeMux)
}
