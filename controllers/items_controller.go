package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/maitungmn/bookstore_items-api/domain/items"
	"github.com/maitungmn/bookstore_items-api/services"
	"github.com/maitungmn/bookstore_items-api/utils/http_utils"
	"github.com/maitungmn/bookstore_oauth-go/oauth"
	"github.com/maitungmn/bookstore_utils-go/rest_errors"
)

var (
	ItemController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// http_utils.RespondError(w, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = oauth.GetClientId(r)

	// item := items.Item{
	// 	Seller: oauth.GetCallerId(r),
	// }

	result, err := services.ItemsService.Create(itemRequest)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
