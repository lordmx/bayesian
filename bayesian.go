package bayesian

import (
	"errors"
	"math"
)

type Class string
type Word string

type Classifier struct {
	Classes []Class
	Learned int
	Seen int
	Data map[Class]*Data
}

func NewClassifier(classes ...Class) *Classifier {
	classifier := &Classifier{
		Data: make(map[Class]*Data, len(classes)),
	}

	for _, class := range classes {
		classifier.Data[class] = newData()
	}

	return classifier
}

func NewClassifierFromStorage(storage Storage, classes ...Class) (*Classifier, error) {
	classifier := NewClassifier(classes)
	err := storage.Load(classifier)

	return classifier, err
}

func (classifier *Classifier) Priors() map[Class]float64 {
	priors := make(map[Class]float64, len(classifier.Classes))
	sum := 0

	for _, class := range classifier.Classes {
		total := classifier.Data[class].Total
		sum += total
		priors[class] = float64(total)
	}

	if sum == 0 {
		return priors
	}

	for class, prior := range priors {
		priors[class] /= float64(sum)
	}

	return priors
}

func (classifier *Classifier) Learn(class Class, words ...Word) {
	for _, word := range words {
		classifier.Data[class].F[word]++
		classifier.Data[class].Total++
	}

	classifier.Learned++
}

func (classifier *Classifier) Classify(words ...Word) map[Class]float64 {
	scores := make(map[Class]float64, len(classifier.Classes))
	priors := classifier.Priors()

	classifier.Seen++

	for _, class := range classifier.Classes {
		score := math.Log(priors[class])

		for _, word := range words {
			score += math.Log(classifer.Data[class].Prob(word))
		}

		scores[class] = score
	}

	return scores
}

func (classifier *Classifier) PersistTo(storage Storage) error {
	return storage.Save(classifier)
}