package utils

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func GetSSID() (string, error) {
	cmd := exec.Command("netsh", "wlan", "show", "interfaces")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	output := out.String()
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "SSID") && !strings.Contains(line, "BSSID") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				ssid := strings.TrimSpace(parts[1])
				return ssid, nil
			}
		}
	}

	return "", errors.New("not SSID found")
}
