package errorutil

import "errors"

var GENERAL_BAD_REQUEST = errors.New("bad request")
var GENERAL_NOT_FOUND = errors.New("not found")
