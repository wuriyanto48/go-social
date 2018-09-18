package social

// Type generic type of social login
type Type int

const (
	// Facebook type
	Facebook Type = iota

	// Google Type
	Google

	// Linkedin Type
	Linkedin

	// Twitter Type
	Twitter

	// Github Type
	Github
)

// String function
// returns string of Type
func (t Type) String() string {
	switch t {
	case Facebook:
		return "Facebook"
	case Google:
		return "Google"
	case Linkedin:
		return "Linkedin"
	case Twitter:
		return "Twitter"
	case Github:
		return "Github"
	default:
		return "Facebook"
	}
}
