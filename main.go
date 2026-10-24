package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/text/cases"
)

type Todo struct {
	//構造体タグ・・・Jsonに変換するときのフィールド名する印これがないと大文字のまま保存される。
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const fileName = "todos.json"

// ファイルから読み込む。無ければ空の一覧を返す
func load() ([]Todo, error) {
	data, err := os.ReadFile(fileName)
	if os.IsNotExist(err) {
		return []Todo{}, nil
	}
	if err != nil {
		return nil, err
	}
	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	} 
	return todos, nil
}

//ファイルへ保存する
func save(todos []Todo) error {
	data, err := json.MarshalIndent(todos, ""," ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使い方: add <内容> | list | done <ID>")
		os.Exit(1)
	}

	todos, err := load()
	if err != nil {
		fmt.Println("内容を入力してください")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		title := strings.Join(os.Args[2:], " ")
		if title = "" {
			fmt.Println("内容を入力してください")
			os.Exit(1)
		}
		maxID := 0
		for _, t := range todos {
			if t.ID > maxID {
				maxID = t.ID
			}
		}
		todos = append(todos, Todo{ID: maxID + 1, Title: Title})
		if err := save(todos); err != nil {
			fmt.Println("保存に失敗:", err)
			os.Exit(1)
		}
		fmt.Println("追加しました:", title)
	}
}