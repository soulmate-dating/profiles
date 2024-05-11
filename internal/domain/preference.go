package domain

const (
	Anyone Preference = "anyone"
	Man    Preference = "man"
	Woman  Preference = "woman"
)

type Preference string

func (s Preference) Preferences() (string, string) {
	if s == Anyone {
		return string(Man), string(Woman)
	}
	return string(s), string(s)
}
