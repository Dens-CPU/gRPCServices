package ordererrors

import "errors"

var (
	Avalible_markets = errors.New("change market is not available")
	Not_exist_order  = errors.New("the order is not exist!")
)
