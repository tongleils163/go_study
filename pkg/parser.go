package pkg

import (
	"fmt"
	"github.com/antonmedv/expr"
)

type QA struct {
	QuestionId int
	AnswerId   int

	Previous func(int) int
}

type TT struct {
	QAS map[string]int
}

func arra() {

}
func Parse() {
	qa := QA{QuestionId: 100, AnswerId: 2}
	qa.Previous = func(s int) int {
		if s == 2 {
			return 5
		}
		return 6
	}

	// qas := make(map[string]int)
	// qas["100"] = 10
	// qas["101"] = 11
	// qas["102"] = 12
	// qas["103"] = 13
	// qas["104"] = 14
	// tt := &TT{QAS: qas}
	// fmt.Println(qas)
	out, err := expr.Eval("AnswerId==2", qa)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	p, err := expr.Compile("Previous(AnswerId)", expr.Env(QA{}))
	out, err = expr.Run(p, qa)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	// p1, err := expr.Compile(`["100"] ==10`, expr.Env(&TT{}))
	// out, err = expr.Run(p1, tt)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(out)

	type Request struct {
		Meta map[string]string
	}
	request := &Request{
		Meta: map[string]string{"q1": "a"},
	}
	// code := `Meta["q1"] =="a"`
	code := `filter(Meta["q1"] =="a")`

	program, err := expr.Compile(code, expr.Env(&Request{}))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	output, err := expr.Run(program, request)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("%v", output)

}

type QuestionAnswers struct {
	Answer   int
	Question int
}
type LearningPathPromoter struct {
	Type    int64
	GoodsId int64
}
type A struct {
	B int
	C int
}

func test11(s string) {

	ch := make(chan int, 2)

	<-ch
	ch <- 4

	aaa := A{B: 55}
	test := map[string]interface{}{
		"Data": []LearningPathPromoter{
			LearningPathPromoter{Type: 1,
				GoodsId: 5},
			LearningPathPromoter{Type: 1,
				GoodsId: 6},
			LearningPathPromoter{Type: 2,
				GoodsId: 5},
		},
		"A": A{B: 55},
		"B": A{B: 11},
	}

	// output, err := expr.Eval(`len(filter(Data, {.Type > 1}))`, test)
	output, err := expr.Eval(`3>1 ?A: B`, test)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)

	env := map[string]interface{}{
		"QuestionAnswers": map[int]int{
			100: 1,
			101: 5,
		},

		// "foo": 1,
		// "bar": []string{"zero", "hello world"},
		"promote100": func(answer map[int]int, answer1 string) A {
			fmt.Println(answer)
			fmt.Println(answer1)
			if answer[100] == 2 {
				return A{B: 1, C: aaa.C}
			}
			return A{B: 1, C: aaa.B}
		},
		"promote101": func(answer map[int]int) int { return 2 },
	}
	// 	output, err := expr.Eval("swipe(bar[foo])", env)
	output, err = expr.Eval(`promote100(QuestionAnswers,"a" )`, env)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v", output.(A))

}
