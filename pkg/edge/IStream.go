//go:build windows

package edge

import (
	"io"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _IStreamVtbl struct {
	_IUnknownVtbl
	Read         ComProc
	Write        ComProc
	Seek         ComProc
	SetSize      ComProc
	CopyTo       ComProc
	Commit       ComProc
	Revert       ComProc
	LockRegion   ComProc
	UnlockRegion ComProc
	Stat         ComProc
	Clone        ComProc
}

type IStream struct {
	vtbl *_IStreamVtbl
}

func NewIStream() *IStream {
	return &IStream{
		vtbl: &_IStreamVtbl{},
	}
}

func (i *IStream) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *IStream) Release() error {
	return i.vtbl.CallRelease(unsafe.Pointer(i))
}

func (i *IStream) Read(p []byte) (int, error) {
	bufLen := len(p)
	if bufLen == 0 {
		return 0, nil
	}

	var n int
	res, _, err := i.vtbl.Read.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&p[0])),
		uintptr(bufLen),
		uintptr(unsafe.Pointer(&n)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}

	switch windows.Handle(res) {
	case windows.S_OK:
		// The buffer has been completely filled
		return n, nil
	case windows.S_FALSE:
		// The buffer has been filled with less than len data and the stream is EOF
		return n, io.EOF
	default:
		return 0, syscall.Errno(res)
	}
}

func (i *IStream) Write(p []byte) (int, error) {
	bufLen := len(p)
	if bufLen == 0 {
		return 0, nil
	}

	var n int
	res, _, err := i.vtbl.Write.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&p[0])),
		uintptr(bufLen),
		uintptr(unsafe.Pointer(&n)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}

	switch windows.Handle(res) {
	case windows.S_OK:
		// The buffer has been completely written
		return n, nil
	case windows.S_FALSE:
		// The buffer has been written with less than len data and the stream is EOF
		return n, io.EOF
	default:
		return 0, syscall.Errno(res)
	}
}

func (i *IStream) Seek(offset int64, whence int) (int64, error) {
	var n int64
	res, _, err := i.vtbl.Seek.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(offset),
		uintptr(whence),
		uintptr(unsafe.Pointer(&n)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}

	switch windows.Handle(res) {
	case windows.S_OK:
		// The buffer has been completely written
		return n, nil
	case windows.S_FALSE:
		// The buffer has been written with less than len data and the stream is EOF
		return n, io.EOF
	default:
		return 0, syscall.Errno(res)
	}
}

func (i *IStream) SetSize(size int64) error {
	res, _, err := i.vtbl.SetSize.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(size),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}

	if windows.Handle(res) != windows.S_OK {
		return syscall.Errno(res)
	}

	return nil
}

func (i *IStream) CopyTo(dest *IStream, size int64) error {
	res, _, err := i.vtbl.CopyTo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(dest)),
		uintptr(size),
		uintptr(0),
		uintptr(0),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}

	if windows.Handle(res) != windows.S_OK {
		return syscall.Errno(res)
	}

	return nil
}

func (i *IStream) Commit(flags int) error {
	res, _, err := i.vtbl.Commit.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(flags),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}

	if windows.Handle(res) != windows.S_OK {
		return syscall.Errno(res)
	}

	return nil
}

func (i *IStream) Revert() error {
	res, _, err := i.vtbl.Revert.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}

	if windows.Handle(res) != windows.S_OK {
		return syscall.Errno(res)
	}

	return nil
}

func (i *IStream) LockRegion(offset int64, size int64, lockType int) error {
	res, _, err := i.vtbl.LockRegion.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(offset),
		uintptr(size),
		uintptr(lockType),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}

	if windows.Handle(res) != windows.S_OK {
		return syscall.Errno(res)
	}

	return nil
}

func (i *IStream) UnlockRegion(offset int64, size int64, lockType int) error {
	res, _, err := i.vtbl.UnlockRegion.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(offset),
		uintptr(size),
		uintptr(lockType),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}

	if windows.Handle(res) != windows.S_OK {
		return syscall.Errno(res)
	}

	return nil
}

const STAT_DEFAULT uint32 = 0
const STAT_NONAME uint32 = 1
const STAT_NOOPEN uint32 = 2

const STGC_DEFAULT int = 0
const STGC_OVERWRITE int = 1
const STGC_ONLYIFCURRENT int = 2
const STGC_DANGEROUSLYCOMMITMERELYTODISKCACHE int = 4
const STGC_CONSOLIDATE int = 8

func (i *IStream) Stat(stat *STATSTG, statFlag uint32) error {
	res, _, err := i.vtbl.Stat.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(stat)),
		uintptr(statFlag),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}

	if windows.Handle(res) != windows.S_OK {
		return syscall.Errno(res)
	}

	return nil
}

type STATSTG struct {
	Name           []uint16
	Type           uint32
	Size           uint64
	Mtime          windows.Filetime
	Ctime          windows.Filetime
	Atime          windows.Filetime
	Mode           uint32
	LocksSupported uint32
	Clsid          syscall.GUID
	StateBits      uint32
	Reserved       uint32
}
