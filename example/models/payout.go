package models

type SymbolPayouts map[uint8][]uint8

type SP map[uint8][]int

func (p SymbolPayouts) GetP(s uint8, count int) uint8 {
	if sp, ok := p[s]; ok {
		if len(sp) > count {
			return sp[count]
		}
	}

	return 0
}

type Getter interface {
	Get(s uint8, count int) uint8
}

type MapPayout struct {
	payout SymbolPayouts
}

func NewMapPayout(p SymbolPayouts) MapPayout {
	return MapPayout{payout: p}
}

func (mp MapPayout) Get(s uint8, count int) uint8 {
	return mp.payout.GetP(s, count)
}

type SwitchPayout struct {
	payout Getter
}

func NewSwitchPayout(g Getter) SwitchPayout {
	return SwitchPayout{payout: g}
}

func (sp SwitchPayout) Get(s uint8, count int) uint8 {
	return sp.payout.Get(s, count)
}

var (
	Lp = SymbolPayouts{
		0: {0, 0, 10, 20},
		1: {0, 0, 20, 30, 40},
		2: {0, 0, 30, 40, 50},
		3: {0, 0, 30, 40, 50, 60, 70, 80, 90},
		4: {0, 50},
		5: {0, 0, 30, 40, 50},
		6: {50},
		7: {0, 0, 30, 40, 50},
	}
)
