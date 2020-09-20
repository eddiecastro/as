package services

import (
	"github.com/ajoses/salesloft-test/backend/pkg/models"
	"math"
)

const (
	percentageDiff = .10
)

type OperationsService struct {
}

func NewOperationsService() *OperationsService {
	ret := &OperationsService{}

	return ret
}

func (o *OperationsService) CharFrequency(text string) *models.Frequency {
	frequency := models.NewFrequency()

	for _, c := range text {
		frequency.AddCount(string(c))
	}

	return frequency
}

func (o *OperationsService) PossibleDuplicate(orig models.People, new models.People) bool {
	if orig.ID == new.ID {
		return false
	}
	freqOrig := o.CharFrequency(orig.EmailAddress)
	freqNew := o.CharFrequency(new.EmailAddress)

	origCount := freqOrig.TotalCount()
	for char, count := range freqNew.Freq {
		actual := freqOrig.Freq[char]
		freqOrig.Freq[char] = int32(math.Abs(float64(actual - count)))
	}

	if percentageDiff > float64(freqOrig.TotalCount())/float64(origCount) {
		return true
	}

	return false
}
