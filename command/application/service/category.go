package service

import (
	"context"

	"commandservice/domain/models/categories"
)

// Category更新サービスインターフェイス
type CategoryService interface {
	// カテゴリを登録する
	Add(ctx context.Context, category *categories.Category)
}
