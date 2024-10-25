package products_test

import (
	"commandservice/domain/models/products"
	"commandservice/errs"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Productエンティティを構成のするオブジェクト", Ordered,
	Label("ProductId構造体の生成"), func() {
		var empty_str *errs.DomainError    // 空文字列 長さ36に違反する
		var length_over *errs.DomainError  // 36文字より大きい文字列 長さ36に違反する
		var not_uuid *errs.DomainError     // UUID 以外の文字列を指定する
		var product_id *products.ProductId // UUID 文字列を指定する
		var uid string
		// 前処理
		BeforeAll(func() {
			_, empty_str = products.NewProductId("")
			_, length_over = products.NewProductId("aaaaaaaaaabbbbbbbbbbccccccccccdddddddddd")
			_, not_uuid = products.NewProductId("aaaaaaaaaabbbbbbbbbbccccccccccdddddd")
			id, _ := uuid.NewRandom()
			uid = id.String()
			product_id, _ = products.NewProductId(id.String())
		})
		// 文字数の検証
		Context("文字数の検証", Label("文字数"), func() {
			It("空文字列の場合、errs.DomainErrorが返る", func() {
				Expect(empty_str).To(
					Equal(errs.NewDomainError("商品IDの長さは36文字でなければなりません。")))
			})
			It("36文字より大きい文字列の場合、errs.DomainErrorが返る", func() {
				Expect(length_over).To(
					Equal(errs.NewDomainError("商品IDの長さは36文字でなければなりません。")))
			})
		})
		// UUID 形式の検証
		Context("UUID形式の検証", Label("UUID形式"), func() {
			It("uuid以外のの文字列の場合、errs.DomainErrorが返る", func() {
				Expect(not_uuid).To(
					Equal(errs.NewDomainError("商品IDはUUIDの形式でなければなりません。")))
			})
			It("36文字のuuid文字列の場合、ProductIdが返る", func() {
				id, _ := products.NewProductId(uid)
				Expect(product_id).To(Equal(id))
			})
		})
	})
