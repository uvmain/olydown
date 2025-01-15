package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"olydown/logic"
	"os"
	"path"
)

func DownloadFile(filename string, destinationFolder string) (done bool, err error) {

	filePath := path.Join(destinationFolder, filename)

	out, err := os.Create(filePath)
	if err != nil {
		return false, err
	}
	defer out.Close()

	dcimFolder, err := logic.GetDcimFolder()
	if err != nil {
		return false, err
	}
	sourceUrl := fmt.Sprintf("http://192.168.0.10/DCIM/%s/%s", dcimFolder, filename)
	log.Printf("downloading %s", sourceUrl)
	response, err := http.Get(sourceUrl)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("bad status: %s", response.Status)
	}
	log.Printf("download status code: %d", response.StatusCode)

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return false, err
	}

	return true, nil
}
