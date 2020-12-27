package types

type Response struct {
	Response struct{
		ID int `json:"id"`
		Title string `json:"title"`
		Body string `json:"body"`
		Date string `json:"date"`
	} `json:"response"`
}