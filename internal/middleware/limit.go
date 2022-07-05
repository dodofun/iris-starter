package middleware

import (
	"time"

	"github.com/kataras/iris/v12/middleware/rate"
)

// 限速
var Limit = rate.Limit(1, 5, rate.PurgeEvery(time.Minute, 5*time.Minute))
