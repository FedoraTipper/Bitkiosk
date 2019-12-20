package models

type AuthMethod struct {
	ID       int
	TTL      int  // minutes to refresh
}

var DefaultAuth = &AuthMethod{
	ID:  1,
	TTL: 3600,
}

func GetAuthMethod(ID int) *AuthMethod {
	switch ID {
		case 1:
			return DefaultAuth
	}

	return nil
}
