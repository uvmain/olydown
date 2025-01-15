package download

import (
	"fmt"
	"io"
	"net/http"
	"olydown/logic"
	"os"
)

func DownloadFile(filename string, destinationPath string) (done bool, err error) {

	out, err := os.Create(destinationPath)
	if err != nil {
		return false, err
	}
	defer out.Close()

	dcimFolder, err := logic.GetDcimFolder()
	if err != nil {
		return false, err
	}
	sourceUrl := fmt.Sprintf("http://192.168.0.10/DCIM/%s/%s", dcimFolder, filename)

	response, err := http.Get(sourceUrl)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("bad status: %s", response.Status)
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return false, err
	}

	return true, nil
}
