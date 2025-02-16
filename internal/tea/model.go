package internal_tea

type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitializeModel() Model {
	return Model{
		choices:  []string{"hello", "world"},
		selected: make(map[int]struct{}),
	}
}
