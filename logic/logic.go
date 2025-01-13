package logic

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"olydown/types"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

var cachedDcimFolder string

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

func fetchBytes(requestUrl string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return resBody, nil
}

func getDcimFolder() (string, error) {
	if len(cachedDcimFolder) > 0 {
		return cachedDcimFolder, nil
	}
	response, err := fetchBody("http://192.168.0.10/get_imglist.cgi?DIR=/DCIM")
	if err != nil {
		return "", err
	}
	dcimFolders := strings.Split(response, "\n")
	dcimFolder := strings.Split(dcimFolders[1], ",")[1]
	cachedDcimFolder = dcimFolder
	return dcimFolder, nil
}

func GetImageList() ([]string, error) {
	var imageList []string
	dcimFolder, err := getDcimFolder()
	if err != nil {
		return []string{}, err
	}
	imageListUrl := fmt.Sprintf("http://192.168.0.10/get_imglist.cgi?DIR=/DCIM/%s", dcimFolder)
	response, err := fetchBody(imageListUrl)
	if err != nil {
		return []string{}, err
	}
	paths := strings.Split(response, "\n")
	for _, path := range paths {
		if strings.Contains(path, "/DCIM/") {
			imageList = append(imageList, strings.Split(path, ",")[1])
		}
	}
	return imageList, nil
}

func GetImageScreennail(filename string) (types.ImageResponse, error) {
	dcimFolder, _ := getDcimFolder()
	url := fmt.Sprintf("http://192.168.0.10/get_screennail.cgi?DIR=/DCIM/%s/%s", dcimFolder, filename)
	response, err := fetchBytes(url)

	if err != nil {
		return types.ImageResponse{}, err
	}

	base64Response := base64.StdEncoding.EncodeToString(response)
	orientation, _ := GetOrientation(filename)

	imageResponse := types.ImageResponse{
		Base64string: base64Response,
		Orientation:  orientation,
	}
	return imageResponse, nil
}

func GetOrientation(filename string) (string, error) {
	dcimFolder, _ := getDcimFolder()
	url := fmt.Sprintf("http://192.168.0.10/get_exif.cgi?DIR=/DCIM/%s/%s", dcimFolder, filename)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	exifData := buf.Bytes()

	exifInfo, err := exif.Decode(bytes.NewReader(exifData))
	if err != nil {
		return "", err
	}

	orientation, _ := exifInfo.Get(exif.Orientation)

	if orientation.String() == "8" {
		return "90", nil
	}

	return "0", nil
}
