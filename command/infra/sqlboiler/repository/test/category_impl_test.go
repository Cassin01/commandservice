package repository_test

import (
	"context"
	"database/sql"
	"fmt"

	"commandservice/domain/models/categories"
	"commandservice/errs"
	"commandservice/infra/sqlboiler/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/sqlboiler/boil"
)

var _ = Describe("categoryRepositorySQLBoiler構造体", Ordered,
	Label("CategoryRepositoryインターフェイスメソッドのテスト"), func() {
		var rep categories.CategoryRepository
		var ctx context.Context
		var tran *sql.Tx

		// 前処理
		BeforeAll(func() {
			rep = repository.NewcategoryRepositorySQLBoiler() // リポジトリの生成
		})

		// テスト毎の前処理
		BeforeEach(func() {
			ctx = context.Background()
			tran, _ = boil.BeginTx(ctx, nil) // トランザクションの開始
		})

		// テスト毎の後処理
		AfterEach(func() {
			tran.Rollback()
		})

		// Exists()メソッドのテスト
		Context("同名の商品カテゴリが存在確認結果を返す", Label("Exists"), func() {
			It("存在しない商品の場合nilが返る", func() {
				name, _ := categories.NewCategoryName("文房具")
				category, _ := categories.NewCategory(name)
				result := rep.Exists(ctx, tran, category)
				Expect(result).To(BeNil())
			})
			It("存在するカテゴリ名の場合、errs.CRUDErrorが返る", func() {
				name, _ := categories.NewCategoryName("文房具")
				category, _ := categories.NewCategory(name)
				result := rep.Exists(ctx, tran, category)
				Expect(result).To(Equal(errs.NewCRUDError(
					fmt.Sprintf("%sは既に登録されています。", category.Name().Value()))))
			})
		})

		// Create()メソッドのテスト
		Context("新しい商品カテゴリを永続化する", Label("Create"), func() {
			It("カテゴリが登録成功し、nilが返る", func() {
				name, _ := categories.NewCategoryName("食品")
				category, _ := categories.NewCategory(name)
				result := rep.Create(ctx, tran, category)
				Expect(result).To(BeNil())
			})
		})
		It("obj_idが同じカテゴリを追加すると、errs.CRUDErrorが返る", func() {
			id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
			name, _ := categories.NewCategoryName("文房具")
			category := categories.BuildCategory(id, name)
			result := rep.Create(ctx, tran, category)
			Expect(result).To(Equal(errs.NewCRUDError("一意制約違反です。")))
		})
	})
