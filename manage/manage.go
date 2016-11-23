package manage

import (
	"os"

	"github.com/kardianos/service"
)

func deregisterService(uuid, serviceType, serviceUsername, servicePassword string) error {
	svcConfig := &service.Config{
		Name: serviceName(uuid),
		Option: service.KeyValue{
			"UserService": true,
			"KeepAlive":   true,
			"UserName":    serviceUsername,
			"Password":    servicePassword,
		},
	}

	prg := &Program{}
	svc, err := service.New(prg, svcConfig)
	if err != nil {
		return err
	}

	return svc.Uninstall()
}

func removeDirectories(homeDir, uuid string) error {
	return os.RemoveAll(connectorDirectory(homeDir, uuid))
}
