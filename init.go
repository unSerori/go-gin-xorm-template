package main

import (
	"fmt"
	"hoge-api/logging"
	"hoge-api/model"
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	// .envから定数をプロセスの環境変数にロード
	err := godotenv.Load(".env") // エラーを格納
	if err != nil {              // エラーがあったら
		//logging.ErrorLog("Error loading .env file", err)
		panic("Error loading .env file.")
	}

	// ログ設定を初期化
	err = logging.SetupLogging() // セットアップ
	if err != nil {              // エラーチェック
		fmt.Printf("error opening file: %v\n", err)
	}
	log.Println("Start server!")

	// DB初期化
	model.DBConnect()
}
