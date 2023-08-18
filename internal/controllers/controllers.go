package controllers

import "github.com/julienschmidt/httprouter"

type Controller interface {
	Register(router *httprouter.Router)
}
