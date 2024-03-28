package model

type Herb struct {
	Id int64
	Name string `form:"name"`
	Dosage string `form:"dosage"`
	Uses string `form:"uses"`
	Precautions string `form:"precautions"`
	Preparations string `form:"preparations"`
}