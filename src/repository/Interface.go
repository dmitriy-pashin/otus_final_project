package repository

import (
	"net"
)

type ListRepository interface {
	Add(ip *net.IPNet) (bool, error)
	Delete(ip *net.IPNet) (bool, error)
	IsInList(ip net.IP) bool
}
