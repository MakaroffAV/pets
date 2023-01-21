package req

func Get(u string, qPrms, hPrms map[string]string) ([]byte, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Do HTTP/GET request
	*/

	req := Request{
		Href:    &u,
		Meth:    "GET",
		ParamsQ: &qPrms,
		ParamsH: &hPrms,
	}

	r, err := req.Do()
	if err != nil {
		return nil, err
	}

}
