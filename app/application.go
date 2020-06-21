package app

import (
	"github.com/maitungmn/bookstore_items-api/utils/env_utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maitungmn/bookstore_items-api/clients/elasticsearch"
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
