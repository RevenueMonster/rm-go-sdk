package sdk

type path struct {
	method  string
	urlPath string
}

func newPath(method, urlpath string) path {
	return path{method, urlpath}
}
