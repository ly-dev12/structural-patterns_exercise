package main

import (
	"encoding/json"
	"fmt"
)

type Profesor struct {
	Name        string
	Area        string
	Edad        int
	Institucion string
}

func NewListProfesor(name, area string, edad int, institucion string) Profesor {
	return Profesor{Name: name, Area: area, Edad: edad, Institucion: institucion}
}

func (p *Profesor) Print() {
	fmt.Printf("%s %s %d %s\n", p.Name, p.Area, p.Edad, p.Institucion)
	fmt.Println("Cargando siguiente...")
}

type Profe interface {
	Add(pr Profesor)
	Execute()
}

type ProfList struct {
	Profesores []Profesor
}

func NewListProfesorEtern() *ProfList {
	return &ProfList{}
}

func (p *ProfList) Add(pr Profesor) {
	p.Profesores = append(p.Profesores, pr)
}

func (p *ProfList) byteAdapter() interface{} {
	var prof []Profesor
	for _, elem := range p.Profesores {
		prof = append(prof, elem)
	}

	return prof
}

func (p *ProfList) Execute() {
	for _, ele := range p.Profesores {
		ele.Print()
	}
}

// JSON DECORATOR
type JsonDec interface {
	byteAdapter() interface{}
}

func TransformToJson(ele JsonDec) []string {
	data := ele.byteAdapter()
	dataSon, _ := json.Marshal(data)
	fmt.Println(string(dataSon))
	return []string{"f"}
}

func main() {
	prof1 := NewListProfesor("Luis", "Matermatica", 23, "POS")
	prof2 := NewListProfesor("Andrew", "Lenguaje", 53, "KLP")
	lista := NewListProfesorEtern()
	lista.Add(prof1)
	lista.Add(prof2)
	TransformToJson(lista)
	//lista.Execute()
}
