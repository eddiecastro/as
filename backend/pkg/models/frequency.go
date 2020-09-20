package models

type Frequency struct {
	Freq map[string]int32
}

func (f *Frequency) AddCount(char string) {
	count := f.Freq[char]

	count = count + 1
	f.Freq[char] = count
}

func (f *Frequency) GetCount(char string) int32 {
	return f.Freq[char]
}

func (f *Frequency) TotalCount() (total int32) {

	for _, count := range f.Freq {
		total = total + count
	}

	return
}
