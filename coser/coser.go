package coser

import (
	"errors"
	"time"
)

type filepath string

type Coser struct {
	name     string
	workdays []time.Time
	boxpath  filepath
	boxs
}

type boxs struct {
	dialog           map[string]string
	dialogconfigpath string
	emoji            map[string]filepath
}

func (coser *Coser) Init() error {
	return errors.New("adf")
}
