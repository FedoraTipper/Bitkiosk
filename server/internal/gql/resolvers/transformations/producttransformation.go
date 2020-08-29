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

func GQLProductToDBProduct(i *gql.NewProduct, adminId uint) (*product.Product, error) {
	dboProduct := &product.Product{
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
		dboProduct.StartDate = parsedStartDate
	}

	if i.EndDate != nil {
		parsedEndDate, err := date.ParseSqlDate(*i.EndDate)

		if err != nil {
			return nil, errors.New("End date: " + err.Error())
		} else {
			dboProduct.EndDate = parsedEndDate
		}
	}

	return dboProduct, nil
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

	rating, reviewCount :=  review.LoadAverageRatingAndRatingCountForProduct(p.Id, db)

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
