package handler

import (
	"errors"
	"log"
	"net"

	"commandservice/errs"

	"github.com/go-sql-driver/mysql"
)

// データベースアクセスエラーのハンドリング
func DBErrHandler(err error) error {
	var opErr *net.OpError
	var diverErr *mysql.MySQLError
	// 接続がタイムアプトかネットワーク関連の問題が原因で接続が確率できない?
	if errors.As(err, &opErr) {
		log.Println(err.Error())
		return errs.NewInternalError(opErr.Error())
	} else if errors.As(err, &driverErr) { // MySQLドライバーエラー?
		log.Printf("Code:%d Message:%s \n", driverErr.Number, driverErr.Message)
		if driverErr.Number == 1062 { // 一意制約違反?
			return errs.NewCRUDError("一意制約違反です。")
		} else {
			return errs.NewInternalError(driverErr.Message)
		}
	}

	// その他エラー
	log.Println(err.Error())
	return errs.NewInternalError(err.Error())
}
