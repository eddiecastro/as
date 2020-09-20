package services

import (
	"fmt"
	"github.com/ajoses/salesloft-test/backend/pkg/models"
	log "github.com/sirupsen/logrus"
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

func (o *OperationsService) CharFrequency(text string) models.Frequency {
	var frequency models.Frequency

	for _, c := range text {
		frequency.AddCount(string(c))
	}

	log.Info(fmt.Sprintf("%v", frequency.Freq))

	return frequency
}

func (o *OperationsService) PossibleDuplicate(orig string, new string) bool {
	var freqOrig models.Frequency
	var freqNew models.Frequency

	origCount := freqOrig.TotalCount()
	for char, count := range freqNew.Freq {
		actual := freqOrig.Freq[char]
		freqOrig.Freq[char] = int32(math.Abs(float64(actual - count)))
	}

	if percentageDiff > float64(freqOrig.TotalCount()/origCount) {
		return true
	}

	return false
}
