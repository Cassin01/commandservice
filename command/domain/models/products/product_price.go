package products

import (
	"fmt"

	"commandservice/errs"
)

// 商品単価を保持する値オブジェクト
type ProductPrice struct {
	value uint32
}

// ゲッター
func (ins *ProductPrice) Value() uint32 {
	return ins.value
}

// コンストラクタ
func NewProductPrice(value uin32) (*ProductPrice, *errs.DomainError) {
	const MIN_VALUE = 50   // 最小単価
	const MAX_VALUE = 1000 // 最大単価
	if value < MIN_VALUE || value > MAX_VALUE {
		return nil, errs.NewDomainError(
			fmt.Sprintf("単価は%d以上、%d以下です。", MIN_VALUE, MAX_VALUE))
	}
	return &ProductPrice{value: value}, nil
}
