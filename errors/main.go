package main

import "fmt"

// lets make some structs

type Teacher struct {
	name                 string
	rateMyProfessorScore float64
}

// lets make a simple method
func (t *Teacher) AssignScoreFromReviews(reviewScores []int) error {

	total := 0
	reviewCount := len(reviewScores)
	// go loops
	for _, review := range reviewScores {

		if review < 0 {
			return fmt.Errorf("Review less than zero -> %d", review)
		}
		total += review
	}

	t.rateMyProfessorScore = float64(total) / float64(reviewCount)
	//RETURN NILL
	return nil
}

func main() {

	var goodReviews []int = []int{5, 5, 5, 5, 5, 5}
	var badReviews []int = []int{-100, 5, 5, 5, 5, 5}

	david := &Teacher{name: "David"}
	matt := &Teacher{name: "Matt"}

	if err := david.AssignScoreFromReviews(goodReviews); err != nil {
		panic(err)
	}

	if err := matt.AssignScoreFromReviews(badReviews); err != nil {
		panic(err)
	}

}
