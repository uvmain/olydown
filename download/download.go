package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"olydown/logic"
	"os"
	"path/filepath"
)

func DownloadFile(filename string, destinationFolder string) (done bool, err error) {

	filePath := filepath.Join(destinationFolder, filename)

	if _, err := os.Stat(filePath); os.IsExist(err) {
		log.Printf("File %s already exists", filePath)
		return false, err
	}

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
