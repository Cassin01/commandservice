package impl_test

import (
	"os"
	"path/filepath"
	"testing"

	"commandservice/infra/sqlboiler/handler"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSrvImplPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "application/implパッケージのテスト")
}

// 全テストが実行される前に1度だけ実行される
var _ = BeforeSuite(func() {
	absPath, _ := filepath.Abs("../../../infra/sqlboiler/config/database.toml")

	// database.tomlファイルのパスを環境変数に設定する
	os.Setenv("DATABASE_TOML_PATH", absPath)
	err := handler.DBConnect() // データベースに接続する
	Expect(err).NotTo(HaveOccurred(), "データベース接続が失敗したのテストを中止します。")
})
