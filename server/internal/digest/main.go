package digest

import (
	"golang.org/x/crypto/bcrypt"
)


func bcryptDigest(secret string) (string){
	digest, _ := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	return string(digest)
}

func bcryptCompare(secret string, dbToken string) (bool){
	return bcrypt.CompareHashAndPassword([]byte(dbToken), []byte(secret)) == nil
}

func GetDigest(secret string, authMethodID uint) (string){
	switch authMethodID {
		case 1:
			return bcryptDigest(secret)
	}

	return ""
}

func CompareDigest(secret string, dbToken string, authMethodId uint) (bool) {
	switch authMethodId{
		case 1:
			return bcryptCompare(secret, dbToken)
	}

	return false
}