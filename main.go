package main

import (
	"encoding/json"

	"contentanalyzer/analyzer"
)

func main() {
	s := "кондитерская глазурь молочная (сахар, заменитель масла какао нетемперируемый нелауринового типа (рафинированные дезодорированные растительные масла в натуральном и модифицированном виде, в том числе соевое, эмульгаторы (сорбитан тристеарат, соевый лецитин), антиокислитель (токоферолы, концентрат смеси)), молоко сухое обезжиренное, какао-порошок, эмульгатор (соевый лецитин), ароматизатор «Молоко-ваниль»), мука пшеничная высшего сорта, шарики рисовые неглазированные (крупа рисовая, сахар, концентрат (экстракт) ячменно-солодовый (солод ячменный пивоваренный, ячмень, вода), соль йодированная (содержит йодат калия, агент антислеживающий (ферроцианид калия)), масло растительное, эмульгатор (соевый лецитин)), жир кондитерский (рафинированные дезодорированные растительные масла в натуральном и модифицированном виде, в том числе соевое, эмульгатор (соевый лецитин), антиокислитель (токоферолы, концентрат смеси)), сахар, измельченные хлопья кукурузные «Золотистые» (крупа кукурузная, соль йодированная (содержит йодат калия, агент антислеживающий (ферроцианид калия))), молоко сухое обезжиренное, крахмал кукурузный, какао-порошок, масло растительное, эмульгатор (соевый лецитин), соль йодированная (содержит йодат калия, агент антислеживающий (ферроцианид калия)), разрыхлитель (гидрокарбонат натрия), ароматизатор «Ванилин». Возможно наличие кусочков арахиса, ореха фундука"
	ings := analyzer.Analyze(s)
	println(json.Marshal(ings))
	return
}
