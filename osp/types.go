package osp

type HandlerFunc func(request *OspRequest) (Response, error)
