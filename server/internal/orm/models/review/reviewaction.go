package review

import (
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/jinzhu/gorm"
)

const (
	noOffset int = 0
	noLimit  int = 0
)

func LoadTotalReviewCountForProduct(productId uint, db *gorm.DB) int {
	var totalReviewCount int

	err := db.New().Raw("select count(id) as totalReviewCount from reviews r where r.product_id = ?", productId).Row().Scan(&totalReviewCount)

	if err != nil {
		logger.Errorfn("LoadTotalReviewCountForProduct", err)
	}

	return totalReviewCount
}

func LoadAverageRatingForProduct(productId uint, db *gorm.DB) float64 {
	var averageRating float64

	err := db.New().Raw("select coalesce(round(avg(rating), 2), 0) as averageRating from reviews r where r.product_id = ?", productId).Row().Scan(&averageRating)

	if err != nil {
		logger.Errorfn("LoadAverageRatingForProduct", err)
	}

	return averageRating
}

func LoadAverageRatingAndRatingCountForProduct(productId uint, db *gorm.DB) (float64, int) {
	averageRating := LoadAverageRatingForProduct(productId, db)
	totalReviewCount  := LoadTotalReviewCountForProduct(productId, db)

	return averageRating, totalReviewCount
}

func LoadReviewsForProduct(productId uint, db *gorm.DB) []Review {
	return LoadReviewsForProductWithLimitAndOffset(productId, noLimit, noOffset, db)
}

func LoadReviewsForProductWithLimitAndOffset(productId uint, limit int, offset int, db *gorm.DB) []Review {
	var reviews []Review

	db.Where("product_id = ?", productId).Limit(limit).Offset(offset).Find(&reviews)

	return reviews
}

func LoadProductReviewForUser(productId uint, userId uint, db *gorm.DB) Review {
	var review Review

	db.Where("product_id = ? and user_id = ?", productId, userId).First(&review)

	return review
}