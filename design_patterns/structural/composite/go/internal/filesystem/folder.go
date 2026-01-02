package filesystem

type Folder struct {
	name       string
	components []Component
}

func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}

func (f *Folder) Size() int {
	total := 0
	for _, c := range f.components {
		total += c.Size()
	}
	return total
}

func NewFolder(name string) *Folder {
	return &Folder{
		name: name,
	}
}
