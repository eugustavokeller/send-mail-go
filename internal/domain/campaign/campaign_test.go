package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {

	// Three A's: Arrange, Act, Assert

	// Arrange: setup variables and input
	assert := assert.New(t)
	id := "123"
	name := "New Campaign"
	content := "Body Campaign"
	contacts := []string{"email1@campaign.com", "email2@campaign.com"}

	// Act: execute the code to test
	campaign := NewCampaign(name, content, contacts)

	// Assert: check the output
	assert.NotNil(campaign)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Len(campaign.Contact, 2)
	assert.Equal(contacts[0], campaign.Contact[0].Email)
	assert.NotEqual(id, campaign.ID)
}
