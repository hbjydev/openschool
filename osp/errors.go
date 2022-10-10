package osp

import "errors"

var (
	ErrorBadVersion = errors.New("an invalid osp version was provided")
	ErrorBadAction  = errors.New("an invalid osp action was provided")
)
