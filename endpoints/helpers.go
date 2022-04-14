package endpoints

import (
	"net/url"
	"strconv"
)

func ParseQueryParams(params url.Values) (from, to, t int64, err error) {

	from, err = StringConv(params.Get("from"))
	if err != nil {
		return 0, 0, 0, err
	}

	to, err = StringConv(params.Get("to"))
	if err != nil {
		return 0, 0, 0, err
	}

	t, err = StringConv(params.Get("t"))
	if err != nil {
		return 0, 0, 0, err
	}

	//todo validate these query params
	// validate here or in function implementing logic of aggregation ?

	return from, to, t, nil

}

func StringConv(s string) (i int64, err error) {

	i, err = strconv.ParseInt(s, 10, 64)
	//todo error handling
	if err != nil {
		return 0, err
	}
	return i, nil
}
