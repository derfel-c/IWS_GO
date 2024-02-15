package models

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

func NewRating(ratingVal int, comment string) (rating Rating, message string) {
	rating = Rating{Rating: ratingVal, Comment: comment}
	message = ValidateRating(rating)
	if message == "Valid" {
		return
	} else {
		return Rating{}, "Rating not valid"
	}
}
func ValidateRating(rating Rating) string {
	if rating.Rating < 0 || rating.Rating > 10 || len(rating.Comment) > 200 {
		return "Invalid"
	} else {
		return "Valid"
	}
}
