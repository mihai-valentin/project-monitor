package mapper

type Mapper struct {
	*Project
}

func New() *Mapper {
	return &Mapper{
		Project: newProjectMapper(),
	}
}
