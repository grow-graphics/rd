package rd

import (
	"unsafe"

	"runtime.link/mmm"
)

type Buffer mmm.Pointer[Interface, Buffer, uintptr]

func (b Buffer) Free() {
	(*mmm.API(b)).FreeRID(mmm.End(b))
}

type IndexBuffer mmm.Pointer[Interface, IndexBuffer, uintptr]

func (b IndexBuffer) Buffer() Buffer {
	return *(*Buffer)(unsafe.Pointer(&b))
}

func (b IndexBuffer) Free() {
	(*mmm.API(b)).FreeRID(mmm.End(b))
}

type StorageBuffer mmm.Pointer[Interface, StorageBuffer, uintptr]

func (b StorageBuffer) Buffer() Buffer {
	return *(*Buffer)(unsafe.Pointer(&b))
}

func (b StorageBuffer) Free() {
	(*mmm.API(b)).FreeRID(mmm.End(b))
}

type UniformBuffer mmm.Pointer[Interface, UniformBuffer, uintptr]

func (b UniformBuffer) Buffer() Buffer {
	return *(*Buffer)(unsafe.Pointer(&b))
}

func (b UniformBuffer) Free() {
	(*mmm.API(b)).FreeRID(mmm.End(b))
}

type TextureBuffer mmm.Pointer[Interface, TextureBuffer, uintptr]

func (b TextureBuffer) Buffer() Buffer {
	return *(*Buffer)(unsafe.Pointer(&b))
}

func (b TextureBuffer) Free() {
	(*mmm.API(b)).FreeRID(mmm.End(b))
}
