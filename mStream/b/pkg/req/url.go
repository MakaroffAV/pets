package req

import "net/url"

type Url struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the url params
	*/

	Prcl *string
	Host *string
}

func (u Url) Build() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Build the url by params
	*/

	u := url.URL{}
}
