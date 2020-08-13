package ccm

import (
	"github.com/getsentry/sentry-go"

	"time"
)

func SentrySendError(errorInfo error) {
	// init by ENVIRONMNET
	_ = sentry.Init(sentry.ClientOptions{
	})
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	// 发送错误 sentry.CaptureException(exception error)
	sentry.CaptureException(errorInfo)
}
