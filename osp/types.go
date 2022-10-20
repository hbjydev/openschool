package osp

type HandlerFunc func(request *Request) (Response, error)
