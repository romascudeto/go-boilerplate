package models

type Employee struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func (b *Employee) TableName() string {
	return "employee"
}
