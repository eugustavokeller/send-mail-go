package campaign

import (
	"testing"
)

func TestNewCampaign(t *testing.T) {

	name := "New Campaign"
	content := "Body Campaign"
	contacts := []string{"email1@campaign.com", "email2@campaign.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID == "" {
		t.Error("ID must not be empty")
	}
	if campaign.CreateOn.IsZero() {
		t.Error("CreateOn must not be empty")
	}
	if campaign.Name != name {
		t.Error("Name must be", name)
	}
	if campaign.Content != content {
		t.Error("Content must be", content)
	}
	if len(campaign.Contact) != len(contacts) {
		t.Error("Contact length must be", len(contacts))
	}

}
