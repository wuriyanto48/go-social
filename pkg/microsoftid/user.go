package microsoftid

// User Microsoft Identity Platform data structure
type User struct {
	OdataContext      string   `json:"@odata.context"`
	BusinessPhones    []string `json:"businessPhones"`
	DisplayName       string   `json:"displayName"`
	GivenName         string   `json:"givenName"`
	JobTitle          string   `json:"jobTitle"`
	Mail              string   `json:"mail"`
	MobilePhone       string   `json:"mobilePhone"`
	OfficeLocation    string   `json:"officeLocation"`
	PreferredLanguage string   `json:"preferredLanguage"`
	Surname           string   `json:"surname"`
	UserPrincipalName string   `json:"userPrincipalName"`
	ID                string   `json:"id"`
	Error             *Error   `json:"error,omitempty"`
}

// Error data structure
type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	InnerError struct {
		Date            string `json:"date"`
		RequestID       string `json:"request-id"`
		ClientRequestID string `json:"client-request-id"`
	} `json:"innerError"`
}
