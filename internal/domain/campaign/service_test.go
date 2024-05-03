package campaign

import (
	"errors"
	"sendmail/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test X",
		Content: "Content Body",
		Emails:  []string{"test1@test.com"},
	}
	service = Service{}
)

func Test_Create_NewCampaign(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_SaveNewCampaign(t *testing.T) {

	repository := new(repositoryMock)

	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		}
		if campaign.Content != newCampaign.Content {
			return false
		}
		if len(campaign.Contact) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = repository

	service.Create(newCampaign)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositoryMock)

	repository.On("Save", mock.Anything).Return(errors.New("Error to save on database"))
	service.Repository = repository

	_, err := service.Create(newCampaign)

	assert.Equal("Error to save on database", err.Error())
}
