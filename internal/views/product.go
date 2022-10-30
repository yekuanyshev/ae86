package views

type ProductSearchResult struct {
	Products []ProductShort `json:"products"`
}

type ProductShort struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Price int    `json:"price"`
}

type ProductFull struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Price       int          `json:"price"`
	Description string       `json:"description"`
	Image       string       `json:"image"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}
