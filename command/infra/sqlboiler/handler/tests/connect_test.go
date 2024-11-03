package handler_test

import (
	"os"
	"path/filepath"
	"testing"

	"commandservice/infra/sqlboiler/handler"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/sqlboiler/hanlderパッケージのテスト")
}

var _ = Describe("データベース接続テスト", func() {
	It("接続が成功した場合、nilが返る", Label("DB接続"), func() {
		absPath, _ := filepath.Abs("../../config/database.toml")
		// database.tomlファイルにパスを環境変数に設定する
		os.Setenv("DATABASE_TOML_PATH", absPath)
		result := handler.DBConnect()
		Expect(result).To(BeNil())
	})
})
