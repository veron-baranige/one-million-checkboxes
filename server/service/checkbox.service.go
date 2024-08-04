package service

import (
	"github.com/veron-baranige/one-million-checkboxes/lib"
)

type CheckBoxService struct {
	bitSetMap lib.BitSetMap
}

func NewCheckBoxService() *CheckBoxService {
	return &CheckBoxService{
		bitSetMap: *lib.NewBitSetMap(),
	}
}

func (cbs *CheckBoxService) SetValue(position uint32, isChecked bool) {
	cbs.bitSetMap.Set(position, isChecked)
}
