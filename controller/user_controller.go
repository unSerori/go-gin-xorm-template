package controller

import (
	"errors"
	"hoge-api/logging"
	"hoge-api/model"
	"hoge-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var userService = service.UserService{} // サービスの実体を作る。

// ユーザ作成
func RegisterUserHandler(c *gin.Context) {
	// 構造体にマッピング
	var bUser model.User // 構造体のインスタンス
	if err := c.ShouldBindJSON(&bUser); err != nil {
		// エラーログ
		logging.ErrorLog("Failure to bind request.", err)
		// レスポンス
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 7001,
			"srvResMsg":  "Failure to bind request.",
			"srvResData": gin.H{},
		})
		return
	}

	// 登録処理と失敗レスポンス
	token, err := userService.RegisterUser(bUser)
	if err != nil { // エラーハンドル
		// 処理で発生したエラーのうちDB関連のエラーのみ
		var mysqlErr *mysql.MySQLError // DBエラーを判定するためのDBインスタンス
		if errors.As(err, &mysqlErr) { // 第一引数のerrが第二引数の型にキャスト可能ならキャストしてtrue
			// 本処理時のエラーごとに処理(:DBエラー)
			switch err.(*mysql.MySQLError).Number {
			case 1062: // 一意性制約違反
				// エラーログ
				logging.ErrorLog("There is already a user with the same primary key. Uniqueness constraint violation.", err)
				// レスポンス
				c.JSON(http.StatusBadRequest, gin.H{
					"srvResCode": 7002,
					"srvResMsg":  "There is already a user with the same primary key. Uniqueness constraint violation.",
					"srvResData": gin.H{},
				})
			default:
				// エラーログ
				logging.ErrorLog("New user registration was not possible due to other DB problems.", err)
				// レスポンス
				c.JSON(http.StatusBadRequest, gin.H{
					"srvResCode": 7003,
					"srvResMsg":  "New user registration was not possible due to other DB problems.",
					"srvResData": gin.H{},
				})
			}
		}
		// 処理で発生したエラーのうちDB関連でないもの
		var serviceErr *service.CustomErr
		if errors.As(err, &serviceErr) {
			// 本処理時のエラーごとに処理(:DBエラー以外)
			switch serviceErr.Type {
			case service.ErrTypeHashingPassFailed: // ハッシュ化に失敗
				// エラーログ
				logging.ErrorLog("Failure to hash passwords.", err)
				// レスポンス
				c.JSON(http.StatusBadRequest, gin.H{
					"srvResCode": 7004,
					"srvResMsg":  "Failure to hash passwords.",
					"srvResData": gin.H{},
				})
			case service.ErrTypeGenTokenFailed: // トークンの作成に失敗
				// エラーログ
				logging.ErrorLog("Failed to generate token.", err)
				// レスポンス
				c.JSON(http.StatusBadRequest, gin.H{
					"srvResCode": 7005,
					"srvResMsg":  "Failed to generate token.",
					"srvResData": gin.H{},
				})
			default:
				// エラーログ
				logging.ErrorLog("New user registration was not possible due to other problems.", err)
				// レスポンス
				c.JSON(http.StatusBadRequest, gin.H{
					"srvResCode": 7006,
					"srvResMsg":  "New user registration was not possible due to other problems.",
					"srvResData": gin.H{},
				})
			}
		}
		return // エラーレスポンス後に終了
	}

	// 処理後の成功
	// 成功ログ
	logging.SuccessLog("Successful user registration.")
	// レスポンス
	c.JSON(http.StatusCreated, gin.H{
		"srvResCode": 1001,
		"srvResMsg":  "Successful user registration.",
		"srvResData": gin.H{
			"authenticationToken": token,
		},
	})
}
