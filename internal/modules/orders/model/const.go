package model

import "errors"

var ErrorProductNotFound = errors.New("product is not found")
var ErrorProductNotEnoughStock = errors.New("product does not have enough stock")
