package log

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	INFO = iota
	ERROR

	requestId = "requestId"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO", 0)
	errorLogger = log.New(os.Stdout, "ERROR", 0)
}

func doLog(cLog *log.Logger, level, callDepth int, v ...interface{}) {
	if level == ERROR {
		cLog.SetOutput(os.Stderr)
		cLog.SetFlags(log.Llongfile)
	}
	cLog.Output(callDepth, fmt.Sprintln(v...))
}

func generateTraceId(ctx context.Context) (retString string) {
	requestID, _ := ctx.Value(requestId).(string)

	if requestID != "" {
		retString = "traceId=" + requestID
	}

	return
}
func infoLog(v ...interface{}) {
	doLog(infoLogger, INFO, 4, v...)
}

func errorLog(v ...interface{}) {
	doLog(errorLogger, ERROR, 4, v...)
}

func GenericInfo(ctx context.Context, infoMessage string, data ...map[string]interface{}) {
	var fields map[string]interface{}
	if len(data) > 0 {
		fields = data[0]
	}
	trackingIDs := generateTraceId(ctx)
	fieldsBytes, _ := json.Marshal(fields)
	fieldsString := string(fieldsBytes)
	msg := fmt.Sprintf("|%s|", trackingIDs)
	if fields != nil && len(fields) > 0 {
		infoLog(msg, infoMessage, "|", fieldsString)
	} else {
		infoLog(msg, infoMessage)
	}

}

func GenericError(ctx context.Context, e error, data ...map[string]interface{}) {
	var fields map[string]interface{}
	if len(data) > 0 {
		fields = data[0]
	}
	trackingIDs := generateTraceId(ctx)
	msg := ""
	if e != nil {
		msg = fmt.Sprintf("|%s|%s", trackingIDs, e.Error())
	} else {
		msg = fmt.Sprintf("|%s", trackingIDs)
	}

	if fields != nil && len(fields) > 0 {
		fieldsBytes, _ := json.Marshal(fields)
		fieldsString := string(fieldsBytes)
		errorLog(msg, "|", fieldsString)
	} else {
		errorLog(msg)
	}
}
