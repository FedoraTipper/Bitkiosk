package review

import (
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/jinzhu/gorm"
)

func LoadAverageRatingAndRatingCountForProduct(productID uint, db *gorm.DB) (float64, int) {
	var averageRating float64
	var totalReviewCount int

	err := db.New().Raw("select count(id) as totalReviewCount, coalesce(avg(rating), 0) as averageRating from reviews r where r.product_id = ?", productID).Row().Scan(&totalReviewCount, &averageRating)

	if err != nil {
		logger.Errorfn("LoadAverageRatingAndRatingCountForProduct", err)
	}

	return averageRating, totalReviewCount
}

func LoadReviewsForProduct(productId uint, db *gorm.DB) []Review {
	var reviews []Review

	db.Where("product_id = ?", productId).Find(&reviews)

	return reviews
}