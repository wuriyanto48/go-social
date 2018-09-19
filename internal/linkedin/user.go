package linkedin

// User structure
type User struct {
	EmailAddress               string `json:"emailAddress"`
	FirstName                  string `json:"firstName"`
	Headline                   string `json:"headline"`
	ID                         string `json:"id"`
	LastName                   string `json:"lastName"`
	NumConnections             int    `json:"numConnections"`
	PictureURL                 string `json:"pictureUrl"`
	SiteStandardProfileRequest struct {
		URL string `json:"url"`
	} `json:"siteStandardProfileRequest"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
	Status    int    `json:"status"`
	Timestamp int64  `json:"timestamp"`
}
