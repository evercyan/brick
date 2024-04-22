package xavatar

// Style ...
type Style int

const (
	StyleLetter Style = iota
	StyleSquare
	StylePornhub
)

func (t Style) String() string {
	switch t {
	case StyleLetter:
		return "letter"
	case StyleSquare:
		return "square"
	case StylePornhub:
		return "pornhub"
	default:
		return ""
	}
}
