package cnf

import "regexp"

var (
	Logs = map[string]string{
		`F001`: ``,
	}
	Host string = "https://www.youtube.com"

	SHost string = "https://www.youtube.com/results"

	RegSome *regexp.Regexp = regexp.MustCompile(`"videoId":"([A-z-\d]+)".*?title":{"runs":\[{"text":"(.*?)"}].*?"longBylineText":{"runs":\[{"text":"(.*?)","navigationEndpoint.*?seconds"}}."simpleText":"([\d:]+).*?viewCountText":{"simpleText":"([\d,]+)`)
)
