/*
@Time : 2020/8/28 15:43
@Author : ZhouHui2
@File : Wallet
@Software: GoLand
*/
package orientation

import "time"

type Wallet struct {
	id                      string
	createTime              int64
	balance                 int64
	balanceLastModifiedTime int64
}

func InitWallet() *Wallet {
	return &Wallet{
		id:         "1", // 随机生成id
		createTime: time.Now().Unix(),
	}
}
func (t *Wallet) increaseBalance(money int64) {
	t.balance += money
	t.balanceLastModifiedTime = time.Now().Unix()

}
func (t *Wallet) decreaseBalance(money int64) {
	t.balance -= money
	t.balanceLastModifiedTime = time.Now().Unix()
}
