package main


type Datum struct {
	Name           string `json:""`
	OccurrenceTime int    `json:""`

	Payload string `json:""`
}

// A DataReceipt is returned when Data
type Receipt struct {
	ID string `json:""`
}
