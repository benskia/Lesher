package power

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/benskia/Thresher/internal/config"
)

const batFilepath string = "/sys/class/power_supply/"

const (
	startFile      string = "charge_control_start_threshold"
	endFile               = "charge_control_end_threshold"
	fullDesignFile        = "energy_full_design"
	fullActualFile        = "energy_full"
	statusFile            = "status"
)

type Battery struct {
	Name             string
	Status           string
	Start            int
	End              int
	FullChargeDesign int
	FullChargeActual int
}

// A map simplifies execution of ops that target batteries by name.
type Batteries map[string]Battery

func (bat *Battery) readThresholdFiles() error {
	startPath := path.Join(batFilepath, bat.Name, startFile)
	endPath := path.Join(batFilepath, bat.Name, endFile)
	statusPath := path.Join(batFilepath, bat.Name, statusFile)

	startValue, err := fileToInt(startPath)
	if err != nil {
		return fmt.Errorf("failed to get start value: %w", err)
	}

	endValue, err := fileToInt(endPath)
	if err != nil {
		return fmt.Errorf("failed to get end value: %w", err)
	}

	status, err := fileToStr(statusPath)
	if err != nil {
		return err
	}

	bat.Start = startValue
	bat.End = endValue
	bat.Status = status
	return nil
}

// Elevated permissions are required to write to power_supply files. It's a bit
// overkill to require these permissions for the entire program, so we can just
// execute sudo shell commands instead of using WriteFile.
func (bat *Battery) writeThresholdFiles(profile config.Profile) error {
	startPath := path.Join(batFilepath, bat.Name, startFile)
	endPath := path.Join(batFilepath, bat.Name, endFile)

	type writeInfo struct {
		path  string
		value []byte
	}

	// The order we write in matters. If the new start is higher than the
	// current end, the command will fail. Same if we try to set an end that
	// is lower than the current start.
	toWrite := []writeInfo{}
	startData := writeInfo{path: startPath, value: []byte(strconv.Itoa(profile.Start))}
	endData := writeInfo{path: endPath, value: []byte(strconv.Itoa(profile.End))}

	if profile.Start >= bat.End {
		toWrite = append(toWrite, endData, startData)
	} else {
		toWrite = append(toWrite, startData, endData)
	}

	for _, data := range toWrite {
		cmd := exec.Command("sudo", "dd", "of="+data.path)
		cmd.Stdin = bytes.NewReader(data.value)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error executing command:\n\t%+v\n\t%w", cmd, err)
		}
	}

	return nil
}

func (bat *Battery) readFullChargeFiles() error {
	fullActualPath := path.Join(batFilepath, bat.Name, fullActualFile)
	fullDesignPath := path.Join(batFilepath, bat.Name, fullDesignFile)

	fullActualValue, err := fileToInt(fullActualPath)
	if err != nil {
		return fmt.Errorf("failed to get actual full-charge value: %w", err)
	}

	fullDesignValue, err := fileToInt(fullDesignPath)
	if err != nil {
		return fmt.Errorf("failed to get design full-charge value: %w", err)
	}

	bat.FullChargeActual = fullActualValue
	bat.FullChargeDesign = fullDesignValue
	return nil
}

func fileToInt(filepath string) (int, error) {
	filename := path.Base(filepath)
	b, err := os.ReadFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("failed %s ReadFile: %w", filename, err)
	}

	trimmedContent := strings.TrimSpace(string(b))
	value, err := strconv.Atoi(trimmedContent)
	if err != nil {
		return 0, fmt.Errorf("error converting %s: %w", trimmedContent, err)
	}

	return value, nil
}

func fileToStr(filepath string) (string, error) {
	filename := path.Base(filepath)
	b, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("error reading %s: %w", filename, err)
	}

	return strings.TrimSpace(string(b)), nil
}
