package views

type Index struct {
	StoreTitle string     `json:"storeTitle"`
	Banners    []string   `json:"banners"`
	Categories []Category `json:"categories"`
}

type Category struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}
