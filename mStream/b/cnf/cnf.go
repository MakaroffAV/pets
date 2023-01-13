package cnf

import "regexp"

var (
	DbPath string = "/Users/makarov/Desktop/pets/mStream/b/mStream.db"
	Logs          = map[string]string{
		`F0000001`: `Error while creating new connection with database`,
	}
	Host string = "https://www.youtube.com"

	SHost string = "https://www.youtube.com/results"

	RegSome *regexp.Regexp = regexp.MustCompile(`"videoId":"([A-z-\d]+)".*?title":{"runs":\[{"text":"(.*?)"}].*?"longBylineText":{"runs":\[{"text":"(.*?)","navigationEndpoint.*?seconds"}}."simpleText":"([\d:]+).*?viewCountText":{"simpleText":"([\d,]+)`)
)
