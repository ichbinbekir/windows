package ntdll

import "syscall"

var _ntdll = syscall.NewLazyDLL("ntdll.dll")

var (
	_ntSuspendProcess = _ntdll.NewProc("NtSuspendProcess")
	_ntResumeProcess  = _ntdll.NewProc("NtResumeProcess")
)
