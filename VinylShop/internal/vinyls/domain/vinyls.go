package domain

import (
	"fmt"
	"time"
)

type Vinyl struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Artist      string    `json:"artist"`
	ReleaseDate int32     `json:"releaseDate"`
	Price       int32     `json:"price"`
	Rating      float32   `json:"rating"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewVinyl(title, artist string, releaseDate, price int32, rating float32) *Vinyl {
	return &Vinyl{
		Title:       title,
		Artist:      artist,
		ReleaseDate: releaseDate,
		Price:       price,
		Rating:      rating,
	}
}

func (v *Vinyl) Validate() error {
	if len(v.Title) < 1 {
		return fmt.Errorf("field title must not be empty")
	}
	if len(v.Artist) < 1 {
		return fmt.Errorf("field artst must not be empty")
	}
	if v.ReleaseDate < 1000 && v.ReleaseDate > 2024 {
		return fmt.Errorf("field year must be between 1000 and 2024")
	}
	return nil
}
