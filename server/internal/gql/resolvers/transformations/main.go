package transformations

import (
	"errors"
	gql "github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/product"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/review"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils/date"
	stringUtil "github.com/fedoratipper/bitkiosk/server/pkg/utils/string"
	"github.com/jinzhu/gorm"
	"strings"
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

	if u.ID == 0 {
		return nil, errors.New("unable to find user")
	}

	if up.ID == 0 {
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

	if up.ID == 0 {
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

	if userProfile.ID == 0 {
		return nil, errors.New("unable to find user with email " + i.Email)
	}

	userProfile.FirstName = &i.FirstName
	userProfile.LastName = &i.LastName

	return &userProfile, nil
}

func GQLProductToDBProduct(i *gql.NewProduct, adminId uint) (*product.Product, error) {
	product := &product.Product{
		Name:             *stringUtil.FormatNameString(&i.Name),
		Sku:              *stringUtil.FormatWhiteSpace(&i.Sku),
		Stock:            i.Stock,
		AdminId:          adminId,
		Description: 	  strings.TrimSpace(i.Description),
		ShortDescription: strings.TrimSpace(i.ShortDescription),
		Price:            i.Price,
	}

	parsedStartDate, err := date.ParseSqlDate(i.StartDate)

	if err != nil {
		return nil, errors.New("Start date: " + err.Error())
	} else {
		product.StartDate = parsedStartDate
	}

	if i.EndDate != nil {
		parsedEndDate, err := date.ParseSqlDate(*i.EndDate)

		if err != nil {
			return nil, errors.New("End date: " + err.Error())
		} else {
			product.EndDate = parsedEndDate
		}
	}

	return product, nil
}

func DBProductToGQLProduct(p *product.Product, db *gorm.DB) (*gql.Product, error) {
	gqlProduct := &gql.Product{
		Sku:           		p.Sku,
		Name:          		p.Name,
		ShortDescription:   p.ShortDescription,
		Description:		p.Description,
		Price:          	p.Price,
		Stock:          	p.Stock,
		StartDate:      	*date.FormatToSqlDate(p.StartDate),
		EndDate:        	date.FormatToSqlDate(p.EndDate),
		CreatedAt:      	date.FormatToSqlDate(&p.CreatedAt),
		UpdatedAt:      	date.FormatToSqlDate(p.UpdatedAt),
	}

	rating, reviewCount :=  review.LoadAverageRatingAndRatingCountForProduct(p.ID, db)

	gqlProduct.Rating = rating
	gqlProduct.ReviewCount = reviewCount

	user := user.LoadUserWithId(p.AdminId, db)

	if user != nil {
		adminUser, err := DBUserToGQLUser(user)

		if err != nil {
			return nil, err
		}

		gqlProduct.CreatedByAdmin = adminUser
	}

	return gqlProduct, nil
}