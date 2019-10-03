package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)
func main(){
	r:=RegisterHandler();
	http.ListenAndServe(":8000", r);
}

func RegisterHandler() * httprouter.Router {
	router:=httprouter.New();
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router

}