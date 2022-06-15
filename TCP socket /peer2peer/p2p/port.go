package p2p

import (
	"math/rand"
)

type PortDuration struct {
	Min uint
	Max uint
}

func (pDuration *PortDuration) SetDuration(min, max uint) {
	pDuration.Min = min
	pDuration.Max = max
}

func (pDuration *PortDuration) RandPort() int {
	return rand.Intn(int(pDuration.Max - pDuration.Min)) + int(pDuration.Min)
}