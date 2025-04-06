package power

// Description:
//	Manages information about available batteries (per Linux power_supply
//	class).
//
//	Active batteries and their files might change between executions, so we
//	probably don't need persisting information about the batteries.
//
//	Lesher sets the same thresholds for all batteries, but there might be cases
//	(such as before the first profile is set) when batteries have different
//	thresholds. Still considering if it makes sense to print details for every
//	battery individually.
//
// Responsibilities:
//	- Read current charge thresholds.
//	- Write new charge thresholds.
//	- Read battery health information.
//	- Calculate battery health.

type Battery struct {
	Name             string
	Start            int
	End              int
	FullChargeSpec   int
	FullChargeActual int
}

// Reads power_supply info into Batteries that ops can work with.
func ReadThresholds() ([]Battery, error) {
	return nil, nil
}

// Writes power supply info for all power_supplies passed by ops.
func WriteThresholds([]Battery) error {
	return nil
}
