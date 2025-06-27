package ntdll

import "github.com/ichbinbekir/windows"

func NtSuspendProcess(hProcess windows.HANDLE) windows.NTSTATUS {
	ret, _, _ := _ntSuspendProcess.Call(
		uintptr(hProcess),
	)
	return windows.NTSTATUS(ret)
}

func NtResumeProcess(hProcess windows.HANDLE) windows.NTSTATUS {
	ret, _, _ := _ntResumeProcess.Call(uintptr(hProcess))
	return windows.NTSTATUS(ret)
}
