package model

type AnimeInfo struct {
	Id       string `json:"id"`
	ImageUrl string `json:"image_url"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Year     string `json:"year"`
}
