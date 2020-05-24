package store

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"net"
)

type Store struct {
	client *consul.Client
}

var (
	Ip = localIP()
	S *Store
)

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func (s *Store)SetUp() error {
	config := consul.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := consul.NewClient(config)
	if err != nil {
		return fmt.Errorf("no consul client newed %s", err)
	}

	port := 8080
	registerInfo := &consul.AgentServiceRegistration{
		ID:fmt.Sprintf("test-server-%s-%d", Ip, port),
		Name:fmt.Sprintf("test-web-%s", Ip),
		Port:port,
		Address:Ip,
		Tags:make([]string, 0),
	}

	if err := client.Agent().ServiceRegister(registerInfo); err != nil {
		return err
	}
	S.client = client
	return nil
}

func (s *Store) Destroy()  {
	s.client.Agent().ServiceDeregister(fmt.Sprintf("test-web-%s", Ip))
}
