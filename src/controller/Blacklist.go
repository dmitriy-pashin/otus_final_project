package controller

import (
	"net"
	"net/http"
	"otus_final_project/src/component/web"
)

type Blacklist struct {
	addHandler    AddHandler
	deleteHandler DeleteHandler
}

type AddHandler interface {
	Add(ipNet *net.IPNet) (bool, error)
}

type DeleteHandler interface {
	Delete(ipNet *net.IPNet) (bool, error)
}

func NewBlacklist(addHandler AddHandler, deleteHandler DeleteHandler) *Blacklist {
	return &Blacklist{
		addHandler:    addHandler,
		deleteHandler: deleteHandler,
	}
}

func (cont *Blacklist) ActionAdd(response http.ResponseWriter, request *http.Request) web.ActionResultInterface {
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

func (cont *Blacklist) ActionDelete(response http.ResponseWriter, request *http.Request) web.ActionResultInterface {
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
