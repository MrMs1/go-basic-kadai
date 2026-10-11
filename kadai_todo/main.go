package main

import "fmt"

type ToDo struct {
	ID        int
	Title     string
	Completed bool
}

func (t *ToDo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", t.ID)
	t.Completed = true
}

func printTodos(todos []ToDo) {
	for _, t := range todos {
		var completed string
		if t.Completed == true {
			completed = "完　了"
		} else {
			completed = "未完了"
		}
		fmt.Printf("[%s]（ID：%d）%s\n", completed, t.ID, t.Title)
	}
}

func main() {
	todos := []ToDo{
		{1, "学習計画", false},
		{2, "環境構築", false},
		{3, "基礎文法", false},
	}

	printTodos(todos)
	todos[0].Complete()
	todos[1].Complete()
	printTodos(todos)
}
