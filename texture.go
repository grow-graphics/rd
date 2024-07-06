package rd

import (
	"io"

	"grow.graphics/uc"
)

// Texture reference.
type Texture interface {
	Resource
	Nameable
	Variable

	/*
		Clear clears the specified texture by replacing all of its pixels with the specified color.
		base_mipmap and mipmap_count determine which mipmaps of the texture are affected by this clear operation,
		while base_layer and layer_count determine which layers of a 3D texture (or texture array) are affected
		by this clear operation. For 2D textures (which only have one layer by design), base_layer must be 0 and
		layer_count must be 1.

		Note: texture can't be cleared while a draw list that uses it as part of a framebuffer is being created.
		Ensure the draw list is finalized (and that the color/depth texture using it is not set to [FrameSuspend])
		to clear this texture.
	*/
	Clear(color uc.Color, base_mipmap, mipmap_count, base_layer, layer_count int, barrier Barrier) error

	// Format returns the data format used to create this texture.
	Format() TextureFormat

	// Handle returns the internal graphics handle for this texture object. For use when communicating with third-party APIs.
	Handle() uintptr

	// IsShared returns true if the texture is shared, false otherwise.
	IsShared() bool

	// IsValid returns true if the texture is valid, false otherwise.
	IsValid() bool

	// Layer returns the texture data for the specified layer.
	Layer(layer int) TextureData
}

// TextureData can be read from and/or written depending on usage flags.
type TextureData interface {
	io.Reader
	io.ReaderFrom
	io.Closer
}

// TextureBuffer stores texture data.
type TextureBuffer interface {
	Variable

	Resource
	Nameable
}

// TextureFormat for a texture.
type TextureFormat struct {
	ArrayLayers int            // The number of layers in the texture. Only relevant for 2D texture arrays.
	Depth       int            // The texture's depth (in pixels). This is always 1 for 2D textures.
	Format      DataFormat     // The texture's pixel data format.
	Height      int            // The texture's height (in pixels).
	Mipmaps     int            // The number of mipmaps available in the texture.
	Samples     TextureSamples // The number of samples used when sampling the texture.
	TextureType TextureType    // The texture type.
	Usage       TextureUsage   // The texture's usage bits, which determine what can be done using the texture.
	Width       int            // The texture's width (in pixels).

	ShareableFormats map[DataFormat]struct{} // The shareable formats for this texture.
}

// TextureType for a texture.
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

// TextureUsage for a texture.
type TextureUsage int

const (
	TextureSampling               TextureUsage = 1 << iota // Texture can be sampled.
	TextureAttachment                                      // Texture can be used as a color attachment in a framebuffer.
	TextureDepthStencilAttachment                          // Texture can be used as a depth/stencil attachment in a framebuffer.
	TextureStorage                                         // Texture can be used as a storage image.
	TextureStorageAtomic                                   // Texture can be used as a storage image with support for atomic operations.
	TextureReadCPU                                         // Texture can be read back on the CPU using texture_get_data faster than without this bit, since it is always kept in the system memory.
	TextureCanUpdate                                       // Texture can be updated using texture_update.
	TextureCanCopyFrom                                     // Texture can be a source for texture_copy.
	TextureCanCopyInto                                     // Texture can be a destination for texture_copy.
	TextureInputAttachment                                 // Texture can be used as a input attachment in a framebuffer.
)

// TextureView for a texture.
type TextureView struct {
	FormatOverride DataFormat // Optional override for the data format to return sampled values in.
	SwizzleAlpha   Swizzle    // The channel to sample when sampling the alpha channel.
	SwizzleBlue    Swizzle    // The channel to sample when sampling the blue channel.
	SwizzleGreen   Swizzle    // The channel to sample when sampling the green channel.
	SwizzleRed     Swizzle    // The channel to sample when sampling the red channel.
}

// Swizzle for textures.
type Swizzle int

const (
	SwizzleIdentity Swizzle = iota // The channel is returned as-is.
	SwizzleZero                    // The channel is always 0.
	SwizzleOne                     // The channel is always 1.
	SwizzleRed                     // Sample the red channel.
	SwizzleGreen                   // Sample the green channel.
	SwizzleBlue                    // Sample the blue channel.
	SwizzleAlpha                   // Sample the alpha channel.
)

// TextureSliceType for textures.
type TextureSliceType int

const (
	TextureSlice2D      TextureSliceType = iota // 2-dimensional texture slice.
	TextureSliceCubemap                         // Cubemap texture slice.
	TextureSlice3D                              // 3-dimensional texture slice.
)

// TextureSamples for textures.
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

// SamplerState for sampling textures.
type SamplerState struct {
	/*
		AnisotropyMax is the maximum anisotropy that can be used when sampling. Only effective if UseAnisotropy is true.
		Higher values result in a sharper sampler at oblique angles, at the cost of performance (due to memory bandwidth).
		This value may be limited by the graphics hardware in use. Most graphics hardware only supports values up to 16.0.

		If AnisotropyMax is 1.0, forcibly disables anisotropy even if UseAnisotropy is true.
	*/
	AnisotropyMax float64
	/*
		The border color that will be returned when sampling outside the sampler's bounds and the RepeatU, RepeatV or
		RepeatW modes have repeating disabled.
	*/
	BorderColor BorderColor
	// Comparison to use. Only effective if EnableCompare is true.
	Comparison Comparison
	/*
		EnableComparison means returned values will be based on the Comparison. This is a hardware-based approach and is
		therefore faster than performing this manually in a shader. For example, compare operations are used for shadow
		map rendering by comparing depth values from a shadow sampler.
	*/
	EnableComparison bool
	/*
		LevelOfDetailBias to use. Positive values will make the sampler blurrier at a given distance, while negative values
		will make the sampler sharper at a given distance (at the risk of looking grainy). Recommended values are between
		-0.5 and 0.0. Only effective if the sampler has mipmaps available.
	*/
	LevelOfDetailBias   float64
	MagnificationFilter Filter
	MaxLevelOfDetail    float64 // Only effective if the sampler has mipmaps available.
	MinificationFilter  Filter
	MinLevelOfDetail    float64 // Only effective if the sampler has mipmaps available.
	MipmapFilter        Filter
	RepeatU             RepeatMode // The repeat mode to use along the U axis of UV coordinates. This affects the returned values if sampling outside the UV bounds.
	RepeatV             RepeatMode // The repeat mode to use along the V axis of UV coordinates. This affects the returned values if sampling outside the UV bounds.
	RepeatW             RepeatMode // The repeat mode to use along the W axis of UV coordinates. This affects the returned values if sampling outside the UV bounds.
	UnnormalizedUVW     bool
	UseAnisotropy       bool
}

// BorderColor options.
type BorderColor int

const (
	BorderColorFloatTransparentBlack BorderColor = iota
	BorderColorInt64TransparentBlack
	BorderColorFloatOpaqueBlack
	BorderColorInt64OpaqueBlack
	BorderColorFloatOpaqueWhite
	BorderColorInt64OpaqueWhite
)

// Filter for sampling.
type Filter int

const (
	FilterNearest Filter = iota
	FilterLinear
)

// RepeatMode when sampling.
type RepeatMode int

const (
	Repeat RepeatMode = iota
	RepeatMirrored
	ClampToEdge
	ClampToBorder
	ClampToEdgeMirrored
)
