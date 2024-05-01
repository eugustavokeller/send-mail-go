package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Test campaign"
	content  = "Body Content"
	contacts = []string{"t1@test.com", "t2@test.com"}
)

func Test_NewCampaign_Create(t *testing.T) {

	// Three A's: Arrange, Act, Assert

	// Arrange: setup variables and input
	assert := assert.New(t)

	// Act: execute the code to test
	campaign, _ := NewCampaign(name, content, contacts)

	// Assert: check the output
	assert.NotNil(campaign)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Len(campaign.Contact, 2)
	assert.Equal(contacts[0], campaign.Contact[0].Email)

}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateOnMustBeNow(t *testing.T) {

	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)
	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.CreateOn)
	assert.Greater(campaign.CreateOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required", err.Error())
}
