package controller

import (
	"net"
	"net/http"
	"otus_final_project/src/component/web"
)

type Login struct {
	loginHandler LoginHandler
}

type LoginHandler interface {
	Login(ipNet net.IP, login string, password string) bool
}

func NewLogin(loginHandler LoginHandler) *Login {
	return &Login{
		loginHandler: loginHandler,
	}
}

func (cont *Login) ActionLogin(response http.ResponseWriter, request *http.Request) web.ActionResultInterface {
	ip := request.FormValue("ip")
	login := request.FormValue("login")
	password := request.FormValue("password")

	address, _, err := net.ParseCIDR(ip)
	if err != nil {
		return web.NewActionResult(web.DefaultFailContent, http.StatusBadRequest, err)
	}

	result := cont.loginHandler.Login(address, login, password)

	if !result {
		return web.NewActionResult(web.DefaultFailContent, http.StatusOK, nil)
	}

	return web.NewActionResult(web.DefaultSuccessContent, http.StatusOK, nil)
}
