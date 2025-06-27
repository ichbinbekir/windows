package user32

import (
	"syscall"
	"unsafe"

	"github.com/ichbinbekir/windows"
)

func SetWindowsHookExW(idHook int, lpfn HOOKPROC, hmod windows.HINSTANCE, dwThreadId windows.DWORD) windows.HHOOK {
	ret, _, _ := _setWindowsHookExW.Call(
		uintptr(idHook),
		syscall.NewCallback(lpfn),
		uintptr(hmod),
		uintptr(dwThreadId),
	)
	return windows.HHOOK(ret)
}

func GetMessageW(lpMsg *MSG, hWnd windows.HWND, wMsgFilterMin, wMsgFilterMax windows.UINT) windows.BOOL {
	ret, _, _ := _getMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
	)
	return windows.BOOL(ret)
}

func TranslateMessage(lpMsg *MSG) windows.BOOL {
	ret, _, _ := _translateMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return windows.BOOL(ret)
}

func DispatchMessageW(lpMsg *MSG) windows.LRESULT {
	ret, _, _ := _dispatchMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return windows.LRESULT(ret)
}

func UnhookWindowsHookEx(hhk windows.HHOOK) windows.BOOL {
	ret, _, _ := _unhookWindowsHookEx.Call(
		uintptr(hhk),
	)
	return windows.BOOL(ret)
}

func CallNextHookEx(hhk windows.HHOOK, nCode int, wParam windows.WPARAM, lParam windows.LPARAM) windows.LRESULT {
	ret, _, _ := _callNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return windows.LRESULT(ret)
}

func GetAsyncKeyState(vKey int) windows.SHORT {
	ret, _, _ := _getAsyncKeyState.Call(
		uintptr(vKey),
	)
	return windows.SHORT(ret)
}

func Keybd_event(bVk, bScan windows.BYTE, dwFlags windows.DWORD, dwExtraInfo windows.ULONG_PTR) {
	_keybd_event.Call(
		uintptr(bVk),
		uintptr(bScan),
		uintptr(dwFlags),
		uintptr(dwExtraInfo),
	)
}
