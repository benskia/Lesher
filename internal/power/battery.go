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
type Battery struct {
	Name             string
	Start            int64
	End              int64
	FullChargeDesign int64
	FullChargeActual int64
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

func (bat *Battery) writeThresholds() error {
	startPath := path.Join(batFilepath, bat.Name, startFile)
	endPath := path.Join(batFilepath, bat.Name, endFile)
	b := []byte{}

	if bytesWritten := binary.PutVarint(b, bat.Start); bytesWritten == 0 {
		return errors.New("failed start PutVarint")
	}

	if err := os.WriteFile(startPath, b, 0644); err != nil {
		return fmt.Errorf("failed start WriteFile: %v", err)
	}

	if bytesWritten := binary.PutVarint(b, bat.End); bytesWritten == 0 {
		return errors.New("failed end PutVarint")
	}

	if err := os.WriteFile(endPath, b, 0644); err != nil {
		return fmt.Errorf("failed end WriteFile: %v", err)
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

func readInt(filepath string) (int64, error) {
	filename := path.Base(filepath)
	b, err := os.ReadFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("failed %s ReadFile: %v", filename, err)
	}
	value, bytesRead := binary.Varint(b)
	if bytesRead == 0 {
		return 0, fmt.Errorf("failed %s Varint: %v", filename, err)
	}

	return value, nil
}
