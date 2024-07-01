package rd

import (
	"time"

	"grow.graphics/uc"
	"grow.graphics/xy"
	"runtime.link/ffi"
	"runtime.link/mmm"
)

type Interface interface {
	Barrier(from, to BarrierMask)

	BufferClear(buffer Buffer, offset, size_bytes int, post_barrier BarrierMask) error
	BufferGetData(lifetime mmm.Lifetime, buffer Buffer, offset_bytes, size_bytes int) ffi.Bytes
	BufferUpdate(buffer Buffer, offset, size_bytes int, data ffi.Bytes, post_barrier BarrierMask) error

	CaptureTimestamp(name string)
	GetCapturedTimestampTimeCPU(index int) time.Duration
	GetCapturedTimestampTimeGPU(index int) time.Duration
	GetCapturedTimestampName(index int) string
	GetCapturedTimestampsCount() int
	GetCapturedTimestampsFrame() int

	ComputeListBegin(allow_draw_overlap bool) ComputeList
	ComputeListAddBarrier(list ComputeList)
	ComputeListBindComputePipeline(list ComputeList, pipeline ComputePipeline)
	ComputeListBindUniformSet(list ComputeList, set UniformSet, set_index int)
	ComputeListDispatch(list ComputeList, groups_x, groups_y, groups_z int)
	ComputeListEnd(mask BarrierMask)
	ComputeListSetPushConstant(list ComputeList, buffer ffi.Bytes, size_bytes int)

	ComputePipelineCreate(span mmm.Lifetime, shader Shader, constants ffi.Managed[[]PipelineSpecializationConstant]) ComputePipeline
	ComputePipelineIsValid(pipeline ComputePipeline) bool

	CreateLocalDevice(span mmm.Lifetime) Interface

	DrawCommandBeginLabel(name string, color uc.Color)
	DrawCommandEndLabel()
	DrawCommandInsertLabel(name string, color uc.Color)

	DrawListBegin(fb Framebuffer,
		initial_color_action InitialAction, final_color_action FinalAction,
		initial_depth_action InitialAction, final_depth_action FinalAction,
		clear_color_values ffi.Slice[uc.Color], clear_depth float64, clear_stencil int,
		region xy.Rect2, storage_textures ffi.Managed[[]Texture],
	)
	DrawListBeginForScreen(screen Screen, clear_color uc.Color) DrawList
	DrawListBeginSplit(span mmm.Lifetime, fb Framebuffer, splits int,
		initial_color_action InitialAction, final_color_action FinalAction,
		initial_depth_action InitialAction, final_depth_action FinalAction,
		clear_color_values ffi.Slice[uc.Color], clear_depth float64, clear_stencil int,
		region xy.Rect2, storage_textures ffi.Managed[[]Texture],
	) ffi.Slice[DrawList]
	DrawListBindIndexArray(list DrawList, array IndexArray)
	DrawListBindRenderPipeline(list DrawList, pipeline RenderPipeline)
	DrawListBindUniformSet(list DrawList, set UniformSet, set_index int)
	DrawListBindVertexArray(list DrawList, array VertexArray)
	DrawListDisableScissor(list DrawList)
	DrawListDraw(list DrawList, use_indices bool, instances int, procedural_vertex_count int)
	DrawListEnableScissor(list DrawList, rect xy.Rect2)
	DrawListEnd(mask BarrierMask)
	DrawListSetBlendConstants(list DrawList, constants uc.Color)
	DrawListSetPushConstant(list DrawList, buffer ffi.Bytes, size_bytes int)
	DrawListSwitchToNextPass() DrawList
	DrawListSwitchToNextPassSplit(span mmm.Lifetime, splits int) ffi.Slice[DrawList]

	FramebufferCreate(span mmm.Lifetime, textures ffi.Managed[[]Texture], validate_with_format, view_count int) Framebuffer
	FramebufferCreateEmpty(span mmm.Lifetime, size xy.Vector2i, samples TextureSamples, validate_with_format int) Framebuffer
	FramebufferCreateMultipass(span mmm.Lifetime, textures ffi.Managed[[]Texture], passes ffi.Managed[[]FramebufferPass], validate_with_format, view_count int) Framebuffer
	FramebufferGetFormat(fb Framebuffer) FramebufferFormat
	FramebufferIsValid(Framebuffer) bool

	FramebufferFormatCreate(attachments ffi.Managed[[]AttachmentFormat], view_count int) FramebufferFormat
	FramebufferFormatCreateEmpty(samples TextureSamples) FramebufferFormat
	FramebufferFormatCreateMultipass(attachments ffi.Managed[[]AttachmentFormat], passes ffi.Managed[[]FramebufferPass], view_count int) FramebufferFormat
	FramebufferFormatGetTextureSamples(format1 FramebufferFormat, render_pass int) TextureSamples

	FullBarrier()
	GetDeviceName() string
	GetDevicePipelineCacheUUID() string
	GetDeviceVendorName() string
	GetDriverResource(resource DriverResource, rid RID, index int) int
	GetFrameDelay() int
	GetMemoryUsage(mtype MemoryType) int

	IndexArrayCreate(span mmm.Lifetime, buf IndexBuffer, offset, count int) IndexArray
	IndexBufferCreate(span mmm.Lifetime, size_indices int, format IndexBufferFormat, data ffi.Bytes, use_restart_indicies bool) IndexBuffer

	LimitGet(limit Limit) int

	RenderPipelineCreate(span mmm.Lifetime, shader Shader, framebuffer_format FramebufferFormat, vertex_format VertexFormat,
		primitive RenderPrimitive, rasterization_state ffi.Managed[PipelineRasterizationState], multisample_state ffi.Managed[PipelineMultisampleState],
		stencil_state ffi.Managed[PipelineDepthStencilState], color_blend_state ffi.Managed[PipelineColorBlendState], dynamic_state_flags PipelineDynamicStateFlags,
		for_render_pass int, specialization_constants ffi.Managed[[]PipelineSpecializationConstant],
	) RenderPipeline
	RenderPipelineIsValid(pipeline RenderPipeline) bool

	SamplerCreate(span mmm.Lifetime, state ffi.Managed[SamplerState]) Sampler
	SamplerIsFormatSupportedForFilter(format DataFormat, filter SamplerFilter) bool

	ScreenGetFramebufferFormat() FramebufferFormat
	ScreenGetHeight(screen Screen) int64
	ScreenGetWidth(screen Screen) int64

	SetResourceName(resource RID, name string)

	ShaderCompileBinaryFromSPIRV(span mmm.Lifetime, shader ffi.Managed[ShaderSPIRV], name string) ffi.Bytes
	ShaderCompileSourceIntoSPIRV(span mmm.Lifetime, shader ffi.Managed[ShaderSource], allow_cache bool) ffi.Managed[ShaderSPIRV]
	ShaderCreateFromBytecode(span mmm.Lifetime, code ffi.Bytes, id ShaderPlaceholder) Shader
	ShaderCreateFromSPIRV(span mmm.Lifetime, shader ffi.Managed[ShaderSPIRV], name string) Shader
	ShaderCreatePlaceholder(span mmm.Lifetime) ShaderPlaceholder
	ShaderGetVertexInputAttributeMask(shader Shader) int64

	StorageBufferCreate(span mmm.Lifetime, size_bytes int, data ffi.Bytes, usage_flags StorageBufferUsage) StorageBuffer

	Submit()
	Sync()

	TextureBufferCreate(span mmm.Lifetime, size_bytes int, data_format DataFormat, data ffi.Bytes) TextureBuffer
	TextureClear(texture Texture, color uc.Color, base_mipmap, mipmap_count, base_layer, layer_count int, post_barrier BarrierMask) error
	TextureCopy(src, dst Texture, src_pos, dst_pos, size xy.Vector3, src_mipmap, dst_mipmap, src_layer, dst_layer int, post_barrier BarrierMask) error
	TextureCreate(span mmm.Lifetime, format ffi.Managed[TextureFormat], view ffi.Managed[TextureView], data ffi.Managed[[]ffi.Bytes]) Texture
	TextureCreateFromExtension(span mmm.Lifetime,
		texture_type TextureType, format DataFormat, samples TextureSamples, usage_flags TextureUsage,
		image, width, height, depth, layers int,
	) Texture
	TextureCreateShared(span mmm.Lifetime, view ffi.Managed[TextureView], with_texture Texture) Texture
	TextureCreateSharedFromSlice(span mmm.Lifetime, view ffi.Managed[TextureView], with_texture Texture, layer, mipmap, mipmaps int, slice_type TextureSliceType) Texture
	TextureGetData(span mmm.Lifetime, texture Texture, layer int) ffi.Bytes
	TextureGetFormat(tex Texture) TextureFormat
	TextureGetNativeHandle(tex Texture) uint64
	TextureIsFormatSupportedForUsage(format DataFormat, usage TextureUsage) bool
	TextureIsShared(tex Texture) bool
	TextureIsValid(tex Texture) bool
	TextureResolveMultisample(src, dst Texture, post_barrier BarrierMask) error
	TextureUpdate(tex Texture, layer int, data ffi.Bytes, post_barrier BarrierMask) error

	UniformBufferCreate(span mmm.Lifetime, size_bytes int, data ffi.Bytes) UniformBuffer
	UniformSetCreate(span mmm.Lifetime, uniforms ffi.Managed[[]Uniform], shader Shader, shader_set int) UniformSet
	UniformSetIsValid(set UniformSet) bool

	VertexArrayCreate(span mmm.Lifetime, vertex_count int, vertex_format VertexFormat, src_buffers ffi.Managed[[]Buffer], offsets ffi.Slice[int64]) VertexArray
	VertexFormatCreate(span mmm.Lifetime, vertex_descriptions ffi.Managed[[]VertexAttribute]) VertexFormat

	FreeRID(uintptr)
}

type ShaderPlaceholder mmm.Pointer[Interface, ShaderPlaceholder, uintptr]

func (placeholder ShaderPlaceholder) Free() {
	(*mmm.API(placeholder)).FreeRID(mmm.End(placeholder))
}

type RID uintptr

type FramebufferPass struct {
	ColorAttachments    ffi.Slice[int32]
	DepthAttachment     int
	InputAttachments    ffi.Slice[int32]
	PreserveAttachments ffi.Slice[int32]
	ResolveAttachments  ffi.Slice[int32]
}

type Screen int64

type FramebufferFormat int64

type DrawList int64

type ComputeList int64

type VertexFormat int64

type ComputePipeline mmm.Pointer[Interface, ComputePipeline, uintptr]

func (pipeline ComputePipeline) Free() {
	(*mmm.API(pipeline)).FreeRID(mmm.End(pipeline))
}

type UniformSet mmm.Pointer[Interface, UniformSet, uintptr]

func (set UniformSet) Free() {
	(*mmm.API(set)).FreeRID(mmm.End(set))
}

type Framebuffer mmm.Pointer[Interface, Framebuffer, uintptr]

func (fb Framebuffer) Free() { (*mmm.API(fb)).FreeRID(mmm.End(fb)) }

type IndexArray mmm.Pointer[Interface, IndexArray, uintptr]

func (array IndexArray) Free() { (*mmm.API(array)).FreeRID(mmm.End(array)) }

type VertexArray mmm.Pointer[Interface, VertexArray, uintptr]

func (array VertexArray) Free() { (*mmm.API(array)).FreeRID(mmm.End(array)) }

type RenderPipeline mmm.Pointer[Interface, RenderPipeline, uintptr]

func (pipeline RenderPipeline) Free() { (*mmm.API(pipeline)).FreeRID(mmm.End(pipeline)) }

type Shader mmm.Pointer[Interface, Shader, uintptr]

func (shader Shader) Free() { (*mmm.API(shader)).FreeRID(mmm.End(shader)) }

type Sampler mmm.Pointer[Interface, Sampler, uintptr]

func (sampler Sampler) Free() { (*mmm.API(sampler)).FreeRID(mmm.End(sampler)) }

type BarrierMask uint32

const (
	BarrierMaskVertex BarrierMask = 1 << iota
	BarrierMaskCompute
	BarrierMaskTransfer
	BarrierMaskFragment

	BarrierMaskRaster = BarrierMaskVertex | BarrierMaskFragment

	BarrierMaskAll       = 32767
	BarrierMaskNoBarrier = 32768
)

type InitialAction int

const (
	InitialActionClear InitialAction = iota
	InitialActionClearRegion
	InitialActionClearRegionContinue
	InitialActionKeep
	InitialActionDrop
	InitialActionContinue
)

type FinalAction int

const (
	FinalActionRead FinalAction = iota
	FinalActionDiscard
	FinalActionContinue
)

type AttachmentFormat struct {
	Format     DataFormat
	Samples    TextureSamples
	UsageFlags int
}

type DriverResource int

const (
	DriverResourceVulkanDevice DriverResource = iota
	DriverResourceVulkanPhysicalDevice
	DriverResourceVulkanInstance
	DriverResourceVulkanQueue
	DriverResourceVulkanQueueFamilyIndex
	DriverResourceVulkanImage
	DriverResourceVulkanImageView
	DriverResourceVulkanImageNativeTextureFormat
	DriverResourceVulkanSampler
	DriverResourceVulkanDescriptorSet
	DriverResourceVulkanBuffer
	DriverResourceVulkanComputePipeline
	DriverResourceVulkanRenderPipeline
)

type MemoryType int

const (
	MemoryTextures MemoryType = iota
	MemoryBuffers
	MemoryTotal
)

type IndexBufferFormat int

const (
	IndexBufferUint16 IndexBufferFormat = iota
	IndexBufferUint32
)

type RenderPrimitive int

const (
	RenderPrimitivePoints RenderPrimitive = iota
	RenderPrimitiveLines
	RenderPrimitiveLinesWithAdjacency
	RenderPrimitiveLineStrips
	RenderPrimitiveLineStripsWithAdjacency
	RenderPrimitiveTriangles
	RenderPrimitiveTrianglesWithAdjacency
	RenderPrimitiveTriangleStrips
	RenderPrimitiveTriangleStripsWithAdjacency
	RenderPrimitiveTriangleStripsWithRestartIndex
	RenderPrimitiveTessellationPatch
)

type PipelineRasterizationState struct {
	CullMode                PolygonCullMode
	DepthBiasClamp          float64
	DepthBiasConstantFactor float64
	DepthBiasEnabled        bool
	DepthBiasSlopeFactor    float64
	DiscardPrimitives       bool
	EnableDepthClamp        bool
	FrontFace               PolygonFrontFace
	LineWidth               float64
	PatchControlPoints      int64
	Wireframe               bool
}

type PolygonCullMode int

const (
	PolygonCullDisabled PolygonCullMode = iota
	PolygonCullFront
	PolygonCullBack
)

type PolygonFrontFace int

const (
	PolygonFrontFaceClockwise PolygonFrontFace = iota
	PolygonFrontFaceCounterClockwise
)

type PipelineMultisampleState struct {
	EnableAlphaToCoverage bool
	EnableAlphaToOne      bool
	EnableSampleShading   bool
	MinSampleShading      float64
	SampleCount           TextureSamples
	SampleMasks           ffi.Managed[[]int64]
}

type PipelineColorBlendState struct {
	Attachments          ffi.Managed[[]PipelineColorBlendStateAttachment]
	BlendConstant        uc.Color
	EnableLogicOperation bool
	LogicOperation       LogicOperation
}

type PipelineColorBlendStateAttachment struct {
	AlphaBlendOperation         BlendOperation
	ColorBlendOperation         BlendOperation
	DestinationAlphaBlendFactor BlendFactor
	DestinationColorBlendFactor BlendFactor
	EnableBlend                 bool
	SourceAlphaBlendFactor      BlendFactor
	SourceColorBlendFactor      BlendFactor
	WriteA                      bool
	WriteB                      bool
	WriteG                      bool
	WriteR                      bool
}

type LogicOperation int

const (
	LogicOperationClear LogicOperation = iota
	LogicOperationAnd
	LogicOperationAndReverse
	LogicOperationCopy
	LogicOperationAndInverted
	LogicOperationNoop
	LogicOperationXor
	LogicOperationOr
	LogicOperationNor
	LogicOperationEquivalent
	LogicOperationInvert
	LogicOperationOrReverse
	LogicOperationCopyInverted
	LogicOperationOrInverted
	LogicOperationNand
	LogicOperationSet
)

type BlendOperation int

const (
	BlendOperationAdd BlendOperation = iota
	BlendOperationSubtract
	BlendOperationReverseSubtract
	BlendOperationMinimum
	BlendOperationMaximum
)

type BlendFactor int

const (
	BlendFactorZero BlendFactor = iota
	BlendFactorOne
	BlendFactorSourceColor
	BlendFactorOneMinusSourceColor
	BlendFactorDestinationColor
	BlendFactorOneMinusDestinationColor
	BlendFactorSourceAlpha
	BlendFactorOneMinusSourceAlpha
	BlendFactorDestinationAlpha
	BlendFactorOneMinusDestinationAlpha
	BlendFactorConstantColor
	BlendFactorOneMinusConstantColor
	BlendFactorConstantAlpha
	BlendFactorOneMinusConstantAlpha
	BlendFactorSourceAlphaSaturate
	BlendFactorSource1Color
	BlendFactorOneMinusSource1Color
	BlendFactorSource1Alpha
	BlendFactorOneMinusSource1Alpha
)

type PipelineDepthStencilState struct {
	BackOperationCompare      CompareOperator
	BackOperationCompareMask  int64
	BackOperationDepthFail    StencilOperation
	BackOperationFail         StencilOperation
	BackOperationPass         StencilOperation
	BackOperationReference    int64
	BackOperationWriteMask    int64
	DepthCompareOperator      CompareOperator
	DepthRangeMax             float64
	DepthRangeMin             float64
	EnableDepthRange          bool
	EnableDepthTest           bool
	EnableDepthWrite          bool
	EnableStencil             bool
	FrontOperationCompare     CompareOperator
	FrontOperationCompareMask int64
	FrontOperationDepthFail   StencilOperation
	FrontOperationFail        StencilOperation
	FrontOperationPass        StencilOperation
	FrontOperationReference   int64
	FrontOperationWriteMask   int64
}

type CompareOperator int

const (
	CompareOperatorNever CompareOperator = iota
	CompareOperatorLess
	CompareOperatorEqual
	CompareOperatorLessOrEqual
	CompareOperatorGreater
	CompareOperatorNotEqual
	CompareOperatorGreaterOrEqual
	CompareOperatorAlways
)

type StencilOperation int

const (
	StencilOperationKeep StencilOperation = iota
	StencilOperationZero
	StencilOperationReplace
	StencilOperationIncrementAndClamp
	StencilOperationDecrementAndClamp
	StencilOperationInvert
	StencilOperationIncrementAndWrap
	StencilOperationDecrementAndWrap
)

type PipelineDynamicStateFlags int

const (
	PipelineDynamicLineWidth PipelineDynamicStateFlags = 1 << iota
	PipelineDynamicDepthBias
	PipelineDynamicBlendConstants
	PipelineDynamicDepthBounds
	PipelineDynamicStencilCompareMask
	PipelineDynamicStencilWriteMask
	PipelineDynamicStencilReference
)

type PipelineSpecializationConstant struct {
	ConstantID int
	Value      any
}

type SamplerState struct {
	AnisotropyMax       float64
	BorderColor         SamplerBorderColor
	CompareOperator     CompareOperator
	EnableCompare       bool
	LevelOfDetailBias   float64
	MagnificationFilter SamplerFilter
	MaxLevelOfDetail    float64
	MinificationFilter  SamplerFilter
	MinLevelOfDetail    float64
	MipmapFilter        SamplerFilter
	RepeatU             SamplerRepeatMode
	RepeatV             SamplerRepeatMode
	RepeatW             SamplerRepeatMode
	UnnormalizedUVW     bool
	UseAnisotropy       bool
}

type SamplerBorderColor int

const (
	SamplerBorderColorFloatTransparentBlack SamplerBorderColor = iota
	SamplerBorderColorIntTransparentBlack
	SamplerBorderColorFloatOpaqueBlack
	SamplerBorderColorIntOpaqueBlack
	SamplerBorderColorFloatOpaqueWhite
	SamplerBorderColorIntOpaqueWhite
)

type SamplerFilter int

const (
	SamplerFilterNearest SamplerFilter = iota
	SamplerFilterLinear
)

type SamplerRepeatMode int

const (
	SamplerRepeatModeRepeat SamplerRepeatMode = iota
	SamplerRepeatModeMirroredRepeat
	SamplerRepeatModeClampToEdge
	SamplerRepeatModeClampToBorder
	SamplerRepeatModeMirrorClampToEdge
)

type ShaderSPIRV struct {
	Compute               ffi.Bytes
	Fragment              ffi.Bytes
	TesselationControl    ffi.Bytes
	TesselationEvaluation ffi.Bytes
	Vertex                ffi.Bytes
}

type ShaderSource struct {
	Language              ShaderLanguage
	Compute               ffi.Bytes
	Fragment              ffi.Bytes
	TesselationControl    ffi.Bytes
	TesselationEvaluation ffi.Bytes
	Vertex                ffi.Bytes
}

type ShaderLanguage int

const (
	ShaderLanguageGLSL ShaderLanguage = iota
	ShaderLanguageHLSL
)

type StorageBufferUsage int

const (
	StorageBufferDispatchIndirect StorageBufferUsage = 1 << iota
)

type Uniform struct {
	Binding     int
	UniformType UniformType
}

type UniformType int

const (
	UniformTypeSampler UniformType = iota
	UniformTypeSamplerWithTexture
	UniformTypeTexture
	UniformTypeImage
	UniformTypeTextureBuffer
	UniformTypeSamplerWithTextureBuffer
	UniformTypeImageBuffer
	UniformTypeUniformBuffer
	UniformTypeStorageBuffer
	UniformTypeInputAttachment
)

type VertexAttribute struct {
	Format    DataFormat
	Frequency VertexFrequency
	Location  int
	Offset    int
	Stride    int
}

type VertexFrequency int

const (
	VertexFrequencyVertex VertexFrequency = iota
	VertexFrequencyInstance
)
