package xtime

// Pattern ...
type Pattern string

// ...
const (
	Layout       Pattern = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC        Pattern = "Mon Jan _2 15:04:05 2006"
	UnixDate     Pattern = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate     Pattern = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822       Pattern = "02 Jan 06 15:04 MST"
	RFC822Z      Pattern = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850       Pattern = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123      Pattern = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z     Pattern = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339      Pattern = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano  Pattern = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen      Pattern = "3:04PM"
	Stamp        Pattern = "Jan _2 15:04:05"
	StampMilli   Pattern = "Jan _2 15:04:05.000"
	StampMicro   Pattern = "Jan _2 15:04:05.000000"
	StampNano    Pattern = "Jan _2 15:04:05.000000000"
	DateTime     Pattern = "2006-01-02 15:04:05"
	DateOnly     Pattern = "2006-01-02"
	TimeOnly     Pattern = "15:04:05"
	DateJoin     Pattern = "20060102"
	DateTimeJoin Pattern = "20060102150405"
	TimeJoinOnly Pattern = "150405"
)

// Desc ...
func (t Pattern) Desc() string {
	switch t {
	case Layout, ANSIC, UnixDate, RubyDate, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, RFC3339, RFC3339Nano, Kitchen, Stamp, StampMilli, StampMicro, StampNano, DateTime, DateOnly, TimeOnly, DateJoin, DateTimeJoin, TimeJoinOnly:
		return string(t)
	default:
		return string(DateTime)
	}
}
