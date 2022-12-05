package definition

import (
	"time"
)

type Author struct {
	FirstName string
	LastName  string
}

type Articles struct {
	Title        string
	Authors      []Author
	Content      string
	Draft        bool
	Category     uint16
	DateCreated  time.Time
	DateModefied time.Time
}
