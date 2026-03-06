package inputs

type Form struct {
	Name       string      `json:"name"`
	Containers []Container `json:"containers"`
}
