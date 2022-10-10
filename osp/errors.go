package osp

import "errors"

var (
	ErrorNotEnoughRequestLineComponents = errors.New("not enough request line components")
	ErrorBadVersion                     = errors.New("an invalid osp version was provided")
	ErrorBadAction                      = errors.New("an invalid osp action was provided")
)
