package main

import (
	"fmt"

	"github.com/insabelter/IWS_GO/models"
)

// Bitte die folgenden beiden Structs in das eigene package "models" auslagern und weitere Structs hinzufügen

func main() {
	var rating = models.Rating{Rating: 1, Comment: "test"}
	var author = models.Author{Name: "Max", Email: "max@max.de"}
	var feedback = models.Feedback{Author: author, Ratings: models.Ratings{Interesting: rating, Learning: rating, Pacing: rating, ExerciseDifficulty: rating, Support: rating, OverallSatisfaction: rating}}
	fmt.Println(feedback)
	var validRating, _ = models.NewRating(2, "comment")
	var invalidRating, e = models.NewRating(23, "comment")
	fmt.Println(*validRating)
	fmt.Println(invalidRating)
	fmt.Println(e)
}
