/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/16
 * +----------------------------------------------------------------------
 * |Time: 11:40 下午
 * +----------------------------------------------------------------------
 */

package limiter

type Limiter interface {
	Allow() (isAllow bool)
}