package transformations

import (
	"errors"
	gql "github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils/date"
	"github.com/jinzhu/gorm"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *user.User) (o *gql.User, err error) {
	o = &gql.User{
		Email:     i.Email,
		Role:	   int(i.Role),
		CreatedAt: date.FormatToSqlDate(&i.CreatedAt),
		UpdatedAt: date.FormatToSqlDate(i.UpdatedAt),
	}

	if i.UserProfile != nil {
		o.UserProfile, err = DBUserProfileToGQLUserProfile(i.UserProfile, i)
	}

	return o, err
}

// GQLInputUserToDBUser transforms [user] gql input to db model
func GQLInputUserToDBUser(i *gql.NewUser, update bool) (o *user.User, err error) {
	o = &user.User{
		Email:     i.Email,
	}

	if i.Token == "" && !update {
		return nil, errors.New("Field [password] is required")
	}

	if i.Email == "" && !update {
		return nil, errors.New("Field [email] is required")
	}

	if i.Email != "" {
		o.Email = i.Email
	}

	return o, err
}

func DBUserProfileToGQLUserProfile(up *user.UserProfile, u *user.User) (o *gql.UserProfile, err error) {

	if u.Id == 0 {
		return nil, errors.New("unable to find user")
	}

	if up.Id == 0 {
		return nil, errors.New("unable to find user profile")
	}

	o = &gql.UserProfile{
		FirstName:   up.FirstName,
		LastName:    up.LastName,
		Email: 		 &u.Email,
	}

	return o, err
}


func UpdatedDBUserProfileToGQLUserProfile(up *user.UserProfile, email string) (o *gql.UserProfile, err error) {

	if up.Id == 0 {
		return nil, errors.New("unable to find user profile")
	}

	o = &gql.UserProfile{
		FirstName:   up.FirstName,
		LastName:    up.LastName,
		Email: 		 &email,
	}

	return o, err
}

func GQLUpdateUserProfileToDBUserProfile(i *gql.UpdatedProfile, db *gorm.DB) (*user.UserProfile, error) {

	if i.Email == "" {
		return nil, errors.New("unable to find user profile")
	}

	var userProfile user.UserProfile

	db = db.Joins("JOIN users on users.id = user_profiles.user_id").Where("users.email like ?", i.Email).Find(&userProfile)

	if userProfile.Id == 0 {
		return nil, errors.New("unable to find user with email " + i.Email)
	}

	userProfile.FirstName = &i.FirstName
	userProfile.LastName = &i.LastName

	return &userProfile, nil
}

