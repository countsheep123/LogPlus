package logplus

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type CallInfo struct {
	FileName string
	FuncName string
	Line     int
}

func (ci CallInfo) String() string {
	return fmt.Sprintf("%s (%s:%d)", ci.FuncName, ci.FileName, ci.Line)
}

func getCallInfo(depth int) *CallInfo {
	pc, filePath, line, _ := runtime.Caller(depth)
	_, fileName := path.Split(filePath)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	funcName := parts[pl-1]
	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
	}

	return &CallInfo{
		FileName: fileName,
		FuncName: funcName,
		Line:     line,
	}
}
