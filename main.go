package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
	"net/http"
	"no_vcs/me/dynamo-db-example/internal/store"
	"os"
)

func main() {
	fmt.Println("starting...")
	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("DYNAMO_REGION")),
		Endpoint: aws.String(os.Getenv("DYNAMO_ENDPOINT")),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)
	s := store.NewLeagueInfoGetter(os.Getenv("DYNAMO_TABLE"), db)

	r := mux.NewRouter()
	r.HandleFunc("/league/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)

		league, err := s.GetLeagueInfo(params["id"])

		if err != nil {
			fmt.Printf("ERR: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if league == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		b, err := json.Marshal(league)

		if err != nil {
			panic(err)
		}

		_, err = w.Write(b)

		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("running server on port :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
		fmt.Println(err)
	}
}