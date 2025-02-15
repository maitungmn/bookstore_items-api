package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maitungmn/bookstore_items-api/src/clients/elasticsearch"
	"github.com/maitungmn/bookstore_utils-go/env_utils"
)

const (
	LocalUrl = "LOCAL_URL"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {

	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    env_utils.GetEnvVariable(LocalUrl),
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
