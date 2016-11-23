package manage

import (
	De "github.com/tj/go-debug"
)

var debug = De.Debug("meshblu-connector-service:manage")

// UninstallOptions are used to call Uninstall
type UninstallOptions struct {
	// HomeDir is the HomeDir of user
	HomeDir string

	// UUID is the UUID of the connector to uninstall
	UUID string

	// ServiceType is the ServiceType of the connector to uninstall
	// can be one of ["Service", "UserService", "UserLogin"]
	ServiceType string

	// ServiceUsername is the name of the user the service is registered
	// under
	ServiceUsername string

	// ServicePassword is a password of some kind. It's definetely nescessary
	// and this totally won't work without it ;-)
	ServicePassword string
}

// Uninstall uninstalls the meshblu-connector
// indicated by the uuid
// [x] Stop/deregister the service
// [ ] Remove all the directories
func Uninstall(options *UninstallOptions) error {
	err := deregisterService(options.UUID, options.ServiceType, options.ServiceUsername, options.ServicePassword)
	if err != nil {
		debug("Service deregistration failed, probably wasn't installed: %v", err.Error())
	}

	return removeDirectories(options.HomeDir, options.UUID)
}
