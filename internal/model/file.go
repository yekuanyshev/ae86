package model

type File struct {
	ID          int64  `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	ContentType string `gorm:"column:content_type"`
	Size        int64  `gorm:"column:size"`
	Content     []byte `gorm:"column:content"`
}

func (File) TableName() string {
	return "file"
}
