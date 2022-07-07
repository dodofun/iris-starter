package middleware

import (
	"time"

	"github.com/kataras/iris/v12/middleware/rate"
)

// 限速：每秒钟 5 个请求，最大爆发数 50；每 1 分钟检查一次，如果是 5 分钟之前访问的旧数据，则删除。
var Limit = rate.Limit(5, 50, rate.PurgeEvery(time.Minute, 5*time.Minute))
