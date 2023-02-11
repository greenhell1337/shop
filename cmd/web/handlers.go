package web

import (
	"fmt"
	"githab/greenhell1337/shop/pkg"
	"html/template"
	"net/http"
	"strconv"
)

func Shop(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./ui/html/shop.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := &pkg.OrderTea{
		TeaCount11: r.FormValue("TeaCount11"),
		TeaName11:  r.FormValue("TeaName11"),
		TeaCount12: r.FormValue("TeaCount12"),
		TeaName12:  r.FormValue("TeaName12"),
		TeaCount13: r.FormValue("TeaCount13"),
		TeaName13:  r.FormValue("TeaName13"),
		TeaCount14: r.FormValue("TeaCount14"),
		TeaName14:  r.FormValue("TeaName14"),
		TeaCount15: r.FormValue("TeaCount15"),
		TeaName15:  r.FormValue("TeaName15"),
		TeaCount16: r.FormValue("TeaCount16"),
		TeaName16:  r.FormValue("TeaName16"),
		TeaCount17: r.FormValue("TeaCount17"),
		TeaName17:  r.FormValue("TeaName17"),
		TeaCount18: r.FormValue("TeaCount18"),
		TeaName18:  r.FormValue("TeaName18"),
		TeaCount19: r.FormValue("TeaCount19"),
		TeaName19:  r.FormValue("TeaName19"),
		TeaCount21: r.FormValue("TeaCount21"),
		TeaName21:  r.FormValue("TeaName21"),
		TeaCount22: r.FormValue("TeaCount22"),
		TeaName22:  r.FormValue("TeaName22"),
		TeaCount23: r.FormValue("TeaCount23"),
		TeaName23:  r.FormValue("TeaName23"),
		TeaCount24: r.FormValue("TeaCount24"),
		TeaName24:  r.FormValue("TeaName24"),
		TeaCount25: r.FormValue("TeaCount25"),
		TeaName25:  r.FormValue("TeaName25"),
		TeaCount26: r.FormValue("TeaCount26"),
		TeaName26:  r.FormValue("TeaName26"),
		TeaCount27: r.FormValue("TeaCount27"),
		TeaName27:  r.FormValue("TeaName27"),
		TeaCount28: r.FormValue("TeaCount28"),
		TeaName28:  r.FormValue("TeaName28"),
		TeaCount29: r.FormValue("TeaCount29"),
		TeaName29:  r.FormValue("TeaName29"),
		TeaCount210: r.FormValue("TeaCount210"),
		TeaName210:  r.FormValue("TeaName210"),
		TeaCount211: r.FormValue("TeaCount211"),
		TeaName211:  r.FormValue("TeaName211"),
		TeaCount212: r.FormValue("TeaCount212"),
		TeaName212:  r.FormValue("TeaName212"),
		TeaCount31: r.FormValue("TeaCount31"),
		TeaName31:  r.FormValue("TeaName31"),
		TeaCount32: r.FormValue("TeaCount32"),
		TeaName32:  r.FormValue("TeaName32"),
		TeaCount33: r.FormValue("TeaCount33"),
		TeaName33:  r.FormValue("TeaName33"),
		TeaCount34: r.FormValue("TeaCount34"),
		TeaName34:  r.FormValue("TeaName34"),
		TeaCount41: r.FormValue("TeaCount41"),
		TeaName41:  r.FormValue("TeaName41"),
		TeaCount42: r.FormValue("TeaCount42"),
		TeaName42:  r.FormValue("TeaName42"),
		TeaCount43: r.FormValue("TeaCount43"),
		TeaName43:  r.FormValue("TeaName43"),
		TeaCount51: r.FormValue("TeaCount51"),
		TeaName51:  r.FormValue("TeaName51"),
		TeaCount52: r.FormValue("TeaCount52"),
		TeaName52:  r.FormValue("TeaName52"),
		TeaCount53: r.FormValue("TeaCount53"),
		TeaName53:  r.FormValue("TeaName53"),
		TeaCount54: r.FormValue("TeaCount54"),
		TeaName54:  r.FormValue("TeaName54"),
		TeaCount55: r.FormValue("TeaCount55"),
		TeaName55:  r.FormValue("TeaName55"),
		TeaCount56: r.FormValue("TeaCount56"),
		TeaName56:  r.FormValue("TeaName56"),
		TeaCount57: r.FormValue("TeaCount57"),
		TeaName57:  r.FormValue("TeaName57"),
	}

	ListTeaValueString := map[string]string{
		details.TeaName11: details.TeaCount11,
		details.TeaName12: details.TeaCount12,
		details.TeaName13: details.TeaCount13,
		details.TeaName14: details.TeaCount14,
		details.TeaName15: details.TeaCount15,
		details.TeaName16: details.TeaCount16,
		details.TeaName17: details.TeaCount17,
		details.TeaName18: details.TeaCount18,
		details.TeaName19: details.TeaCount19,
		details.TeaName21:  details.TeaCount21,
		details.TeaName22:  details.TeaCount22,
		details.TeaName23:  details.TeaCount23,
		details.TeaName24:  details.TeaCount24,
		details.TeaName25:  details.TeaCount25,
		details.TeaName26:  details.TeaCount26,
		details.TeaName27:  details.TeaCount27,
		details.TeaName28:  details.TeaCount28,
		details.TeaName29:  details.TeaCount29,
		details.TeaName210: details.TeaCount210,
		details.TeaName211: details.TeaCount211,
		details.TeaName212: details.TeaCount212,
		details.TeaName31: details.TeaCount31,
		details.TeaName32: details.TeaCount32,
		details.TeaName33: details.TeaCount33,
		details.TeaName34: details.TeaCount34,
		details.TeaName41: details.TeaCount41,
		details.TeaName42: details.TeaCount42,
		details.TeaName43: details.TeaCount43,
		details.TeaName51: details.TeaCount51,
		details.TeaName52: details.TeaCount52,
		details.TeaName53: details.TeaCount53,
		details.TeaName54: details.TeaCount54,
		details.TeaName55: details.TeaCount55,
		details.TeaName56: details.TeaCount56,
		details.TeaName57: details.TeaCount57,
	}

	ListTeaValueInt := make(map[string]int)
	for index, value := range ListTeaValueString {
		pkg.ValueInt, _ = strconv.Atoi(value)
		ListTeaValueInt[index] = pkg.ValueInt
	}

	amount := pkg.SumCost(ListTeaValueInt)
	fmt.Println(amount)

	discount := pkg.Discount(amount)

	details1 := &pkg.Order{
		Success: true,
		ListTea: ListTeaValueInt,
		Cost: amount * 1000,
		Amount: amount,
		Discount: discount*1000,
	}

	ListCity := map[string]float64{
		"Стерлитамак": 128.0,
		"Салават": 162.0,
		"Ташкиново": 206.0,
		"Нефтекамск": 216.0,
		"Бугульма": 235.0,
		"Сарапул": 282.0,
		"Альметьевск": 291.0,
		"Златоуст": 284.0,
		"Магнитогорск": 340.0,
	}
	
	_ = ListCity
	

	fmt.Println(ListTeaValueInt)

	tmpl.Execute(w, details1)
}
