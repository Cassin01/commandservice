package service

import (
	"context"

	"commandservice/domain/models/products"
)

// 商品更新サービスインターフェイス
type ProductService interface {
	// 商品を登録する
	Add(ctx context.Context, product *products.Product) error
	// 商品を変更する
	Update(ctx context.Context, product *products.Product) error
	// 商品を削除する
	Delete(ctx context.Context, product *products.Product) error
}
