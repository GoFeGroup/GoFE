package router

type Router struct {
	Path       string
	Target     string
	SubRouters []Router
}
