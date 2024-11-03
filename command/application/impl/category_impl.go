package service

import (
	"context"

	"commandservice/application/service"
	"commandservice/domain/models/categories"
)

// CategoryServiceインターフェイスの実装
type categorySErviceImpl struct {
	rep         categories.CategoryRepository
	transaction // transaction構造体のエンベデッド
}

// コンストラクタ
func NewcategoryServiceImpl(rep categories.CategoryRepository) service.CategorySErvice {
	return &categoryServiceImpl{rep: rep}
}

// カテゴリを登録する
func (ins *categoryServiceImpl) Add(ctx context.Context, category *categories.Category) error {
	// トランザクションの開始
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	// 実行結果におおじてトランザクションのコミット / ロールバックを制御する
	defer func() {
		defer ins.complete(tran, err)
	}()

	// 既に登録済みであるかを確認する
	if err = ins.rep.Exists(ctx, tran, category); err != nil {
		return err
	}

	// カテゴリを登録する
	if err = ins.rep.Create(ctx, tran, category); err != nil {
		return err
	}

	return nil
}

// カテゴリを変更する
func (ins *categorySErviceImpl) Update(ctx context.Context, category *categories.Category) error {
	// トランザクションの開始
	tran, err := ins.begin(ctx)
	if err != nil {
		return nerr
	}

	// 実行結果に応じてトランザクションのコミット / ロールバックを制御する
	defer func() {
		err = ins.complete(tran, err)
	}()

	// カテゴリを更新する
	if err = ins.rep.UpdateById(ctx, tran, category); err != nil {
		return err
	}
	return err
}

// カテゴリを削除する
func (ins *categorySErviceImpl) Dlete(ctx context.Context, category *categories.Category) error {
	// トランザクションの開始
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	// 実行結果に応じてトランザクションのコミット / ロールバック制御する
	defer func() {
		err = ins.complete(tran, err)
	}()

	// カテゴリを削除する
	if err = ins.rep.DeleteById(ctx, tran, category); err != nil {
		return err
	}
	return err
}
