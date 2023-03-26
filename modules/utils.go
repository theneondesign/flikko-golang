package modules

import (
	"flikko/models"
	"math"
	"sort"

	"gorm.io/gorm"
)

func getNearestProducts(db *gorm.DB, lat, lng, maxDistance float64, limit int) ([]models.Product, error) {

	// Calculate the bounding box of the user's location
	minLat, maxLat, minLng, maxLng := getBoundingBox(lat, lng, maxDistance)

	// Query the database for products within the bounding box
	var products []models.Product
	err := db.Where("latitude BETWEEN ? AND ?", minLat, maxLat).
		Where("longitude BETWEEN ? AND ?", minLng, maxLng).
		Limit(limit).
		Find(&products).Error
	if err != nil {
		return nil, err
	}

	// Sort the products by their distance to the user
	sort.Slice(products, func(i, j int) bool {
		return haversineDistance(lat, lng, products[i].Latitude, products[i].Longitude) < haversineDistance(lat, lng, products[j].Latitude, products[j].Longitude)
	})

	// Limit the number of results
	if len(products) > limit {
		products = products[:limit]
	}

	return products, nil
}

func getBoundingBox(lat, lng, radius float64) (float64, float64, float64, float64) {
	const earthRadius = 6371 // in kilometers

	// Convert radius from kilometers to radians
	r := radius / earthRadius

	// Calculate bounding box coordinates
	minLat := lat - r
	maxLat := lat + r
	minLng := lng - r
	maxLng := lng + r

	return minLat, maxLat, minLng, maxLng
}

func haversineDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadius = 6371 // in kilometers

	// Convert latitudes and longitudes from degrees to radians
	lat1Rad := degreesToRadians(lat1)
	lng1Rad := degreesToRadians(lng1)
	lat2Rad := degreesToRadians(lat2)
	lng2Rad := degreesToRadians(lng2)

	// Calculate the differences between the latitudes and longitudes
	latDiff := lat2Rad - lat1Rad
	lngDiff := lng2Rad - lng1Rad

	// Calculate the great-circle distance using the Haversine formula
	a := math.Pow(math.Sin(latDiff/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(lngDiff/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
