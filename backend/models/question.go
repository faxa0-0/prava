package models

type Question struct {
	ID          int64    `json:"id"`
	Text        string   `json:"text"`
	Answer      int      `json:"answer"`
	Options     []string `json:"options"`
	Explanation string   `json:"explanation"`
	ImgURL      string   `json:"img_url"`
}
