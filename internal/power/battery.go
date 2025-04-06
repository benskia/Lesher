package power

import (
	"encoding/binary"
	"errors"
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

// We're working with int64 instead of int for the sake of using binary.Varint()
// and binary.PutVarint() to easily read and write ints to files.
type battery struct {
	name             string
	start            int64
	end              int64
	fullChargeSpec   int64
	fullChargeActual int64
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

func (bat *battery) writeThresholds() error {
	startPath := path.Join(batFilepath, bat.name, startFile)
	endPath := path.Join(batFilepath, bat.name, endFile)
	b := []byte{}

	if bytesWritten := binary.PutVarint(b, bat.start); bytesWritten == 0 {
		return errors.New("failed start PutVarint")
	}

	if err := os.WriteFile(startPath, b, 0644); err != nil {
		return fmt.Errorf("failed start WriteFile: %v", err)
	}

	if bytesWritten := binary.PutVarint(b, bat.end); bytesWritten == 0 {
		return errors.New("failed end PutVarint")
	}

	if err := os.WriteFile(endPath, b, 0644); err != nil {
		return fmt.Errorf("failed end WriteFile: %v", err)
	}

	return nil
}

func readInt(filepath string) (int64, error) {
	filename := path.Base(filepath)
	b, err := os.ReadFile(filepath)
	if err != nil {
		return -1, fmt.Errorf("failed %s ReadFile: %v", filename, err)
	}
	value, bytesRead := binary.Varint(b)
	if bytesRead == 0 {
		return -1, fmt.Errorf("failed %s Varint: %v", filename, err)
	}

	return value, nil
}
