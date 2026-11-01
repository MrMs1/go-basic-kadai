package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

func saveDiary(entries []Entry) error {
	data, err := json.Marshal(entries)
	if err != nil {
		return err
	}
	werr := os.WriteFile("diary.json", data, 0644)
	if werr != nil {
		return werr
	}
	return nil
}

func loadDiary() ([]Entry, error) {
	data, err := os.ReadFile("diary.json")
	if os.IsNotExist(err) {
		return []Entry{}, nil
	} else if err != nil {
		return nil, err
	}
	var entries []Entry
	uerr := json.Unmarshal(data, &entries)
	if uerr != nil {
		return nil, uerr
	}

	return entries, nil
}

func main() {
	// コマンドライン引数の数をチェック
	if len(os.Args) < 2 {
		fmt.Println("使い方：go run main.go [add|list] ...")
		return
	}

	command := os.Args[1]
	// コマンドの内容に応じた処理
	switch command {
	case "add": // 日記エントリーの記録
		title := os.Args[2]
		contents := os.Args[3]
		diary, err := loadDiary()
		if err != nil {
			fmt.Printf("日記の読み込みでエラーが発生しました。")
		}
		diary = append(diary, Entry{time.Now(), title, contents})

		serr := saveDiary(diary)
		if serr != nil {
			fmt.Printf("日記の記録でエラーが発生しました。")
		}

	case "list": // 日記の一覧表示
		diary, err := loadDiary()
		if err != nil {
			fmt.Printf("日記の読み込みでエラーが発生しました。")
		}
		fmt.Printf("--- 日記データ ---\n")
		for _, note := range diary {
			fmt.Printf("日時：%s\nタイトル：%s\n本文：\n%s\n", note.Timestamp.Format("2006-01-02 15:04:05"), note.Title, note.Content)
		}
	default:
		fmt.Println("不明なコマンドです：", command)
	}

}
