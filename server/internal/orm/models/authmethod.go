package models

type AuthMethod struct {
	ID         int
	TTL        int // minutes to refresh
	RefreshTTL int
}

var DefaultAuth = &AuthMethod{
	ID:  1,
	TTL: 6000,
	RefreshTTL: 3600, // 60 hours
}

func GetAuthMethod(ID int) *AuthMethod {
	switch ID {
		case 1:
			return DefaultAuth
	}

	return nil
}
