/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/16
 * +----------------------------------------------------------------------
 * |Time: 1:03 上午
 * +----------------------------------------------------------------------
 */

package test

import (
	"fmt"
	"github.com/gongxulei/go_kit/balancer"
	"github.com/gongxulei/go_kit/balancer/base"
	"testing"
)

func TestNewBalancer(t *testing.T) {
	addrInfoList := make([]*base.AddrInfo, 0)
	addrInfoList = append(addrInfoList, &base.AddrInfo{
		Addr:   "127.0.0.1:9000",
		Weight: 1,
	})
	addrInfoList = append(addrInfoList, &base.AddrInfo{
		Addr:   "127.0.0.1:8000",
		Weight: 2,
	})
	addrInfoList = append(addrInfoList, &base.AddrInfo{
		Addr:   "127.0.0.1:7000",
		Weight: 1,
	})

	// 权重值不能超过10
	// 如果不使用权重可将每个权重设置为相等的数字
	random := balancer.NewBalancer(addrInfoList, "random")

	for i := 0; i < 10; i++ {
		addr := random.Balance()
		fmt.Println("random address: ", addr)
	}

	round := balancer.NewBalancer(addrInfoList, "round")

	for i := 0; i < 10; i++ {
		addr := round.Balance()
		fmt.Println("round address: ", addr)
	}
}
