package utils

import "math"

func CalculateTotalPages(totalResults int64, limitPerPage int) int64 {
	if limitPerPage <= 0 {
		return 0
	}

	return int64(math.Ceil(float64(totalResults) / float64(limitPerPage)))
}
