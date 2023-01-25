package link

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShortUrl(t *testing.T) {
	testTable := []struct {
		longUrl     []rune
		expectedRes string
	}{
		{
			longUrl:     []rune("https://www.youtube.com/watch?v=y6SLjY5Z1Os&ab_channel=JimmyManiglia"),
			expectedRes: "JRWLlBAMdt",
		},
		{
			longUrl:     []rune("https://www.youtube.com/watch?v=URurUv3NO14&t=200s&ab_channel=VariousArtists-Topic"),
			expectedRes: "XILOTnkljw",
		},
		{
			longUrl:     []rune("https://www.google.com/search?q=go&sxsrf=AJOqlzWUc1id05UWGJjP7nI3kjFaQOpasg:1674639402393&source=lnms&tbm=isch&sa=X&ved=2ahUKEwj3ltP1teL8AhXjkosKHZ4wBpoQ_AUoAXoECAEQAw&biw=1548&bih=882&dpr=1.59"),
			expectedRes: "qusvjXW3lI",
		},
	}

	for _, testCase := range testTable {
		res := NewShortUrl(testCase.longUrl)

		t.Logf("Calling NewShortUrl(%v), result %s\n", testCase.longUrl, testCase.expectedRes)

		assert.Equal(t, testCase.expectedRes, res,
			fmt.Sprintf("Incorrect result. Expect %s, got %s", testCase.expectedRes, res))
	}
}
