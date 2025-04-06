package power

import (
	"encoding/binary"
	"fmt"
	"os"
	"path"
)

const batFilepath string = "/sys/class/power_supply/"

const (
	startFile      string = "charge_control_start_threshold"
	endFile               = "charge_control_end_threshold"
	fullDesignFile        = "energy_full_design"
	fullActualFile        = "energy_full"
	statusFile            = "status"
)

type battery struct {
	name             string
	start            int
	end              int
	fullChargeSpec   int
	fullChargeActual int
}

func (bat *battery) readThresholds() error {
	startPath := path.Join(batFilepath, bat.name, startFile)
	endPath := path.Join(batFilepath, bat.name, endFile)

	startValue, err := readInt(startPath)
	if err != nil {
		return fmt.Errorf("failed to get start value: %v", err)
	}

	endValue, err := readInt(endPath)
	if err != nil {
		return fmt.Errorf("failed to get end value: %v", err)
	}

	bat.start = startValue
	bat.end = endValue
	return nil
}

func readInt(filepath string) (int, error) {
	filename := path.Base(filepath)
	b, err := os.ReadFile(filepath)
	if err != nil {
		return -1, fmt.Errorf("failed %s ReadFile: %v", filename, err)
	}
	value, bytesRead := binary.Varint(b)
	if bytesRead == 0 {
		return -1, fmt.Errorf("failed %s Varint: %v", filename, err)
	}

	return int(value), nil
}
