package models

type AuthenticationMatrix struct {
	BaseModelSoftDelete
	UserID       uint       `gorm:"not null;index:user_auth_matrix_idx" db:"user_id"`
	User         User       `gorm:"-"`
	AuthMethodID uint       `gorm:"not null" db:"auth_method_id"`
	AuthMethod   AuthMethod `gorm:"-"`
	Token        string     `gorm:"not null" db:"token"`
}

