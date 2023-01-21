package req

import (
	"fmt"
	"net/url"
)

type rUrl struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the url params
	*/

	href    string
	qParams map[string]string
}

func (u rUrl) Build() string {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Build the url by params
	*/

	q := url.Values{}
	for k, v := range u.qParams {
		q.Add(k, v)
	}

	return fmt.Sprintf("%s?%s", u.href, q.Encode())

}
