package ccm

import (
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)

func SentrySendError(errorInfo error) {
	// init by ENVIRONMNET
	err := sentry.Init(sentry.ClientOptions{
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	// 发送错误 sentry.CaptureException(exception error)
	sentry.CaptureException(errorInfo)
}
