package domain

import "errors"

var ErrorInvalidAccessToken = errors.New("invalid auth tokenizer")
var ErrorUserDoesNotExist = errors.New("user does not exist")
var ErrorUserAlreadyExists = errors.New("user already exists")
