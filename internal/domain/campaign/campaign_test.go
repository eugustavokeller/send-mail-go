package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_Create(t *testing.T) {

	// Three A's: Arrange, Act, Assert

	// Arrange: setup variables and input
	assert := assert.New(t)
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

}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {

	assert := assert.New(t)
	name := "ID is Not Nil"
	content := "Body Content"
	contacts := []string{"t1@test.com", "t2@test.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateOnIsNotNill(t *testing.T) {

	assert := assert.New(t)
	name := "CreateOn is Not Nil"
	content := "Body Content"
	contacts := []string{"t1@test.com", "t2@test.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.CreateOn)
	assert.Greater(campaign.CreateOn, now)
}
