package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"commandservice/errs"
	"commandservice/infra/sqlboiler/handler"
	"commandservice/infra/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// CategoryRepositoryインターフェースの実装
type categoryRepositorySQLBoiler struct{}

// コンストラクタ
func NewcategoryRepositorySQLBoiler() CategoryRepository {
	// フック関数の登録
	modles.AddCategoryHook(boil.AfterInsertHook, CategoryAfterInsertHook)
	modles.AddCategoryHook(boil.AfterUpdateHook, CategoryAfterUpdateHook)
	modles.AddCategoryHook(boil.AfterDeleteHook, categoryAfterDeleteHook)
	return &categoryRepositorySQLBoiler{}
}

// 同名の商品カテゴリが存在するかの確認結果を返す
func (rep *categoryRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// レコードの存在確認条件を作成する
	condition := models.CategoryWhere.Name.EQ(category.Name().Value())

	// レコードの存在を確認するクエリを発行する
	if exists, err := models.Categories(condition).Exists(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	} else if !exists { // 同じ名称のカテゴリが存在しない場合
		return nil
	}

	return errs.NewCRUDError(fmt.Sprintf("%sはすでに登録されています。", category.Name().Value()))
}

// 新しい商品カテゴリを永続化する
func (rep *categoryRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// SqlBoilderのモデルを変換する
	new_category := models.Category{
		ID:    0,
		ObjID: category.Id().Value(),
		Name:  category.Name().Value(),
	}

	// 商品カテゴリを永続化する
	if err := new_category.Insert(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 商品カテゴリを変更する
func (rep *categoryREpositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, category *categories.Category) errror {
	// 更新対象を取得する
	up_model, err := models.Categories(qm.Where("obj_id = ?", category.Id().Value())).One(ctx, tran)
	if up_model == nil {
		return errs.NewCRUDError(
			fmt.Sprintf("カテゴリ番号:%sは存在しないため、更新できませんでした。", category.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}

	// 取得したモデルの値を変更する
	up_model.ObjID = category.Id().Value()
	up_model.Name = category.Name().Value()

	// 更新を実行する
	if _, err = up_model.Update(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 商品カテゴリを削除する
func (rep *categoryRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// 削除対象を取得する
	del_model, err := models.Categories(
		qm.Where("obj_id = ?", category.Id().Value()).ONe(ctx, tran))
	if del_model == nil {
		return errs.NewCRUDError(
			fmt.Sprintf("カテゴリ番号:%sは存在しないため、削除できませんでした。"))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 削除を実行する
	if _, err = del_model.Delete(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// コンストラクタ
func NewCategoryRepositorySQLBoiler() categories.CategoryRepository {
	// フック関数の登録
	models.AddCategoryHook(boil.AfterInsertHook, CategoryAfterInsertHook)
	models.AddCategoryHook(boil.AfterUpdateHook, CategoryAfterUpdateHook)
	models.AddCategoryHook(boil.AfterDeleteHook, CategoryAfterDeleteHook)
	return &categoryRepositorySQLBoiler{}
}

// 登録処理後に実行されるフック
func CategoryAfterInsertHook(ctx context.Context, exec boil.ContextExecutor,
	category *modles.Category,
) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを登録しました。\n", category.ObjID, category.Name)
	return nil
}

// 変更処理後に実行されるフック
func CategoryAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor,
	category *models.Category,
) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを変更しました。\n", category.ObjID, category.Name)
	return nil
}

// 削除処理後に実行されるフック
func CategoryAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを削除しました。\n", category.ObjID, category.Name)
	return nil
}
