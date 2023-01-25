package URL

import (
	"ShortLink/internal/core"
	"ShortLink/internal/models"
	"ShortLink/internal/store"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestGetByShortLinkIM(t *testing.T) {

	u := []models.URLStorage{{uuid.New(), "kj5_6b76Z3", "https://www.last.fm/ru/"},
		{uuid.New(), "JRWLlBAMdt", "https://www.youtube.com/watch?v=y6SLjY5Z1Os&ab_channel=JimmyManiglia"},
		{uuid.New(), "TT28kjdZY3", "https://habr.com/ru/all/"}}
	s := store.NewIM(zap.L().Sugar())
	s.IMStorage = u

	testTable := []struct {
		store            *store.Store
		shortLink        string
		expectedLongLink string
		expectedError    error
	}{
		{
			store:            s,
			shortLink:        "TT28kjdZY3",
			expectedLongLink: "https://habr.com/ru/all/",
			expectedError:    nil,
		},
		{
			store:            s,
			shortLink:        "TG48kj_ZY3",
			expectedLongLink: "",
			expectedError:    core.ErrNotFound,
		},
	}

	type pair struct {
		LongLink string
		Error    error
	}

	for _, testCase := range testTable {
		resUrl, err := GetByShortLinkIM(testCase.store, testCase.shortLink)
		if resUrl == nil {
			resUrl = &models.URLStorage{LongURL: ""}
		}

		t.Logf("Calling GetByShortLinkIM(%v, %v), result %s\n", testCase.store, testCase.shortLink, resUrl)

		assert.Equal(t, pair{LongLink: testCase.expectedLongLink, Error: testCase.expectedError},
			pair{LongLink: resUrl.LongURL, Error: err}, fmt.Sprintf("Incorrect result. Expect %s, %e, got %s, %e",
				testCase.expectedLongLink, testCase.expectedError, resUrl.LongURL, err))
	}
}
