package osp

type Resource struct {
	Name string

	GET    HandlerFunc
	LIST   HandlerFunc
	UPDATE HandlerFunc
	CREATE HandlerFunc
	DELETE HandlerFunc
}
