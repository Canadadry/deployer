package fs

type memory struct {
	files map[string]File
}

func (m *memory) Open(name string) (File, error) {
	f, ok := m.files[name]
	if !ok {
		return nil, ErrFileNotFound
	}
	return f, nil
}

func (m *memory) Delete(name string) error {
	return nil
}

func (m *memory) Mkdir(name string) error {
	return nil
}

func (m *memory) New(name string) (File, error) {
	return nil, nil
}

func (m *memory) ReadDir(name string) ([]FileInfo, error) {
	return nil, nil
}
