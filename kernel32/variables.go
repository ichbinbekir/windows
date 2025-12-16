package kernel32

import "syscall"

var _kernel32 = syscall.NewLazyDLL("kernel32.dll")

var (
	_writeProcessMemory       = _kernel32.NewProc("WriteProcessMemory")
	_readProcessMemory        = _kernel32.NewProc("ReadProcessMemory")
	_openProcess              = _kernel32.NewProc("OpenProcess")
	_closeHandle              = _kernel32.NewProc("CloseHandle")
	_virtualAlloc             = _kernel32.NewProc("VirtualAlloc")
	_virtualAllocEx           = _kernel32.NewProc("VirtualAllocEx")
	_createRemoteThread       = _kernel32.NewProc("CreateRemoteThread")
	_waitForSingleObject      = _kernel32.NewProc("WaitForSingleObject")
	_createToolhelp32Snapshot = _kernel32.NewProc("CreateToolhelp32Snapshot")
	_process32First           = _kernel32.NewProc("Process32FirstW")
	_process32Next            = _kernel32.NewProc("Process32NextW")
	_module32First            = _kernel32.NewProc("Module32FirstW")
	_module32Next             = _kernel32.NewProc("Module32NextW")
	_getCurrentThreadId       = _kernel32.NewProc("GetCurrentThreadId")
)
