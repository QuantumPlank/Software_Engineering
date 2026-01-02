package filesystem

type File struct {
	name string
	size int
}

func (f *File) Size() int {
	return f.size
}

func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}
