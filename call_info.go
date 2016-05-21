package logplus

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type CallInfo struct {
	PackageName string
	FileName    string
	FuncName    string
	Line        int
}

func (ci CallInfo) String() string {
	return fmt.Sprintf("%s#%s (%s:%d)", ci.PackageName, ci.FuncName, ci.FileName, ci.Line)
}

func getCallInfo(depth int) *CallInfo {
	pc, filePath, line, _ := runtime.Caller(depth)
	_, fileName := path.Split(filePath)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]
	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &CallInfo{
		PackageName: packageName,
		FileName:    fileName,
		FuncName:    funcName,
		Line:        line,
	}
}
