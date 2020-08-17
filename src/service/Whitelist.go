package service

import (
	"net"
	"otus_final_project/src/repository"
)

type Whitelist struct {
	rep *repository.Whitelist
}

func NewWhitelist(rep *repository.Whitelist) *Whitelist {
	return &Whitelist{rep: rep}
}

func (cont *Whitelist) Add(ip *net.IPNet) (bool, error) {
	res, err := cont.rep.Add(ip)

	return res, err
}

func (cont *Whitelist) Delete(ip *net.IPNet) (bool, error) {
	res, err := cont.rep.Delete(ip)

	return res, err
}
