package test

import (
	"testing"
  api"cabservice/api"
  models"cabservice/model"
)
// Testing the cab package
func TestInitiaizeCabs(t *testing.T) {

	var cabs []models.CabInfo
	cabs = api.InitializeCabs()
	if len(cabs) <= 0 {
		t.Errorf("Cab Initializion failed %+v", len(cabs))
	}
}
