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
