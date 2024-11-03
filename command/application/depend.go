package application

import (
	"commandservice/application/impl"
	"commandservice/infra/sqlboiler"

	"go.uber.org/fx"
)

// アプリケーション層の依存定義
var SrvDepend = fx.Options(
	sqlboiler.RepDepend, // SQLBoilerを利用したリポジトリインターフェイス実行
	fx.Provide(
		// サービスインターフェイス実装のコンストラクタ
		impl.NewcategoryServiceImpl,
		impl.NewproductServiceImpl,
	),
)
