package domain

import "gorm.io/gorm"

type ProductMysql struct {
	gorm.Model
	Name  string
	Price int
}
