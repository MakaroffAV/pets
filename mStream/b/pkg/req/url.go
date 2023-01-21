package req

import (
	"fmt"
	"net/url"
)

type Url struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the url params
	*/

	Href string
	Qprm map[string]string
}

func (u Url) Build() string {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Build the url by params
	*/

	q := url.Values{}
	for k, v := range u.Qprm {
		q.Add(k, v)
	}

	return fmt.Sprintf("%s?%s", u.Href, q.Encode())

}
