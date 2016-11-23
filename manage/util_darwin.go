package manage

import (
	"fmt"
	"path/filepath"
)

func connectorDirectory(homeDir, uuid string) string {
	return filepath.Join(userServiceDirectory(homeDir), uuid)
}

func serviceName(uuid string) string {
	return fmt.Sprintf("com.octoblu.%s", uuid)
}

func userServiceDirectory(homeDir string) string {
	return filepath.Join(homeDir, "Library", "MeshbluConnectors")
}
