package manage

import (
	"testing"

	mocks "github.com/b-bianca/melichallenge/notify-api/adapter/handler/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestManageRegisterRoutes(t *testing.T) {
	u := mocks.NewApps(t)

	t.Run("when everything goes ok; should call apps", func(t *testing.T) {
		u.On("RegisterRoutes", mock.AnythingOfType("*gin.RouterGroup")).Return().Once()
		m := &Manage{
			user: u,
		}
		m.RegisterRoutes(nil)

		u.AssertExpectations(t)
	})
}

func TestNew(t *testing.T) {
	got := New(&UseCases{})
	assert.NotNil(t, got.user)
}
