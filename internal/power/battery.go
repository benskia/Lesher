package power

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
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
	Start            int
	End              int
	FullChargeDesign int
	FullChargeActual int
}

// A map simplifies execution of ops that target batteries by name.
type Batteries map[string]Battery

func (bat *Battery) readThresholds() error {
	startPath := path.Join(batFilepath, bat.Name, startFile)
	endPath := path.Join(batFilepath, bat.Name, endFile)

	startValue, err := readInt(startPath)
	if err != nil {
		return fmt.Errorf("failed to get start value: %v", err)
	}

	endValue, err := readInt(endPath)
	if err != nil {
		return fmt.Errorf("failed to get end value: %v", err)
	}

	bat.Start = startValue
	bat.End = endValue
	return nil
}

// Elevated permissions are required to write to power_supply files. It's a bit
// overkill to require these permissions for the entire program, so we can just
// execute sudo shell command to write.
func (bat *Battery) writeThresholds() error {
	startPath := path.Join(batFilepath, bat.Name, startFile)
	endPath := path.Join(batFilepath, bat.Name, endFile)
	var echo string
	var cmd *exec.Cmd

	echo = fmt.Sprintf("echo > %s", startPath)
	cmd = exec.Command("sudo", echo)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed start write: %v", err)
	}

	echo = fmt.Sprintf("echo > %s", endPath)
	cmd = exec.Command("sudo", echo)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed end write: %v", err)
	}

	return nil
}

func (bat *Battery) readFullCharges() error {
	fullActualPath := path.Join(batFilepath, bat.Name, fullActualFile)
	fullDesignPath := path.Join(batFilepath, bat.Name, fullDesignFile)

	fullActualValue, err := readInt(fullActualPath)
	if err != nil {
		return fmt.Errorf("failed to get actual full-charge value: %v", err)
	}

	fullDesignValue, err := readInt(fullDesignPath)
	if err != nil {
		return fmt.Errorf("failed to get design full-charge value: %v", err)
	}

	bat.FullChargeActual = fullActualValue
	bat.FullChargeDesign = fullDesignValue
	return nil
}

func readInt(filepath string) (int, error) {
	filename := path.Base(filepath)
	b, err := os.ReadFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("failed %s ReadFile: %v", filename, err)
	}

	// We should decoding directly from our []byte to int, because the files
	// might (often do) contain whitespaces that will result in wrong numbers.
	trimmedContent := strings.TrimSpace(string(b))
	value, err := strconv.Atoi(trimmedContent)
	if err != nil {
		return 0, fmt.Errorf("error converting %s: %v", trimmedContent, err)
	}

	return value, nil
}
