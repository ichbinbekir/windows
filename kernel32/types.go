package kernel32

import "github.com/ichbinbekir/windows"

type PROCESSENTRY32 struct {
	DwSize              windows.DWORD
	CntUsage            windows.DWORD
	Th32ProcessID       windows.DWORD
	Th32DefaultHeapID   windows.ULONG_PTR
	Th32ModuleID        windows.DWORD
	CntThreads          windows.DWORD
	Th32ParentProcessID windows.DWORD
	PcPriClassBase      windows.LONG
	DwFlags             windows.DWORD
	SzExeFile           [windows.MAX_PATH]windows.CHAR
}

type MODULEENTRY32 struct {
	DwSize        windows.DWORD
	Th32ModuleID  windows.DWORD
	Th32ProcessID windows.DWORD
	GlblcntUsage  windows.DWORD
	ProccntUsage  windows.DWORD
	ModBaseAddr   *windows.BYTE
	ModBaseSize   windows.DWORD
	HModule       windows.HMODULE
	SzModule      [MAX_MODULE_NAME32 + 1]byte
	SzExePath     [windows.MAX_PATH]byte
}
