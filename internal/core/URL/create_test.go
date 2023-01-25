package URL

import (
	"ShortLink/internal/core/link"
	"ShortLink/internal/models"
	"ShortLink/internal/store"
	mock_store "ShortLink/internal/store/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestDBCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_store.NewMockUrler(ctrl)
	uid := uuid.New()

	testTable := []struct {
		name          string
		store         *store.Store
		call          *gomock.Call
		url           models.URLStorage
		expectedError error
	}{
		{
			name:  "Success create",
			store: store.NewTest(nil, mu),
			url: models.URLStorage{uid, link.NewShortUrl([]rune(`https://www.google.com/search?q=go&sxsrf=
				AJOqlzWxF6ypcf7s2nM7n5M-22z0iZ87dw:1674586757635&source=lnms&tbm=isch&sa=X&ved=
				2ahUKEwjgjtbm8eD8AhWulosKHTPNA9gQ_AUoAXoECAEQAw&biw=1548&bih=882&dpr=1.59#imgrc=AfPvlfaD_83mKM`)),
				`https://www.google.com/search?q=go&sxsrf=AJOqlzWxF6ypcf7s2nM7n5M-22z0iZ87dw:1674586757635&source=
				lnms&tbm=isch&sa=X&ved=2ahUKEwjgjtbm8eD8AhWulosKHTPNA9gQ_AUoAXoECAEQAw&biw=1548&bih=882&dpr=1.59#imgrc=
				AfPvlfaD_83mKM`},
			call: mu.EXPECT().Create(&models.URLStorage{uid, link.NewShortUrl([]rune(`https://www.google.com/search?q=go&sxsrf=
				AJOqlzWxF6ypcf7s2nM7n5M-22z0iZ87dw:1674586757635&source=lnms&tbm=isch&sa=X&ved=
				2ahUKEwjgjtbm8eD8AhWulosKHTPNA9gQ_AUoAXoECAEQAw&biw=1548&bih=882&dpr=1.59#imgrc=AfPvlfaD_83mKM`)),
				`https://www.google.com/search?q=go&sxsrf=AJOqlzWxF6ypcf7s2nM7n5M-22z0iZ87dw:1674586757635&source=
				lnms&tbm=isch&sa=X&ved=2ahUKEwjgjtbm8eD8AhWulosKHTPNA9gQ_AUoAXoECAEQAw&biw=1548&bih=882&dpr=1.59#imgrc=
				AfPvlfaD_83mKM`}),
			expectedError: nil,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			err := DBCreate(subtest.store, subtest.url)
			if err != nil && subtest.expectedError != nil {
				t.Errorf("expected error (%v), got error (%v)", subtest.expectedError, err)
			} else {
				if err == nil && subtest.expectedError == nil {
					return
				}
				t.Errorf("expected error (%v), got error (%v)", subtest.expectedError, err)
			}
		})
	}
}
