package bayesian

type Storage interface {
	Load(classifier *Classifier) error
	Save(classifier *Classifier) error
}
