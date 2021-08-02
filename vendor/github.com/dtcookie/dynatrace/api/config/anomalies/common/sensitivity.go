package common

// Sensitivity Sensitivity of the threshold.
// With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.
// With `high` sensitivity, no statistical confidence is used. Each violation triggers an alert.
type Sensitivity string

// Sensitivities offers the known enum values
var Sensitivities = struct {
	High   Sensitivity
	Low    Sensitivity
	Medium Sensitivity
}{
	"HIGH",
	"LOW",
	"MEDIUM",
}
