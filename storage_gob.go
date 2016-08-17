package bayesian

import(
	"os"
	"encoding/gob"
)

type StorageGob struct {
	path string
}

func NewStorageGob(path string) *StorageGob {
	return &StorageGob{
		path: path,
	}
}

func (storage *StorageGob) Load(classifier *Classifier) error {
	reader, err := os.Open(storage.path)

	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(reader)
	raw := &Classifier{}
	err = decoder.Decoder(raw)

	if err != nil {
		return err
	}

	classifier.Learned = raw.Learned
	classifier.Seen = raw.Seen
	classifier.Data = raw.Data
}

func (storage *StorageGob) Save(classifier *Classifier) error {
	writer, err := os.OpenFile(storage.path, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	return gob.NewEncoder(writer).Encode(classifier)
}