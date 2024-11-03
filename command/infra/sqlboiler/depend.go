package sqlboiler

import (
	"commandservice/infra/sqlboiler/respository"

	"go.uber.org/fx"
)

// SQLBoilerを利用したRepositoryの依存定義
var RepDepend = fx.Options(
	fx.Provide(
		// Repositoryインターフェイス実装のコンストラクタを指定
		respository.NewcategoryRepositorySQLBoiler, // カテゴリ用Repository
		respository.NewproductRepositorySQLBoiler,  // 商品用Repository
	),
)
