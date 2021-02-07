package models

type SearchLogModel struct {
	UUID      string `gorm:"column:uuid" json:"uuid"`
	SearchKey string `gorm:"column:search_key" json:"search_key"`
	Page      int    `gorm:"column:page" json:"page"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
}

func (SearchLogModel) TableName() string {
	return "search_log"
}
