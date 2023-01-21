package req

var (
	methGet string = "GET"
)

func Get(u string, qPrms, hPrms map[string]string) ([]byte, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Do HTTP/GET request
	*/

	r := Request{
		Href: &u,
		Meth: &methGet,
	}

	resp, err := r.

	return nil, nil

}
