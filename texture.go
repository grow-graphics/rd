package rd

import (
	"runtime.link/mmm"
)

type TextureFormat struct {
	ArrayLayers int
	Depth       int
	Format      DataFormat
	Height      int
	Mipmaps     int
	Samples     TextureSamples
	TextureType TextureType
	Usage       TextureUsage
	Width       int
}

type TextureType int

const (
	TextureType1D TextureType = iota
	TextureType2D
	TextureType3D
	TextureTypeCube
	TextureTypeArray1D
	TextureTypeArray2D
	TextureTypeArrayCube
)

type TextureUsage int

const (
	TextureUsageSampling TextureUsage = 1 << iota
	TextureUsageAttachment
	TextureUsageDepthStencilAttachment
	TextureUsageStorage
	TextureUsageStorageAtomic
	TextureUsageReadCPU
	TextureUsageCanUpdate
	TextureUsageCanCopyFrom
	TextureUsageCanCopyTo
	TextureUsageInputAttachment
)

type TextureView struct {
	FormatOverride DataFormat
	SwizzleAlpha   TextureSwizzle
	SwizzleBlue    TextureSwizzle
	SwizzleGreen   TextureSwizzle
	SwizzleRed     TextureSwizzle
}

type TextureSwizzle int

const (
	TextureSwizzleIdentity TextureSwizzle = iota
	TextureSwizzleZero
	TextureSwizzleOne
	TextureSwizzleRed
	TextureSwizzleGreen
	TextureSwizzleBlue
	TextureSwizzleAlpha
)

type TextureSliceType int

const (
	TextureSlice2D TextureSliceType = iota
	TextureSliceCubemap
	TextureSlice3D
)

type TextureSamples int

const (
	TextureSamples1 TextureSamples = iota
	TextureSamples2
	TextureSamples4
	TextureSamples8
	TextureSamples16
	TextureSamples32
	TextureSamples64
)

type Texture mmm.Pointer[Interface, Texture, uintptr]

func (t Texture) Free() {
	(*mmm.API(t)).FreeRID(mmm.End(t))
}
