package req

import (
	"testing"
)

type caseTestUrlBuild struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the case
	*/

	arg rUrl
	exp string
}

var (
	testCasesUrlBuild = []caseTestUrlBuild{
		{
			arg: rUrl{
				href: "https://yandex.ru/search/",
				qParams: map[string]string{
					"arg1": "golang for web",
				},
			},
			exp: "https://yandex.ru/search/?arg1=golang+for+web",
		},
		{
			arg: rUrl{
				href: "https://yandex.ru/search/",
				qParams: map[string]string{
					"arg1": "golang for web",
					"arg2": "duck&duck2gooo",
				},
			},
			exp: "https://yandex.ru/search/?arg1=golang+for+web&arg2=duck%26duck2gooo",
		},
	}
)

func TestUrlBuild(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Build the url by params"
	*/

	for _, c := range testCasesUrlBuild {
		if exp := c.arg.Build(); exp != c.exp {
			t.Fatalf(
				`
					Test failed:	Test "Build the url by params"
									(actual and expected values don't match) 	
				`,
			)
		}
	}

}
