package ipacl

import (
	"net"
)

type ipEntry struct {
	ipaddr net.IP
	ipnet  *net.IPNet
}

type iplistImpl struct {
	list []*ipEntry
}

func (self *iplistImpl) add(ip net.IP, ipnet *net.IPNet) {
	entry := &ipEntry{
		ipaddr: ip,
		ipnet:  ipnet,
	}
	self.list = append(self.list, entry)
}

func (self *iplistImpl) AddEntry(ipstr ...string) error {
	for _, item := range ipstr {
		if ip, ipnet, err := net.ParseCIDR(item); err != nil {
			ip = net.ParseIP(item)
			if ip == nil {
				return err
			} else {
				self.add(ip, nil)
			}
		} else {
			self.add(ip, ipnet)
		}
	}
	return nil
}

func (self *iplistImpl) Contains(ipstr string) bool {
	if ip := net.ParseIP(ipstr); ip != nil {
		for _, entry := range self.list {
			if entry.ipnet != nil {
				if entry.ipnet.Contains(ip) {
					return true
				}
			} else if entry.ipaddr.Equal(ip) {
				return true
			}
		}
	}
	return false
}

func New() IPListMgmt {
	return new(iplistImpl)
}
