package network

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	osRuntime "runtime"
	"strings"
	"time"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

var quit chan struct{}

func StartSsidTicker(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				ssid, err := GetSSID()
				if err == nil {
					// Emit the SSID event to the frontend
					wailsRuntime.EventsEmit(ctx, "ssid:update", ssid)
				} else {
					log.Printf("Error fetching SSID: %s", err)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func StopSsidTicker() {
	if quit != nil {
		close(quit)
		quit = nil
	}
}

func GetSSID() (string, error) {
	var gotSSID string
	var err error
	currentOs := osRuntime.GOOS
	log.Printf("current OS: %s", currentOs)

	switch currentOs {
	case "darwin":
		gotSSID, err = getSSIDMac()
	case "linux":
		gotSSID, err = getSSIDLinux()
	case "windows":
		gotSSID, err = getSSIDWindows()
	default:
		gotSSID, err = "", fmt.Errorf("unsupported platform: %s", osRuntime.GOOS)
	}
	log.Printf("ssid: %s", gotSSID)
	return gotSSID, err
}

func getSSIDMac() (string, error) {
	cmd := exec.Command("networksetup", "-getairportnetwork", "en0")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	output := strings.TrimSpace(string(out))
	parts := strings.Split(output, ": ")
	if len(parts) == 2 {
		return parts[1], nil
	}

	return "", fmt.Errorf("failed to parse SSID")
}

func getSSIDLinux() (string, error) {
	cmd := exec.Command("iwgetid", "-r")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

var windowsRegex = regexp.MustCompile(`\s+SSID\s+:\s+(.*)`)

func getSSIDWindows() (string, error) {
	cmd := exec.Command("cmd", "/c", "netsh", "wlan", "show", "interfaces")
	out, err := cmd.Output()
	if err != nil {
		log.Printf("error getting ssid: %s", err)
		return "", err
	}
	matches := windowsRegex.FindStringSubmatch(string(out))
	if len(matches) > 1 {
		log.Printf("ssid: %s", matches[1])
		return matches[1], nil
	}
	return "", errors.New("ssid cannot be determined")
}
