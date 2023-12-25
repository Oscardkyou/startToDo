package domain

import (
	"math"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))

type EntityId int

func CreateEntityId() EntityId {
	return EntityId(randomizer.Int31n(math.MaxInt32))
}
