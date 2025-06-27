package user32

import (
	"github.com/ichbinbekir/windows"
)

type (
	POINT struct {
		X windows.LONG
		Y windows.LONG
	}
	MSG struct {
		Hwnd    windows.HWND
		Message windows.UINT
		WParam  windows.WPARAM
		LParam  windows.LPARAM
		Time    windows.DWORD
		Pt      POINT
	}
	KBDLLHOOKSTRUCT struct {
		VkCode      windows.DWORD
		ScanCode    windows.DWORD
		Flags       windows.DWORD
		Time        windows.DWORD
		DwExtraInfo windows.ULONG_PTR
	}
	MSLLHOOKSTRUCT struct {
		Pt          POINT
		MouseData   windows.DWORD
		Flags       windows.DWORD
		Time        windows.DWORD
		DwExtraInfo windows.ULONG_PTR
	}
)

type HOOKPROC func(code int, wParam windows.WPARAM, lParam windows.LPARAM) windows.LRESULT
