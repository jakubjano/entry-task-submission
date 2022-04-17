package endpoints

import (
	"errors"
	"net/url"
	"strconv"
	"time"
)

//parseQueryParams is used to parse query parameters from http request
func parseQueryParams(params url.Values) (from, to, t int64, err error) {

	from, err = stringConv(params.Get("from"))
	if err != nil {
		return 0, 0, 0, err
	}

	to, err = stringConv(params.Get("to"))
	if err != nil {
		return 0, 0, 0, err
	}

	t, err = stringConv(params.Get("t"))
	if err != nil {
		return 0, 0, 0, err
	}

	return from, to, t, nil

}

// stringConv is used to convert string to int64
// returns converted integer value and error that is passed to handlers
func stringConv(s string) (i int64, err error) {

	i, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// validateInput is used to validate query parameters that are used in event aggregation.
// Takes time range and interval t from a request that are later used in event aggregation
// Basic scenarios where parsed unix timestamps are
func validateInput(from, to, t int64) (err error) {
	//todo implement methods on input struct {from, to, t} ?
	// chain errors ? wrapping ?
	if to-from < 0 {
		err = errors.New("to must be greater than from")
		return err
	} else if to-from == 0 {
		err = errors.New("please choose time range with a minimum length of 1 sec")
		return err
	} else if to-from < t {
		err = errors.New("time range from-to is too short for the given interval t")
		return err
	} else if to > time.Now().Unix() || from > time.Now().Unix() {
		err = errors.New("at least one of the interval boundaries was set to timestamp in the future")
		return err
	} else if t < 0 {
		err = errors.New("t must be greater than 0")
		return err
	} else if from < 0 {
		err = errors.New("from must be after 1970-01-01 00:00")
		return err
	}
	return nil
}
