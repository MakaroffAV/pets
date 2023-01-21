package req

var (
	methGet  string = "GET"
	methPost string = "POST"
)

func Get(u string, qParams, hParams map[string]string) ([]byte, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Do HTTP/GET request
	*/

	r := request{
		href:    &u,
		meth:    &methGet,
		hParams: &hParams,
		qParams: &qParams,
	}

	respBd, err := r.Do()
	if err != nil {
		return nil, err
	}

	return respBd, nil

}
