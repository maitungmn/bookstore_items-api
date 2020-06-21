package services

import (
	"github.com/maitungmn/bookstore_items-api/domain/items"
	"github.com/maitungmn/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(i items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, nil
}

func (s *itemsService) Get(i string) (*items.Item, rest_errors.RestErr) {
	return nil, nil
}
