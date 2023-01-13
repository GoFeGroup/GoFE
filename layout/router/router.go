package router

import "github.com/gin-gonic/gin"

type HandleFunc func(*gin.Context)
type MethodFunc func(string, ...gin.HandlerFunc) gin.IRoutes

type Router struct {
	Path       string
	Method     MethodFunc
	Target     HandleFunc
	SubRouters []Router
}

//
//type Routers struct {
//	R      []Router
//	engine *gin.Engine
//}
//
//func (routers *Routers) begin() *Routers {
//	if routers.engine != nil {
//		return routers
//	}
//	routers.engine = gin.Default()
//	pprof.Register(routers.engine, "debug/pprof")
//	return routers
//}
//
//func (routers *Routers) Static() *Routers {
//	r := routers.begin()
//	r.engine.Static("/", "")
//	return r
//}
//
//func (routers *Routers) Run(addr string) error {
//	r := routers.begin()
//	for _, v := range routers.R {
//		v.Method(v.Path, makeHandler(v.Target))
//	}
//	return r.Run(addr)
//}
//
//func makeHandler(handler HandleFunc) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		// check auth
//		handler(ctx)
//	}
//}
