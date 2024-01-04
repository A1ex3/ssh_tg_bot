package main

import (
	"errors"
	"fmt"
	"net"
	"time"
)

type MessageTemplate struct{}

type template struct {
	DateTime time.Time
	RemoteIpPort string
	HostName  string
	UserName string
	RemoteIp string
}

func (messageTemplate *MessageTemplate) GetLocalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if messageTemplate.isPrivateIP(ip) {
				return ip, nil
			}
		}
	}

	return nil, errors.New("no IP")
}

func (*MessageTemplate) isPrivateIP(ip net.IP) bool {
	var privateIPBlocks []*net.IPNet
	for _, cidr := range []string{
		// don't check loopback ips
		//"127.0.0.0/8",    // IPv4 loopback
		//"::1/128",        // IPv6 loopback
		//"fe80::/10",      // IPv6 link-local
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
	} {
		_, block, _ := net.ParseCIDR(cidr)
		privateIPBlocks = append(privateIPBlocks, block)
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}

	return false
}

func (messageTemplate *MessageTemplate) Create(
	username string,
	dateTime time.Time,
	remoteIp string,
	remoteIpPort string,
) string {
	templ := &template{
		UserName: username,
		DateTime: dateTime,
		RemoteIp: remoteIp,
		RemoteIpPort: remoteIpPort,
	}

	var formatedDateTime string = templ.DateTime.UTC().Format(time.RFC3339)

	return fmt.Sprintf(
		"UserName: %s\nDateTime: %s\nRemoteIp: %s:%s",
		templ.UserName,
		formatedDateTime,
		templ.RemoteIp,
		templ.RemoteIpPort,
	)
}