package abstractFactory

import "github.com/marmotedu/iam/pkg/log"

type Opener interface {
	OpenFile() error
}

type Closer interface {
	CloseFile() error
}

type OpenAndCloser interface {
	Opener
	Closer
}

type FileA struct {
}

func (f *FileA) OpenFile() error {
	log.Infof("open fileA\n")
	return nil
}

func (f *FileA) CloseFile() error {
	log.Infof("close fileA\n")
	return nil
}

type FileB struct {
}

func (f *FileB) OpenFile() error {
	log.Infof("open fileB\n")
	return nil
}

func (f *FileB) CloseFile() error {
	log.Infof("close fileB\n")
	return nil
}

func NewFileA() OpenAndCloser {
	return &FileA{}
}

func NewFileB() OpenAndCloser {
	return &FileB{}
}

func OpenFile(opt OpenAndCloser) {
	opt.OpenFile()
}

func CloseFile(opt OpenAndCloser) {
	opt.CloseFile()
}
