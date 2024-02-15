package models

import (
	"fmt"
)

type Rating struct {
	Rating  int
	Comment string
}

type Author struct {
	Name  string
	Email string
}

type Feedback struct {
	Author  Author
	Ratings Ratings
}

type Ratings struct {
	Interesting         Rating
	Learning            Rating
	Pacing              Rating
	ExerciseDifficulty  Rating
	Support             Rating
	OverallSatisfaction Rating
}

func NewRating(ratingVal int, comment string) (rating *Rating, e error) {
	rating = &Rating{Rating: ratingVal, Comment: comment}
	e = ValidateRating(*rating)
	if e == nil {
		return
	} else {
		return nil, e
	}
}
func ValidateRating(rating Rating) error {
	if rating.Rating < 0 {
		return fmt.Errorf("invalid, rating < 0 not allowed, rating is %d", rating.Rating)
	} else if rating.Rating > 10 {
		return fmt.Errorf("invalid, rating > 10 not allowed, rating is %d", rating.Rating)
	} else if len(rating.Comment) > 200 {
		return fmt.Errorf("invalid, comments > 200 characters not allowed, length is %d", len(rating.Comment))
	} else {
		return nil
	}
}

func (r *Rating) ChangeRating(newVal int, newComment string) {
	r.Rating = newVal
	r.Comment = newComment
}
