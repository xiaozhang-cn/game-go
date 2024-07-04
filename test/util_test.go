package test

import (
	"test2/utils"
	"testing"
)

var log = utils.GetLog()

func TestUtils(t *testing.T) {
	types, err := utils.ExtractStructTypes("./game")
	if err != nil {
		log.Error(err)
	}
	log.Info("scan types: ", types)
}
