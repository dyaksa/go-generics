package fund_go_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type DataManager struct {
	Name string
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

func (m *DataManager) GetName() string {
	return m.Name
}

func (m *DataManager) GetManagerName() string {
	return m.Name
}

type DataVicePresident struct {
	Name string
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

func (m *DataVicePresident) GetName() string {
	return m.Name
}

func (m *DataVicePresident) GetVicePresidentName() string {
	return m.Name
}

type MyWife struct {
	Name string
}

type Wife interface {
	GetName() string
}

func (m *MyWife) GetName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	wife := GetName[Wife](&MyWife{Name: "Putri Oktaviani"})
	assert.Equal(t, "Putri Oktaviani", wife)

	manager := GetName[Manager](&DataManager{Name: "Dyaksa"})
	assert.Equal(t, "Dyaksa", manager)

	vicePresident := GetName[VicePresident](&DataVicePresident{Name: "Jauharuddin"})
	assert.Equal(t, "Jauharuddin", vicePresident)
}
