package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/maitungmn/bookstore_items-api/domain/queries"
	"strings"

	"github.com/maitungmn/bookstore_items-api/clients/elasticsearch"
	"github.com/maitungmn/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
	typeItem = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.ID
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %v", i.ID), errors.New("database error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	i.ID = itemId
	return nil
}

func (i *Item)Search(query queries.EsQuery)([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}
	fmt.Println(result)

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error wen trying to parse response", errors.New("database error"))
		}
		item.ID = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found matching given criteria")
	}

	return items, nil
}
