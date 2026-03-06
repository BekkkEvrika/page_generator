package inputs

type Container struct {
	Key          string      `json:"key"`
	Orientation  string      `json:"orientation"` // vertical,horizontal
	Direction    string      `json:"direction"`   // left,right,both,center
	Padding      float64     `json:"padding"`
	Margin       float64     `json:"margin"`
	Border       float64     `json:"border"`
	BorderRadius float64     `json:"borderRadius"`
	BackColor    string      `json:"backColor"`
	Title        string      `json:"title"`
	Inputs       []Input     `json:"inputs"`
	Childs       []Container `json:"childs"`
}

// GetContainerByKey ищет контейнер по ключу в текущем контейнере и его потомках (рекурсивно)
func (c *Container) GetContainerByKey(key string) *Container {
	if c.Key == key {
		return c
	}
	// Рекурсивный поиск в вложенных контейнерах
	for i := range c.Childs {
		if result := c.Childs[i].GetContainerByKey(key); result != nil {
			return result
		}
	}
	return nil
}

// GetContainerByKeyInSlice ищет контейнер по ключу в срезе контейнеров
func GetContainerByKeyInSlice(containers []Container, key string) *Container {
	for i := range containers {
		if result := containers[i].GetContainerByKey(key); result != nil {
			return result
		}
	}
	return nil
}
