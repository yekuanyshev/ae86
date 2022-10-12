package model

type Banner struct {
	ID      int64  `gorm:"column:id"`
	Image   string `gorm:"column:image"`
	StoreID int64  `gorm:"column:store_id"`
}

func (Banner) TableName() string {
	return "banner"
}
