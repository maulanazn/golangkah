package generics_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SoftwareEngineer interface {
	GetName() string
}

func GetName[T SoftwareEngineer](value T) string {
	return value.GetName()
}

type SoftwareEngineerII interface {
	GetName() string
	GetBackendEngineerName() string
}

type MyBackendEngineer struct {
	Name string
}

func (Sfw *MyBackendEngineer) GetName() string {
	return Sfw.Name
}

func (Sfw *MyBackendEngineer) GetBackendEngineerName() string {
	return Sfw.Name
}

type SoftwareEngineerIII interface {
	GetName() string
	GetFrontendEngineerName() string
}

type MyFrontendEngineer struct {
	Name string
}

func (Sfw *MyFrontendEngineer) GetName() string {
	return Sfw.Name
}

func (Sfw *MyFrontendEngineer) GetFrontendEngineerName() string {
	return Sfw.Name
}
func TestInheritance(t *testing.T) {
	t.Run("softwareengineer", func(t *testing.T) {
		assert.Exactly(t, "Maulanazn", GetName(&MyBackendEngineer{Name: "Maulanazn"}))
		assert.Exactly(t, "Faisal", GetName(&MyFrontendEngineer{Name: "Faisal"}))
	})
}
