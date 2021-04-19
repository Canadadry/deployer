package fs

import ()

type local struct {
	base string
}

func (l *local) Open(name string) (File, error) {
	return nil, nil
}

func (l *local) Delete(name string) error {
	return nil
}

func (l *local) Mkdir(name string) error {
	return nil
}

func (l *local) New(name string) (File, error) {
	return nil, nil
}

func (l *local) ReadDir(name string) ([]FileInfo, error) {
	return nil, nil
}
