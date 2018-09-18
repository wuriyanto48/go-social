package facebook

// User data structure
type User struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Birthday string      `json:"birthday"`
	Gender   string      `json:"gender"`
	Picture  *Picture    `json:"picture,omitempty"`
	Error    interface{} `json:"error"`
}

// Picture structure
type Picture struct {
	Data struct {
		Height       int    `json:"height"`
		IsSilhouette bool   `json:"is_silhouette"`
		URL          string `json:"url"`
		Width        int    `json:"width"`
	} `json:"data"`
}
