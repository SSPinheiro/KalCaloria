package schemas

import (
	"gorm.io/gorm"
)

type Comida struct {
	gorm.Model
	Nome    string
	Tipo    string
	Caloria int64
}
