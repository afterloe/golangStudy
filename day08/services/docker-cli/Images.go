package docker_cli

import (
	"github.com/docker/docker/api/types"
	"fmt"
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

func InspectImage(imageID string) (interface{}, error) {
	cli, err := getCli()
	if nil != err {
		return nil, err
	}
	imageType, image, err := cli.ImageInspectWithRaw(getContext(), imageID)
	if nil != err {
		return nil, err
	}
	fmt.Println(string(image))
	return imageType, nil
}
