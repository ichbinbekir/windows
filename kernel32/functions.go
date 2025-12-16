package kernel32

import (
	"unsafe"

	"github.com/ichbinbekir/windows"
)

func WriteProcessMemory(hProcess windows.HANDLE, lpBaseAddress windows.LPVOID, lpBuffer windows.LPCVOID, nSize windows.SIZE_T, lpNumberOfBytesWritten *windows.SIZE_T) windows.BOOL {
	ret, _, _ := _writeProcessMemory.Call(
		uintptr(hProcess),
		uintptr(lpBaseAddress),
		uintptr(lpBuffer),
		uintptr(nSize),
		uintptr(unsafe.Pointer(lpNumberOfBytesWritten)),
	)
	return windows.BOOL(ret)
}

func ReadProcessMemory(hProcess windows.HANDLE, lpBaseAddress windows.LPCVOID, lpBuffer windows.LPVOID, nSize windows.SIZE_T, lpNumberOfBytesRead *windows.SIZE_T) windows.BOOL {
	ret, _, _ := _readProcessMemory.Call(
		uintptr(hProcess),
		uintptr(lpBaseAddress),
		uintptr(lpBuffer),
		uintptr(nSize),
		uintptr(unsafe.Pointer(lpNumberOfBytesRead)),
	)
	return windows.BOOL(ret)
}

func OpenProcess(dwDesiredAccess windows.DWORD, bInheritHandle windows.BOOL, dwProcessId windows.DWORD) windows.HANDLE {
	ret, _, _ := _openProcess.Call(
		uintptr(dwDesiredAccess),
		uintptr(bInheritHandle),
		uintptr(dwProcessId),
	)
	return windows.HANDLE(ret)
}

func CloseHandle(hObject windows.HANDLE) windows.BOOL {
	ret, _, _ := _closeHandle.Call(
		uintptr(hObject),
	)
	return windows.BOOL(ret)
}

func VirtualAlloc(lpAddress windows.LPVOID, dwSize windows.SIZE_T, flAllocationType, flProtect windows.DWORD) windows.LPVOID {
	ret, _, _ := _virtualAlloc.Call(
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flAllocationType),
		uintptr(flProtect),
	)
	return windows.LPVOID(ret)
}

func VirtualAllocEx(hProcess windows.HANDLE, lpAddress windows.LPVOID, dwSize windows.SIZE_T, flAllocationType, flProtect windows.DWORD) windows.LPVOID {
	ret, _, _ := _virtualAllocEx.Call(
		uintptr(hProcess),
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flAllocationType),
		uintptr(flProtect),
	)
	return windows.LPVOID(ret)
}

func CreateRemoteThread(hProcess windows.HANDLE, lpThreadAttributes windows.LPVOID, dwStackSize windows.SIZE_T, lpStartAddress, lpParameter windows.LPVOID, dwCreationFlags windows.DWORD, lpThreadId *windows.DWORD) windows.HANDLE {
	ret, _, _ := _createRemoteThread.Call(
		uintptr(hProcess),
		uintptr(lpThreadAttributes),
		uintptr(dwStackSize),
		uintptr(lpStartAddress),
		uintptr(lpParameter),
		uintptr(dwCreationFlags),
		uintptr(unsafe.Pointer(lpThreadId)),
	)
	return windows.HANDLE(ret)
}

func WaitForSingleObject(hHandle windows.HANDLE, dwMilliseconds windows.DWORD) windows.DWORD {
	ret, _, _ := _waitForSingleObject.Call(
		uintptr(hHandle),
		uintptr(dwMilliseconds),
	)
	return windows.DWORD(ret)
}

func CreateToolhelp32Snapshot(dwFlags, th32ProcessID windows.DWORD) windows.HANDLE {
	ret, _, _ := _createToolhelp32Snapshot.Call(
		uintptr(dwFlags),
		uintptr(th32ProcessID),
	)
	return windows.HANDLE(ret)
}

func Process32First(hSnapshot windows.HANDLE, lppe *PROCESSENTRY32) windows.BOOL {
	ret, _, _ := _process32First.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(lppe)),
	)
	return windows.BOOL(ret)
}

func Process32Next(hSnapshot windows.HANDLE, lppe *PROCESSENTRY32) windows.BOOL {
	ret, _, _ := _process32Next.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(lppe)),
	)
	return windows.BOOL(ret)
}

func Module32First(hSnapshot windows.HANDLE, lpme *MODULEENTRY32) windows.BOOL {
	ret, _, _ := _module32First.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(lpme)),
	)
	return windows.BOOL(ret)
}

func Module32Next(hSnapshot windows.HANDLE, lpme *MODULEENTRY32) windows.BOOL {
	ret, _, _ := _module32Next.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(lpme)),
	)
	return windows.BOOL(ret)
}

func GetCurrentThreadId() windows.DWORD {
	ret, _, _ := _getCurrentThreadId.Call()
	return windows.DWORD(ret)
}
