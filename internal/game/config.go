package game

import (
	"github.com/franciscolkdo/family-feud/internal/family"
	"github.com/franciscolkdo/family-feud/internal/round"
)

type Config struct {
	Rounds     []round.Config `json:"rounds"`
	BlueFamily family.Config  `json:"blueFamily"`
	RedFamily  family.Config  `json:"redFamily"`
}
