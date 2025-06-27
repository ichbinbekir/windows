package windows

import "unsafe"

type (
	HWND      uintptr
	UINT      uint32
	WPARAM    uintptr
	LPARAM    uintptr
	DWORD     uint32
	LONG      int32
	HHOOK     uintptr
	HINSTANCE uintptr
	LRESULT   uintptr
	BOOL      int32
	ULONG_PTR uintptr
	SHORT     int16
	BYTE      byte
	HANDLE    uintptr
	SIZE_T    ULONG_PTR
	LPVOID    unsafe.Pointer
	LPCVOID   unsafe.Pointer
	NTSTATUS  uint32
	HMODULE   uintptr
	CHAR      byte
)
