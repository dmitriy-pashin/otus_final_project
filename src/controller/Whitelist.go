package controller

import (
	"net"
	"net/http"
	"otus_final_project/src/component/web"
)

type Whitelist struct {
	addHandler    AddHandler
	deleteHandler DeleteHandler
}

func NewWhitelist(addHandler AddHandler, deleteHandler DeleteHandler) *Whitelist {
	return &Whitelist{
		addHandler:    addHandler,
		deleteHandler: deleteHandler,
	}
}

func (cont *Whitelist) ActionAdd(response http.ResponseWriter, request *http.Request) web.ActionResultInterface {
	ip := request.FormValue("ip")

	_, ipnet, err := net.ParseCIDR(ip)
	if err != nil {
		return web.NewActionResult(web.DefaultFailContent, http.StatusBadRequest, err)
	}

	_, err = cont.addHandler.Add(ipnet)

	if err != nil {
		return web.NewActionResult(web.DefaultFailContent, http.StatusInternalServerError, err)
	}

	return web.NewActionResult(web.DefaultSuccessContent, http.StatusOK, nil)
}

func (cont *Whitelist) ActionDelete(response http.ResponseWriter, request *http.Request) web.ActionResultInterface {
	ip := request.FormValue("ip")

	_, ipnet, err := net.ParseCIDR(ip)
	if err != nil {
		return web.NewActionResult(web.DefaultFailContent, http.StatusBadRequest, err)
	}

	_, err = cont.deleteHandler.Delete(ipnet)

	if err != nil {
		return web.NewActionResult(web.DefaultFailContent, http.StatusInternalServerError, err)
	}

	return web.NewActionResult(web.DefaultSuccessContent, http.StatusOK, nil)
}
