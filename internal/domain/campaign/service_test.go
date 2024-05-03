package campaign

import (
	"sendmail/internal/contract"
	"testing"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_NewCampaign(t *testing.T) {
	// assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	service := Service{Repository: repositoryMock}
	newCampaign := contract.NewCampaign{
		Name:    "Test X",
		Content: "Content Body",
		Emails:  []string{"test1@test.com"},
	}

	// assert.NotNil(newCampaign.Name)
	// assert.NotNil(newCampaign.Content)
	// assert.Greater(len(newCampaign.Emails), 0)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		}
		if campaign.Content != newCampaign.Content {
			return false
		}

		return true
	})).Return(nil)

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_SaveNewCampaign(t *testing.T) {
	newCampaign := contract.NewCampaign{
		Name:    "Test X",
		Content: "Content Body",
		Emails:  []string{"test1@test.com"},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
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
	service := Service{Repository: repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}
