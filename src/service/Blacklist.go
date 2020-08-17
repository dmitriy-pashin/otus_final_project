package service

import (
	"net"
	"otus_final_project/src/repository"
)

type Blacklist struct {
	rep *repository.Blacklist
}

func NewBlacklist(rep *repository.Blacklist) *Blacklist {
	return &Blacklist{rep: rep}
}

func (cont *Blacklist) Add(ip *net.IPNet) (bool, error) {
	res, err := cont.rep.Add(ip)

	return res, err
}

func (cont *Blacklist) Delete(ip *net.IPNet) (bool, error) {
	res, err := cont.rep.Delete(ip)

	return res, err
}
