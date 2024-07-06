package rd

// Shader reference.
type Shader interface {
	Resource
	Nameable

	// Compile the shader instance from a binary compiled shader.
	Compile([]byte)

	// VertexInputAttributeMask returns the internal vertex input mask. Internally, the vertex input mask
	// is an unsigned integer consisting of the locations (specified in GLSL via. layout(location = ...))
	// of the input variables (specified in GLSL by the in keyword).
	VertexInputAttributeMask() uint32

	// Variables creates a set of variables. Each variable has a binding address, which is where the variable
	// value will be bound to inside the shader.
	Variables(variables map[int]Variable) Variables
}

// ShaderSource representation.
type ShaderSource struct {
	Language              ShaderLanguage
	Compute               []byte
	Fragment              []byte
	TesselationControl    []byte
	TesselationEvaluation []byte
	Vertex                []byte
}

// ShaderLanguage types.
type ShaderLanguage int

const (
	ShaderLanguageGLSL ShaderLanguage = iota
	ShaderLanguageHLSL
)

// SPIRV source representation.
type SPIRV interface {
	Compute() []byte               // Compute SPIR-V.
	Fragment() []byte              // Fragment SPIR-V.
	TesselationControl() []byte    // TesselationControl SPIR-V.
	TesselationEvaluation() []byte // TesselationEvaluation SPIR-V.
	Vertex() []byte                // Vertex SPIR-V.

	// Shader creates a new shader instance from SPIR-V source.
	Shader(name string) Shader
}

/*
VariableLevel identifies a group of variables with a shared lifetime, typically, up to four levels
of variables are supported to be bound to a rendering operation. They should be organised based on
how often they change.
*/
type VariableLevel int

const (
	VariablesForFrame VariableLevel = iota
	VariablesForShader
	VariablesForMaterial
	VariablesForInstance
)

// Variables are global CPU-mutable values for shaders.
type Variables interface {
	Resource

	// AreValid returns true if the variables are valid, false otherwise.
	AreValid() bool
}

// Variable can be used to set shader variables.
type Variable interface {
	variable()
}

// SamplerWithTexture can be used as a variable.
type SamplerWithTexture struct {
	Variable

	Sampler Sampler
	Texture Texture
}

// SamplerWithTextureBuffer can be used as a variables.
type SamplerWithTextureBuffer struct {
	Variable

	Sampler       Sampler
	TextureBuffer TextureBuffer
}

// InputAttachment defined by a [FramebufferPass].
type InputAttachment struct {
	Variable

	Texture Texture
}
