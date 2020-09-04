package matematica

import (
	"fmt"
	"strconv"
)

// media e reponsavel por calcular a media
func Media(numeors ...float64) float64 {
	total := 0.0
	for _, num := range numeors {
		total += num
	}

	media := total / float64(len(numeors))
	mediaArrendodada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArrendodada
}
