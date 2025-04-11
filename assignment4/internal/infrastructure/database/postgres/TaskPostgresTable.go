package postgres

import "time"

type Task struct {
	ID uint `gorm:"column:id;unique;primaryKey;autoIncrement" json:"id"`
	Title string `gorm:"column:title;" json:"title"`
	Description string `gorm:"column:description;" json:"description"`
	Completed bool `gorm:"column:completed;" json:"completed"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;autoCreate" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdate" json:"updated_at"`
}

