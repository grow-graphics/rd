package rd

import "grow.graphics/uc"

// Renderer captures a shader along with a set of [RenderingOptions].
type Renderer interface {
	Resource

	// IsValid returns true if the render pipeline is valid, false otherwise.
	IsValid() bool
}

// RenderingOptions to configure the renderer.
type RenderingOptions struct {
	FramebufferFormat FramebufferFormat
	VertexFormat      VertexFormat

	PrimitiveType PrimitiveType
	Rasterization Rasterization
	Multisampling Multisampling
	DepthStencils DepthStencils
	ColorBlending ColorBlending
	DynamicStates DynamicStates
	RenderingPass int
	ShaderDefines []any // can be boolean and/or numeric values
}

// PrimitiveType defines how vertices will be interpreted by the renderer.
type PrimitiveType int

const (
	Points                         PrimitiveType = iota // points with constant size, regardless of distance to camera.
	Lines                                               // lines are drawn seperated from each other.
	LinesWithAdjacency                                  // reserved for future use.
	LineStrips                                          // lines drawn are connected to the previous vertex.
	LineStripsWithAdjacency                             // reserved for future use.
	Triangles                                           // triangles are drawn seperated from each other.
	TrianglesWithAdjacency                              // reserved for future use.
	TriangleStrips                                      // triangles drawn are connected to the previous triangle.
	TriangleStripsWithAdjacency                         // reserved for future use.
	TriangleStripsWithRestartIndex                      // triangles drawn are connected to the previous triangle, with a restart index.
	TessellationPatch                                   // used for tessellation shaders, which can deform these patches.
)

// Rasterization options.
type Rasterization struct {
	CullMode                CullMode // controls which polygon faces are skipped.
	DepthBiasClamp          float64
	DepthBiasConstantFactor float64
	DepthBiasEnabled        bool
	DepthBiasSlopeFactor    float64
	DiscardPrimitives       bool // if true, primitives are discarded before rasterization.
	EnableDepthClamp        bool
	FrontFace               FrontFace // winding order.
	LineWidth               float64   // when drawing lines.
	PatchControlPoints      int64     // higher value, higher quality (performance cost).
	Wireframe               bool      // draw triangles as wireframe.
}

// Multisampling options.
type Multisampling struct {
	// EnableAlphaToCoverage generates a temporary coverage value based on the alpha component of the fragment's
	// first color output. This allows alpha transparency to make use of multisample antialiasing.
	EnableAlphaToCoverage bool
	// EnableAlphaToOne forces alpha to either 0.0 or 1.0. This allows hardening the edges of antialiased alpha
	// transparencies. Only relevant if EnableAlphaToCoverage is true.
	EnableAlphaToOne bool
	// EnableSampleShading enables per-sample shading which replaces MSAA by SSAA. This provides higher quality
	// antialiasing that works with transparent (alpha scissor) edges. This has a very high performance cost.
	// See also MinSampleShading. See the per-sample shading Vulkan documentation for more details.
	EnableSampleShading bool
	// MinSampleShading determines how many samples are performed for each fragment. Must be between 0.0 and 1.0
	// (inclusive). Only effective if EnableSampleShading is true. If MinSampleShading is 1.0, fragment invocation
	// must only read from the coverage index sample. Tile image access must not be used if MinSampleShading is not 1.0.
	MinSampleShading float64
	// Samples determines the number of MSAA samples (or SSAA samples if EnableSampleShading is true) to perform.
	// Higher values result in better antialiasing, at the cost of performance.
	Samples TextureSamples
	// SampleMasks array. See the sample mask Vulkan documentation for more details.
	SampleMasks []int64
}

type DepthStencils struct {
	BackOperationComparison      Comparison       // method used for compareing the previous back stencil value and BackOperationReference.
	BackOperationComparisonMask  uint64           // which bits of the stencil value are compared.
	BackOperationDepthFail       StencilOperation // operation to perform on the stencil buffer for back pixels that pass the stencil test but fail the depth test.
	BackOperationFail            StencilOperation // operation to perform on the stencil buffer for back pixels that fail the stencil test.
	BackOperationPass            StencilOperation // operation to perform on the stencil buffer for back pixels that pass the stencil test.
	BackOperationReference       int64            // reference value for the stencil test.
	BackOperationWriteMask       int64            // which bits of the stencil buffer are written to for back pixels.
	DepthComparison              Comparison       // method used for comparing the previous and current depth values.
	DepthRangeMax                float64          // maximum depth that returns true for EnableDepthRange.
	DepthRangeMin                float64          // minimum depth that returns true for EnableDepthRange.
	EnableDepthRange             bool             // if true, depth values are discarded if they fall outside DepthRangeMin and DepthRangeMax.
	EnableDepthTest              bool             // if true, enables depth testing, which occludes objects based on depth.
	EnableDepthWrite             bool             // if true, writes to the depth buffer whenever depth test passes.
	EnableStencil                bool             // if true, enables stenciling.
	FrontOperationComparison     Comparison       // method used for compareing the previous front stencil value and FrontOperationReference.
	FrontOperationComparisonMask int64            // which bits of the stencil value are compared.
	FrontOperationDepthFail      StencilOperation // operation to perform on the stencil buffer for front pixels that pass the stencil test.
	FrontOperationFail           StencilOperation // operation to perform on the stencil buffer for front pixels that fail the stencil test.
	FrontOperationPass           StencilOperation // operation to perform on the stencil buffer for front pixels that pass the stencil test.
	FrontOperationReference      int64            // reference value for the stencil test.
	FrontOperationWriteMask      int64            // which bits of the stencil buffer are written to for front pixels.
}

// ColorBlending options.
type ColorBlending struct {
	Attachments          []ColorBlendingAttachment // Attachments that are blended together.
	BlendConstant        uc.Color                  // BlendConstant to blend with.
	EnableLogicOperation bool                      // EnableLogicOperation enables LogicOperation.
	LogicOperation       LogicOperation            // LogicOperation to use for blending.
}

// ColorBlendingAttachment used for blending.
type ColorBlendingAttachment struct {
	EnableBlend bool // If true, performs blending between the source and destination according to the other fields.

	AlphaBlendOperation blendOperation // blend mode for the alpha channel.
	ColorBlendOperation blendOperation // blend mode for the color channels.

	DestinationAlphaBlendFactor BlendFactor // controls how the blend factor for the alpha channel is determined based on the destination's fragments.
	DestinationColorBlendFactor BlendFactor // controls how the blend factor for the color channels is determined based on the destination's fragments.

	SourceAlphaBlendFactor BlendFactor // controls how the blend factor for the alpha channel is determined based on the source's fragments.
	SourceColorBlendFactor BlendFactor // controls how the blend factor for the color channels is determined based on the source's fragments.

	WriteA bool // If true, writes to the alpha channel.
	WriteB bool // If true, writes to the blue channel.
	WriteG bool // If true, writes to the green channel.
	WriteR bool // If true, writes to the red channel.
}

// Mix performs standard mix blending with straight (non-premultiplied) alpha.
func Mix() ColorBlendingAttachment {
	return ColorBlendingAttachment{
		EnableBlend: true,

		AlphaBlendOperation:         Source + Destination,
		ColorBlendOperation:         Source + Destination,
		SourceColorBlendFactor:      SourceAlpha,
		DestinationColorBlendFactor: One - SourceAlpha,
		SourceAlphaBlendFactor:      SourceAlpha,
		DestinationAlphaBlendFactor: One - SourceAlpha,

		WriteA: true,
		WriteB: true,
		WriteG: true,
		WriteR: true,
	}
}

type blendOperation int

/*
BlendOperation for blending, can be one of:

	Source + Destination
	Source - Destination
	Destination - Source
	min(Source, Destination)
	max(Source, Destination)
*/
const (
	Source      blendOperation = 1
	Destination blendOperation = 3
)

// LogicOperation for blending.
type LogicOperation int

const (
	CLEAR LogicOperation = iota // 0
	AND                         // a && b
	ANDR                        // a && !b
	COPY                        // a
	ANDN                        // !a && b
	NOOP                        // b
	XOR                         // a ^ b
	OR                          // a || b
	NOR                         // !(a || b)
	XNOR                        // !(a ^ b)
	NOT                         // !b
	ORR                         // a || !b
	COPYN                       // !a
	ORN                         // !a || b
	NAND                        // !(a && b)
	SET                         // 1
)

// DynamicStates that can be enabled on the renderer.
type DynamicStates int

const (
	UsesLineWidth DynamicStates = 1 << iota
	UsesDepthBias
	UsesBlendConstants
	UsesDepthBounds
	UsesStencilComparisonMask
	UsesStencilWriteMask
	UsesStencilReference
)

// CullMode defines which polygon faces are skipped.
type CullMode int

const (
	CullDisabled CullMode = iota // no culling.
	CullFront                    // skip front faces.
	CullBack                     // skip back faces.
)

// FrontFace winding order.
type FrontFace int

const (
	FrontIsClockwise        FrontFace = iota // clockwise winding order.
	FrontIsCounterclockwise                  // counterclockwise winding order.
)

type BlendFactor int

const (
	Zero                BlendFactor = iota // 0.0
	One                                    // 1.0
	SourceColor                            // src.color
	_                                      // 1.0 - src.color
	DestinationColor                       // dst.color
	_                                      // 1.0 - dst.color
	SourceAlpha                            // src.color
	_                                      // 1.0 - src.color
	DestinationAlpha                       // dst.alpha
	_                                      // 1.0 - dst.alpha
	ConstantColor                          // const.color
	_                                      // 1.0 - const.color
	ConstantAlpha                          // const.alpha
	_                                      // 1.0 - const.alpha
	SourceAlphaSaturate                    // min(src.alpha, 1 - dst.alpha)
	SecondSourceColor                      // src2.color
	_                                      // 1.0 - src1.color
	SecondSourceAlpha                      // src2.alpha
	_                                      // 1.0 - src1.alpha
)

type Comparison int

const (
	CompareNever          Comparison = iota // false
	CompareLess                             // a < b
	CompareEqual                            // a == b
	CompareLessOrEqual                      // a <= b
	CompareGreater                          // a > b
	CompareNotEqual                         // a != b
	CompareGreaterOrEqual                   // a >= b
	CompareAlways                           // true
)

type StencilOperation int

const (
	StencilKeep              StencilOperation = iota // keep the current value.
	StencilZero                                      // set to 0.
	StencilReplace                                   // replace with the new value.
	StencilIncrementAndClamp                         // increment by 1, clamp to the maximum value.
	StencilDecrementAndClamp                         // decrement by 1, clamp to the minimum value.
	StencilInvert                                    // invert the bits.
	StencilIncrementAndWrap                          // increment by 1, wrap to 0 when the maximum value is reached.
	StencilDecrementAndWrap                          // decrement by 1, wrap to the maximum value when 0 is reached.
)
