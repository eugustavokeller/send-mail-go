package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID       string
	Name     string
	CreateOn time.Time
	Content  string
	Contact  []Contact
}

func NewCampaign(name string, content string, emails []string) *Campaign {

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}

	return &Campaign{
		ID:       xid.New().String(),
		Name:     name,
		CreateOn: time.Now(),
		Content:  content,
		Contact:  contacts,
	}
}
