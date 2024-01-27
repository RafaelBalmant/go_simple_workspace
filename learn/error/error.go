package error

import (
	"fmt"
	"net/http"
)

func HandlerError(executeWithError bool) (resp *http.Response, err error) {
	var url string
	if executeWithError {
		url = "http://www.123ddasccc.com/robots.txt"
	} else {
		url = "http://www.google.com/robots.txt"
	}

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return res, nil
	}

	return res, nil
}
