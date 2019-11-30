package models

type AuthenticationMatrix struct {
	BaseModelSoftDelete
	UserID      uint   `gorm:"not null" db:"user_id"`
	User User `gorm:"foreignkey:id;association_foreignkey:user_id"`
	AuthMethodID uint    `gorm:"not null" db:"auth_method_id"`
	AuthMethod AuthMethod `gorm:"foreignkey:id;association_foreignkey:auth_method_id"`
	Token        string `db:"token"`
}

