package utils

import (
	"context"
	"fmt"
	"github.com/containers/libpod/v2/pkg/bindings/images"
	"github.com/containers/libpod/v2/pkg/domain/entities"
	"os"

	//"github.com/containers/libpod/v2/libpod/define"
	"github.com/containers/libpod/v2/pkg/bindings"
	//"github.com/containers/libpod/v2/pkg/bindings/containers"
	//"github.com/containers/libpod/v2/pkg/bindings/images"
	//"github.com/containers/libpod/v2/pkg/domain/entities"
	//"github.com/containers/libpod/v2/pkg/specgen"
)

func getConnection() (context.Context, error) {
	// Get Podman socket location
	sock_dir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sock_dir + "/podman/podman.sock"

	// Connect to Podman socket
	return bindings.NewConnection(context.Background(), socket)
}

func PullImage(image string) error {
	conn, err := getConnection()
	if (err != nil) {
		fmt.Errorf("Error pulling image with podman %s. %s", image, err)
		return err
	}

	// Pull Busybox image (Sample 1)
	fmt.Println("Attempt to pull image with podman ", image)
	_, err = images.Pull(conn, image, entities.ImagePullOptions{})
	if err != nil {
		fmt.Errorf("Error pulling image with podman %s. %s", image, err)
		return err
	}
	return nil
}