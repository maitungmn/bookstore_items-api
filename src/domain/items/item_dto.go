package items

type Item struct {
	ID                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	Plaintext string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}
