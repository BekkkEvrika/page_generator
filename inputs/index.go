package inputs

type Index struct {
	Title          string  `json:"title"`
	Function       string  `json:"function"` //count, sum, avg
	Column         string  `json:"column"`
	ColumnIdentity string  `json:"columnIdentity"`
	Where          string  `json:"where"`
	WValue         float64 `json:"wValue"`
}

func (i Index) GetName() string {
	return i.Title
}
