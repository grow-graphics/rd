// Package rd provides a rendering device interface for low-level rendering operations.
package rd

import (
	"io"
	"time"

	"grow.graphics/uc"
	"grow.graphics/xy"
)

// Interface to the rendering device.
type Interface interface {
	// Barrier puts a memory barrier in place. This is used for synchronization to avoid data races.
	// See also [Interface.FullBarrier], which may be useful for debugging.
	Barrier(from, upto Barrier)

	// BarrierFull puts a full memory barrier in place. This is a memory barrier with all flags enabled.
	// It should only be used for debugging as it can severely impact performance.
	BarrierFull()

	// CaptureTimestamp creates a timestamp marker with the specified name. This is used for performance
	// reporting with the [Timestamp.CPU], [Timestamp.GPU] and [Timestamp.Name] methods.
	CaptureTimestamp(name string) Timestamp

	// CompileBinary creates a new shader instance from a binary compiled shader.
	CompileBinary(data []byte) Shader

	/*
		Compiles a binary shader from spirv and returns the compiled binary data as bytes. This compiled
		shader is specific to the GPU model and driver version used; it will not work on different GPU models
		or even different driver versions. See also [Interface.CompileSource].

		'name' is an optional human-readable name that can be given to the compiled shader for organizational
		purposes.
	*/
	CompileSPIRV(name string, spriv SPIRV) []byte

	/*
		Compiles a SPIR-V from the shader source code in shader_source and returns the SPIR-V as a [ShaderSPIRV].
		This intermediate language shader is portable across different GPU models and driver versions, but cannot
		be run directly by GPUs until compiled into a binary shader using [Interface.CompileSPIRV].

		If 'cache' is true, make use of the shader cache. This avoids a potentially lengthy shader compilation
		step if the shader is already in cache. If 'cache' is false, the shader cache is ignored and the shader
		will always be recompiled.
	*/
	CompileSource(cache bool, source ShaderSource) SPIRV

	/*
		Compute is used to prepare a list of compute commands with [ComputeList] methods.
		You may not call this method from within 'fn'.

		A simple compute operation might look like this (code is not a complete example):

			RD.Compute(func(compute rd.ComputeList) {
				compute.BindPipeline(computeShaderDilatePipeline)
				compute.BindUniformSet(computeBaseUniformSet, 0)
				compute.BindUniformSet(dilateUniformSet, 1)
				for i := range atlasSlices {
					compute.SetPushConstant(pushConstant, pushConstant.Size())
					compute.Dispatch(groupSize.X, groupSize.Y, groupSize.Z)
				}
			})
	*/
	Compute(fn func(Compute))

	// DeviceName returns the name of the video adapter (e.g. "GeForce GTX 1080/PCIe/SSE2").
	DeviceName() string

	// DeviceVendor returns the name of the video adapter vendor (e.g. "NVIDIA Corporation").
	DeviceVendor() string

	/*
		Drawing starts a list of raster drawing commands created with the [Drawing] methods.

		A simple drawing operation might look like this (code is not a complete example):

			target := rd.DrawingTarget{
				Framebuffer: framebuffers[i],
				InitialColorAction: rd.InitialActionClear,
				FinalColorAction: rd.FinalActionRead,
				InitialDepthAction: rd.InitialActionClear,
				FinalDepthAction: rd.FinalActionDiscard,
				ClearColorValues: []uc.Color{{},{},{}},
			}
			RD.Drawing(target, func(drawing rd.Drawing) {
				drawing.SetShader(shader)
				drawing.SetVariables(rd.VariablesForFrame, globalVariables)
				drawing.SetData(data)
				drawing.Submit(false, 1, slice_triangle_count[i] * 3)
			})
	*/
	Drawing(frame Frame, fn func(Drawing))

	/*
	   DrawingOnScreen is a high-level variant of [Interface.Drawing], with the [Frame] automatically
	   set for drawing onto the window specified by the screen ID.

	   Note: Cannot be used with local rendering devices, as these don't have a screen. If called on
	   a local RenderingDevice, [DrawingOnScreen] panics.
	*/
	DrawingOnScreen(screen Screen, clear uc.Color, fn func(Drawing))

	// ExtensionTexture returns the texture for an existing image (VkImage) with the given type, format, samples,
	// usage_flags, width, height, depth, and layers. This can be used to allow the rendering device to render onto foreign images.
	ExtensionTexture(ttype TextureType, format DataFormat, samples TextureSamples, usage TextureUsage, image uintptr, width, height, depth, layers int) Texture

	// FrameDelay returns the frame count kept by the graphics API. Higher values result in higher input lag, but with
	// more consistent throughput. For the main RenderingDevice, frames are cycled (usually 3 with triple-buffered V-Sync
	// enabled). However, local RenderingDevices only have 1 frame.
	FrameDelay() int

	/*
		FramebufferFormat creates a new framebuffer format with the specified attachments and eyes.

		If 'eyes' is greater than or equal to 2, enables multiview which is used for VR rendering.
		In Vulkan, this requires support for the multiview extension.
	*/
	FramebufferFormat(eyes int, attachments []AttachmentFormat, passes []FramebufferPass) FramebufferFormat

	// IndexBufferU16 creates a new uint16 index buffer.
	IndexBufferU16(data []uint16) IndexBuffer

	// IndexBufferU32 creates a new uint32 index buffer.
	IndexBufferU32(data []uint32) IndexBuffer

	// Limit returns the value of the specified limit. This limit varies depending on the current graphics hardware (and
	// sometimes the driver version). If the given limit is exceeded, rendering errors will occur.
	Limit(limit Limit) int

	// MemoryUsage returns the current memory usage of the rendering device.
	MemoryUsage(mtype MemoryType) int

	// PipelineCache returns the universally unique identifier for the pipeline cache. This is used to cache shader files on disk,
	// which avoids shader recompilations on subsequent engine runs. This UUID varies depending on the graphics card
	// model, but also the driver version. Therefore, updating graphics drivers will invalidate the shader cache.
	PipelineCache() string

	// Processor creates a new [Processor] for [Compute]. Defines can be boolean and/or numeric values.
	Processor(shader Shader, defines []any) Processor

	// Renderer creates a new [Renderer] for [Drawing].
	Renderer(shader Shader, options RenderingOptions) Renderer

	// Local create a new local rendering device. This is most useful for performing
	// compute operations on the GPU independently from the rest of the program.
	RenderingDevice() Local

	// Sampler creates a new sampler.
	Sampler(state SamplerState) Sampler

	// Screen returns the Nth screen.
	Screen(n int) Screen

	/*
		Shader creates a placeholder [Shader] without initializing. This allows you to create a [Shader]
		and pass it around, but defer compiling it to a later time.
	*/
	Shader() Shader

	// SharedTexture creates a shared texture using the specified view and the texture information from 'with'.
	SharedTexture(view TextureView, with Texture) Texture

	// StorageBuffer creates a storage buffer with the specified data and usage.
	StorageBuffer(usage StorageBufferUsage, data []byte) StorageBuffer

	// Texture creates a new texture with the specified parameters.
	Texture(format TextureFormat, view TextureView, data [][]byte) Texture

	// TextureBuffer creates a new texture buffer.
	TextureBuffer(format DataFormat, data []byte) TextureBuffer

	/*
		TextureCopy copies the src to dst with the specified from, into position and size coordinates.
		The Z axis of the from, into and size must be 0 for 2-dimensional textures. Source and
		destination mipmaps/layers must also be specified, with these parameters being 0 for textures
		without mipmaps or single-layer textures.

		Note: src texture can't be copied while a draw list that uses it as part of a framebuffer is being
		created. Ensure the draw list is finalized (and that the color/depth texture using it is not set
		to [FrameSuspend]) to copy this texture.

		Note: src texture requires the [TextureCanCopyFrom] to be retrieved.

		Note: dst can't be copied while a draw list that uses it as part of a framebuffer is being created.
		Ensure the draw list is finalized (and that the color/depth texture using it is not set to
		[FrameSuspend]) to copy this texture.

		Note: dst requires the [TextureCanCopyInto] to be retrieved.

		Note: src and dst must be of the same type (color or depth).
	*/
	TextureCopy(src, dst Texture, from, into, size xy.Vector3, src_mipmap, dst_mipmp, src_layer, dst_layer int, barrier Barrier) error

	// TextureFormatIsSupportedForUsage returns true if the specified format is supported for the given usage, false otherwise.
	TextureFormatIsSupportedForUsage(format DataFormat, usage TextureUsage) bool

	/*
		TextureResolveMultiSample resolves the from_texture texture onto to_texture with multisample antialiasing enabled.
		This must be used when rendering a framebuffer for MSAA to work.

		Note: from and into textures must have the same dimension, format and type (color or depth).

		Note: from can't be copied while a draw list that uses it as part of a framebuffer is being created.
		Ensure the draw list is finalized (and that the color/depth texture using it is not set to
		[FrameSuspend]) to resolve this texture.

		Note: from requires the [TextureCanCopyFrom] to be retrieved.

		Note: from must be multisampled and must also be 2D (or a slice of a 3D/cubemap texture).

		Note: into can't be copied while a draw list that uses it as part of a framebuffer is being created.
		Ensure the draw list is finalized (and that the color/depth texture using it is not set to
		[FrameSuspend]) to resolve this texture.

		Note: into texture requires the [TextureCanCopyInto] to be retrieved.

		Note: into texture must not be multisampled and must also be 2D (or a slice of a 3D/cubemap texture).
	*/
	TextureResolveMultiSample(from, into Texture, barrier Barrier) error

	// UniformBuffer creates a new uniform buffer with the specified data.
	UniformBuffer(data []byte) UniformBuffer

	// VertexArray creates a new vertex from the specified buffers, optionally, with offsets.
	VertexArray(vertices int, format VertexFormat, buffers []Buffer, offsets []int64) VertexArray

	// VertexBuffer creates a new vertex buffer with the specified data.
	VertexBuffer(data []byte) VertexBuffer

	// VertexFormat creates a new vertex format with the specified attributes.
	VertexFormat(attributes []VertexAttribute) VertexFormat
}

// Barrier used for syncronisation.
type Barrier uint32

const (
	BarrierVertex Barrier = 1 << iota
	BarrierCompute
	BarrierTransfer
	BarrierFragment

	BarrierRaster = BarrierVertex | BarrierFragment

	BarrierFull    = 32767
	BarrierDisable = 32768 // no barrier for any type.
)

// MemoryType classifies memory usage.
type MemoryType int

const (
	MemoryTextures MemoryType = iota // texture related memory in use.
	MemoryBuffers                    // buffers in use.
	MemoryTotal                      // total GPU memory.
)

// Local rendering device, can be used in an independent thread.
type Local interface {
	Interface

	// Submit pushes the frame setup and draw command buffers then marks the local device as currently
	// processing (which allows calling [Local.Sync]).
	Submit()

	// Sync forces a synchronization between the CPU and GPU, which may be required in certain cases.
	// Only call this when needed, as CPU-GPU synchronization has a performance cost.
	Sync()
}

// AttachmentFormat describes the format of a framebuffer attachment.
type AttachmentFormat struct {
	Format  DataFormat     // The attachment's data format.
	Samples TextureSamples // The number of samples used when sampling the attachment.
	Usage   TextureUsage   // The attachment's usage flags, which determine what can be done with it.
}

// FramebufferPass contains the list of attachment descriptions for a framebuffer pass. Each points with an
// index to a previously supplied list of texture attachments.
//
// Multipass framebuffers can optimize some configurations in mobile. On desktop, they provide little to no
// advantage.
type FramebufferPass struct {
	ColorAttachments      []int32 // Color attachments in order starting from 0.
	DepthAttachment       int32   // Depth attachment. -1 should be used if no depth buffer is required for this pass.
	InputAttachments      []int32 // Used for multipass framebuffers (more than one render pass). Converts an attachment to an input. Make sure to also supply it properly in the [Variables].
	AttachmentsToPreserve []int32 // (otherwise they are erased).
	AttachmentsToResolve  []int32 // If the color attachments are multisampled, non-multisampled resolve attachments can be provided.
}

// Processor for compute operations.
type Processor any

// Drawing operation.
type Drawing interface {
	// DebugBlock creates a command buffer debug label region that can be displayed in third-party tools such as RenderDoc.
	//
	// In Vulkan, the VK_EXT_DEBUG_UTILS_EXTENSION_NAME extension must be available and enabled for command buffer debug
	// label region to work. See also [DebugLabel].
	DebugBlock(name string, color uc.Color, block func())

	// DebugLabel inserts a command buffer debug label region in the current command buffer.
	DebugLabel(name string, color uc.Color)

	// SetBlendConstant sets the blend constant for the next draw command. Only used if the
	// shader is created with [BlendConstants] flag set.
	SetBlendConstant(color uc.Color)

	// SetData sets the push constant data buffer for the specified compute operation. The shader
	// determines how this binary data is used.
	SetData(data []byte)

	// SetIndexArray sets the index array to use for the next draw command.
	SetIndexArray(array IndexArray)

	// SetRenderer sets the renderer to use for the next draw command.
	SetRenderer(r Renderer)

	// SetScissor sets the scissor region for the next draw command. Setting the
	// region to nil will disable the scissor. An empty region will default to
	// the framebuffer size.
	SetScissor(region *xy.Rect2)

	// SetVariables binds the [Variables] to the drawing operation.
	SetVariables(level VariableLevel, variables Variables)

	// SetVertexArray sets the vertex array to use for the next draw command.
	SetVertexArray(array VertexArray)

	// Submit the drawing for rendering on the GPU. This is the raster equivalent to [Compute.Submit].
	Submit(indices bool, instances, vertices int)

	// SwitchToNextPass switches to the next draw pass.
	SwitchToNextPass()
}

// FrameStart action.
type FrameStart int

const (
	FrameClear             FrameStart = iota // Start rendering and clear the whole framebuffer.
	FrameClearRegion                         // Start rendering and clear the framebuffer in the specified region.
	FrameClearRegionResume                   // Continue suspended rendering and clear the framebuffer in the specified region.
	FrameKeep                                // Start rendering, but keep attached color texture contents.
	FrameWrite                               // Start rendering, ignore what is there; write above it.
	FrameResume                              // Continue suspended rendering.
)

// FrameEnded action.
type FrameEnded int

const (
	FrameRead    FrameEnded = iota // Store the texture for reading and make it read-only.
	FrameDrop                      // Discard the texture data and make it read-only.
	FrameSuspend                   // Store the texture so that drawing can be resumed later.
)

// Frame options.
type Frame struct {
	Buffer  Framebuffer
	Color   func() (FrameStart, FrameEnded)
	Depth   func() (FrameStart, FrameEnded)
	Clear   Clear
	Region  xy.Rect2
	Storage []Texture
}

// Clear values that the buffers will be cleared to.
type Clear struct {
	Colors  []uc.Color
	Depth   float64
	Stencil int
}

type FramebufferFormat interface {
	// Framebuffer creates a new framebuffer out of the specified textures.
	Framebuffer(textures []Texture) Framebuffer

	// TextureSamples returns the number of texture samples used for the given framebuffer.
	TextureSamples(pass int) TextureSamples
}

type Framebuffer interface {
	// IsValid returns true if the framebuffer is valid, false otherwise.
	IsValid() bool
}

// IndexArray for storing mesh indices.
type IndexArray any

// VertexArray for storing mesh vertices.
type VertexArray any

// Compute operation.
type Compute interface {
	// SetData sets the push constant data buffer for the specified compute operation. The shader
	// determines how this binary data is used.
	SetData(data []byte)

	/*
		SetProcessor sets the processor to use when processing the compute operation. If the
		shader has changed since the last time this function was called, the rendering device will
		unbind all descriptor sets and will re-bind them inside [Compute.Submit].
	*/
	SetProcessor(p Processor)

	/*
		SetVariables binds the [Variables] to the compute operation. The rendering device ensures
		that all textures in the uniform set have the correct access masks. If the rendering device had
		to change access masks of textures, it will raise a image memory barrier.
	*/
	SetVariables(level VariableLevel, variables Variables)

	// Submit submits the compute list for processing on the GPU. This is the compute equivalent
	// to [Drawing.Submit].
	Submit(x, y, z int)
}

// Timestamp for performance monitoring.
type Timestamp interface {
	// CPU returns the time since the rendering device started until 't' or false
	// if the timestamp is not available.
	CPU() (time.Duration, bool)

	// GPU returns the time since the rendering device started until 't' or false
	// if the timestamp is not available.
	GPU() (time.Duration, bool)
}

// Resource allocated by the rendering device.
type Resource interface {
	// RID returns the underlying handle for the resource.
	RID() uint64

	// Free any resources associated with the resource. Subsequent use will result in a panic.
	Free()
}

// Buffer that stores data in GPU memory.
type Buffer interface {
	Resource
	Nameable

	/*
	   Clear the contents of the buffer. Always raises a memory barrier.

	   Returns an error if:

	     - the size isn't a multiple of four

	     - the region specified by offset + size_bytes exceeds the buffer

	     - a draw list is currently active (created by [Interface.DrawList])

	     - a compute list is currently active (created by [Interface.ComputeList])
	*/
	Clear() error

	io.ReaderAt
	io.WriterAt
}

// VertexBuffer for storing mesh vertices.
type VertexBuffer interface {
	Buffer
}

// IndexBuffer for storing mesh indices.
type IndexBuffer interface {
	Buffer
}

// UniformBuffer for storing uniform data.
type UniformBuffer interface {
	Variable
	Buffer
}

// StorageBuffer for storing data in GPU memory.
type StorageBuffer interface {
	Variable
	Buffer
}

// Nameable resources can be debugged in RenderDoc.
type Nameable interface {
	// SetResourceName sets the resource name for id to name. This is used for debugging with third-party tools such as RenderDoc.
	SetResourceName(name string)
}

// Screen identifies a display region, such as an OS window.
type Screen interface {
	// FramebufferFormat returns the screen's framebuffer format.
	FramebufferFormat() FramebufferFormat

	// Height returns the window height matching the graphics API context for the given window ID (in pixels).
	Height() int

	// Width returns the window width matching the graphics API context for the given window ID (in pixels).
	Width() int
}

// Sampler for sampling textures.
type Sampler interface {
	Resource
	Nameable
	Variable

	// FormatSupportedForFilter returns true if implementation supports using a texture of format with the given [SamplerFilter].
	FormatSupportedForFilter(format DataFormat, filter Filter) bool
}

// VertexFormat ID.
type VertexFormat int

// StorageBufferUsage flags.
type StorageBufferUsage int

const (
	StorageBufferDispatchIndirect StorageBufferUsage = 1 << iota
)

// VertexAttribute structure.
type VertexAttribute struct {
	Format    DataFormat
	Frequency AttributeFrequency
	Location  int
	Offset    int
	Stride    int
}

// AttributeFrequency
type AttributeFrequency int

const (
	// Vertex attribute addressing is a function of the vertex. This is used to specify the rate at which vertex attributes are pulled from buffers.
	AttributePerVertex AttributeFrequency = iota
	// Vertex attribute addressing is a function of the instance index. This is used to specify the rate at which vertex attributes are pulled from buffers.
	AttributePerInstance
)
