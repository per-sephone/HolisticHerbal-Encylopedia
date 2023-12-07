package model

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	got := New()
	assert.Equal(t, reflect.TypeOf(got), reflect.TypeOf(&Model{}))

}

func TestInsert(t *testing.T) {
	db := New()
	name := "TestName"
	dosage := 1
	uses := "TestUses"
	precautions := "TestPrecautions"
	preparations := "TestPreparations"
	id := db.Insert(name,dosage,uses,precautions,preparations)
	assert.Equal(t, id, int64(1))
}


func TestSelectByName(t *testing.T) {
	db := New()
	herb := Herb{}
	herb = db.SelectByName("TestName")
	assert.Equal(t, herb.Name, "TestName")
	assert.Equal(t, herb.Dosage, 1)
	assert.Equal(t, herb.Uses, "TestUses")
	assert.Equal(t, herb.Precautions, "TestPrecautions")
	assert.Equal(t, herb.Preparations, "TestPreparations")
}


func TestSelect(t *testing.T) {
	db := New()
	expected := []Herb{{Id: 1, Name: "TestName", Dosage: 1, Uses: "TestUses", Precautions: "TestPrecautions", Preparations: "TestPreparations"}}
	herbs := db.Select()
	assert.Equal(t, herbs, expected)
}
