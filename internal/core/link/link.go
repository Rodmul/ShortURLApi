package link

var (
	available = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_1234567890")
)

func NewShortUrl(longUrl []rune) string {
	var res []rune
	for i := len(longUrl) - 1; i >= len(longUrl)-10; i-- {
		temp := 0
		for k := 0; k < 4; k++ {
			temp += int(longUrl[i-k])
		}
		res = append(res, available[temp%63])
	}

	return string(res)
}
