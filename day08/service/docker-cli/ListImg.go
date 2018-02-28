package docker_cli

import (
	"github.com/docker/docker/api/types"
)

func ListImage() ([]types.ImageSummary, error) {
	cli, err := getCli()
	if nil != err {
		return nil, err
	}
	images, err := cli.ImageList(getContext(), types.ImageListOptions{})

	if nil != err {
		return nil, err
	}

	return images, nil
}
