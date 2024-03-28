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
	dosage := "1"
	uses := "TestUses"
	precautions := "TestPrecautions"
	preparations := "TestPreparations"
	id := db.Insert(name,dosage,uses,precautions,preparations)
	assert.Equal(t, id, int64(1))
}

func TestSelect(t *testing.T) {
	db := New()
	expected := []Herb{{Id: 1, Name: "TestName", Dosage: "1", Uses: "TestUses", Precautions: "TestPrecautions", Preparations: "TestPreparations"}}
	herbs := db.Select()
	assert.Equal(t, herbs, expected)
}

func TestSelectByUse(t *testing.T) {
	db := New()
	herbs := []Herb{}
	expected := []Herb{{Id: 1, Name: "TestName", Dosage: "1", Uses: "TestUses", Precautions: "TestPrecautions", Preparations: "TestPreparations"}}
	herbs = db.SelectByUse("TestUses")
	assert.Equal(t, herbs, expected)
}

