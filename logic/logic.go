package logic

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func fetchBody(requestUrl string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func getDcimFolder() (string, error) {
	response, err := fetchBody("http://192.168.0.10/get_imglist.cgi?DIR=/DCIM")
	if err != nil {
		return "", err
	}
	dcimFolders := strings.Split(response, "\n")
	dcimFolder := strings.Split(dcimFolders[1], ",")[1]
	return dcimFolder, nil
}

func GetImageList() (string, error) {
	dcimFolder, err := getDcimFolder()
	if err != nil {
		return "", err
	}
	imageListUrl := fmt.Sprintf("http://192.168.0.10/get_imglist.cgi?DIR=/DCIM/%s", dcimFolder)
	imageList, err := fetchBody(imageListUrl)
	if err != nil {
		return "", err
	}
	return imageList, nil
}
