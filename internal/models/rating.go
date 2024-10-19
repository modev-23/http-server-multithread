package models

type Rating struct {
	RatingId int     `json:"ratingId"`
	UserId   int     `json:"userId"`
	MovieId  int     `json:"movieId"`
	Rating   float32 `json:"rating"`
}
