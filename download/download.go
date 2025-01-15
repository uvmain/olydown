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

func fileAlreadyExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DownloadFile(filename string, destinationFolder string) (done bool, skipped bool, err error) {

	filePath := filepath.Join(destinationFolder, filename)

	if fileAlreadyExists(filePath) {
		log.Printf("File %s already exists", filePath)
		return false, true, err
	}

	out, err := os.Create(filePath)
	if err != nil {
		return false, false, err
	}
	defer out.Close()

	dcimFolder, err := logic.GetDcimFolder()
	if err != nil {
		return false, false, err
	}
	sourceUrl := fmt.Sprintf("http://192.168.0.10/DCIM/%s/%s", dcimFolder, filename)
	log.Printf("downloading %s", sourceUrl)
	response, err := http.Get(sourceUrl)
	if err != nil {
		return false, false, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false, false, fmt.Errorf("bad status: %s", response.Status)
	}
	log.Printf("download status code: %d", response.StatusCode)

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return false, false, err
	}

	return true, false, nil
}
