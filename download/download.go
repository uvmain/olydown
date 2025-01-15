package download

import (
	"bytes"
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
		return false, true, nil
	}

	dcimFolder, err := logic.GetDcimFolder()
	if err != nil {
		return false, false, err
	}

	sourceUrl := fmt.Sprintf("http://192.168.0.10/DCIM/%s/%s", dcimFolder, filename)
	log.Printf("Downloading %s", sourceUrl)

	// Perform the HTTP GET request
	response, err := http.Get(sourceUrl)
	if err != nil {
		return false, false, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false, false, fmt.Errorf("bad status: %s", response.Status)
	}

	log.Printf("Download status code: %d", response.StatusCode)

	// Read the entire response into memory
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, response.Body)
	if err != nil {
		return false, false, err
	}

	// Create the output file once the download is complete
	out, err := os.Create(filePath)
	if err != nil {
		return false, false, err
	}
	defer out.Close()

	// Write the buffer to the file
	_, err = buffer.WriteTo(out)
	if err != nil {
		return false, false, err
	}

	log.Printf("File %s downloaded successfully", filePath)
	return true, false, nil
}
