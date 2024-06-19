package models

type Patient struct {
    ID        uint   `gorm:"primaryKey"`
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Email     string `json:"email"`
}
