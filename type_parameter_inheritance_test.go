package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

// TODO: create compatible like Employee

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetViceName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetViceName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Misno", GetName[Manager](&MyManager{Name: "Misno"}))

	assert.Equal(t, "Misno", GetName[VicePresident](&MyVicePresident{Name: "Misno"}))
}
