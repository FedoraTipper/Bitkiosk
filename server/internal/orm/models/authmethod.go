package models

type AuthMethod struct {
	BaseModelSoftDelete
	MethodId uint `db:"method_id";AUTO_INCREMENT`
	Name string `db:"name" gorm:"not null"`
}