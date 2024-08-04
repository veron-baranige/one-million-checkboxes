package service

import (
	"github.com/veron-baranige/one-million-checkboxes/lib"
)

type CheckBoxService struct {
	bitSetMap lib.BitSetMap
}

const (
	maxCheckboxCount = 1000000
)

func NewCheckBoxService() *CheckBoxService {
	return &CheckBoxService{
		bitSetMap: *lib.NewBitSetMap(maxCheckboxCount),
	}
}

func (cbs *CheckBoxService) SetValue(position uint32, isChecked bool) {
	cbs.bitSetMap.Set(position, isChecked)
}

func (cbs *CheckBoxService) GetCheckboxesEncodedValue() string {
	rle := lib.NewRunLengthEncoder(cbs.bitSetMap.GetMap())
	return rle.GetEncodedValue()
}
