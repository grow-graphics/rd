package rd

type DataFormat int

const (
	DataFormat_R4G4_UNORM_PACK8 DataFormat = iota
	DataFormat_R4G4B4A4_UNORM_PACK16
	DataFormat_B4G4R4A4_UNORM_PACK16
	DataFormat_R5G6B5_UNORM_PACK16
	DataFormat_B5G6R5_UNORM_PACK16
	DataFormat_R5G5B5A1_UNORM_PACK16
	DataFormat_B5G5R5A1_UNORM_PACK16
	DataFormat_A1R5G5B5_UNORM_PACK16
	DataFormat_R8_UNORM
	DataFormat_R8_SNORM
	DataFormat_R8_USCALED
	DataFormat_R8_SSCALED
	DataFormat_R8_UINT
	DataFormat_R8_SINT
	DataFormat_R8_SRGB
	DataFormat_R8G8_UNORM
	DataFormat_R8G8_SNORM
	DataFormat_R8G8_USCALED
	DataFormat_R8G8_SSCALED
	DataFormat_R8G8_UINT
	DataFormat_R8G8_SINT
	DataFormat_R8G8_SRGB
	DataFormat_R8G8B8_UNORM
	DataFormat_R8G8B8_SNORM
	DataFormat_R8G8B8_USCALED
	DataFormat_R8G8B8_SSCALED
	DataFormat_R8G8B8_UINT
	DataFormat_R8G8B8_SINT
	DataFormat_R8G8B8_SRGB
	DataFormat_B8G8R8_UNORM
	DataFormat_B8G8R8_SNORM
	DataFormat_B8G8R8_USCALED
	DataFormat_B8G8R8_SSCALED
	DataFormat_B8G8R8_UINT
	DataFormat_B8G8R8_SINT
	DataFormat_B8G8R8_SRGB
	DataFormat_R8G8B8A8_UNORM
	DataFormat_R8G8B8A8_SNORM
	DataFormat_R8G8B8A8_USCALED
	DataFormat_R8G8B8A8_SSCALED
	DataFormat_R8G8B8A8_UINT
	DataFormat_R8G8B8A8_SINT
	DataFormat_R8G8B8A8_SRGB
	DataFormat_B8G8R8A8_UNORM
	DataFormat_B8G8R8A8_SNORM
	DataFormat_B8G8R8A8_USCALED
	DataFormat_B8G8R8A8_SSCALED
	DataFormat_B8G8R8A8_UINT
	DataFormat_B8G8R8A8_SINT
	DataFormat_B8G8R8A8_SRGB
	DataFormat_A8B8G8R8_UNORM_PACK32
	DataFormat_A8B8G8R8_SNORM_PACK32
	DataFormat_A8B8G8R8_USCALED_PACK32
	DataFormat_A8B8G8R8_SSCALED_PACK32
	DataFormat_A8B8G8R8_UINT_PACK32
	DataFormat_A8B8G8R8_SINT_PACK32
	DataFormat_A8B8G8R8_SRGB_PACK32
	DataFormat_A2R10G10B10_UNORM_PACK32
	DataFormat_A2R10G10B10_SNORM_PACK32
	DataFormat_A2R10G10B10_USCALED_PACK32
	DataFormat_A2R10G10B10_SSCALED_PACK32
	DataFormat_A2R10G10B10_UINT_PACK32
	DataFormat_A2R10G10B10_SINT_PACK32
	DataFormat_A2B10G10R10_UNORM_PACK32
	DataFormat_A2B10G10R10_SNORM_PACK32
	DataFormat_A2B10G10R10_USCALED_PACK32
	DataFormat_A2B10G10R10_SSCALED_PACK32
	DataFormat_A2B10G10R10_UINT_PACK32
	DataFormat_A2B10G10R10_SINT_PACK32
	DataFormat_R16_UNORM
	DataFormat_R16_SNORM
	DataFormat_R16_USCALED
	DataFormat_R16_SSCALED
	DataFormat_R16_UINT
	DataFormat_R16_SINT
	DataFormat_R16_SFLOAT
	DataFormat_R16G16_UNORM
	DataFormat_R16G16_SNORM
	DataFormat_R16G16_USCALED
	DataFormat_R16G16_SSCALED
	DataFormat_R16G16_UINT
	DataFormat_R16G16_SINT
	DataFormat_R16G16_SFLOAT
	DataFormat_R16G16B16_UNORM
	DataFormat_R16G16B16_SNORM
	DataFormat_R16G16B16_USCALED
	DataFormat_R16G16B16_SSCALED
	DataFormat_R16G16B16_UINT
	DataFormat_R16G16B16_SINT
	DataFormat_R16G16B16_SFLOAT
	DataFormat_R16G16B16A16_UNORM
	DataFormat_R16G16B16A16_SNORM
	DataFormat_R16G16B16A16_USCALED
	DataFormat_R16G16B16A16_SSCALED
	DataFormat_R16G16B16A16_UINT
	DataFormat_R16G16B16A16_SINT
	DataFormat_R16G16B16A16_SFLOAT
	DataFormat_R32_UINT
	DataFormat_R32_SINT
	DataFormat_R32_SFLOAT
	DataFormat_R32G32_UINT
	DataFormat_R32G32_SINT
	DataFormat_R32G32_SFLOAT
	DataFormat_R32G32B32_UINT
	DataFormat_R32G32B32_SINT
	DataFormat_R32G32B32_SFLOAT
	DataFormat_R32G32B32A32_UINT
	DataFormat_R32G32B32A32_SINT
	DataFormat_R32G32B32A32_SFLOAT
	DataFormat_R64_UINT
	DataFormat_R64_SINT
	DataFormat_R64_SFLOAT
	DataFormat_R64G64_UINT
	DataFormat_R64G64_SINT
	DataFormat_R64G64_SFLOAT
	DataFormat_R64G64B64_UINT
	DataFormat_R64G64B64_SINT
	DataFormat_R64G64B64_SFLOAT
	DataFormat_R64G64B64A64_UINT
	DataFormat_R64G64B64A64_SINT
	DataFormat_R64G64B64A64_SFLOAT
	DataFormat_B10G11R11_UFLOAT_PACK32
	DataFormat_E5B9G9R9_UFLOAT_PACK32
	DataFormat_D16_UNORM
	DataFormat_X8_D24_UNORM_PACK32
	DataFormat_D32_SFLOAT
	DataFormat_S8_UINT
	DataFormat_D16_UNORM_S8_UINT
	DataFormat_D24_UNORM_S8_UINT
	DataFormat_D32_SFLOAT_S8_UINT
	DataFormat_BC1_RGB_UNORM_BLOCK
	DataFormat_BC1_RGB_SRGB_BLOCK
	DataFormat_BC1_RGBA_UNORM_BLOCK
	DataFormat_BC1_RGBA_SRGB_BLOCK
	DataFormat_BC2_UNORM_BLOCK
	DataFormat_BC2_SRGB_BLOCK
	DataFormat_BC3_UNORM_BLOCK
	DataFormat_BC3_SRGB_BLOCK
	DataFormat_BC4_UNORM_BLOCK
	DataFormat_BC4_SNORM_BLOCK
	DataFormat_BC5_UNORM_BLOCK
	DataFormat_BC5_SNORM_BLOCK
	DataFormat_BC6H_UFLOAT_BLOCK
	DataFormat_BC6H_SFLOAT_BLOCK
	DataFormat_BC7_UNORM_BLOCK
	DataFormat_BC7_SRGB_BLOCK
	DataFormat_ETC2_R8G8B8_UNORM_BLOCK
	DataFormat_ETC2_R8G8B8_SRGB_BLOCK
	DataFormat_ETC2_R8G8B8A1_UNORM_BLOCK
	DataFormat_ETC2_R8G8B8A1_SRGB_BLOCK
	DataFormat_ETC2_R8G8B8A8_UNORM_BLOCK
	DataFormat_ETC2_R8G8B8A8_SRGB_BLOCK
	DataFormat_EAC_R11_UNORM_BLOCK
	DataFormat_EAC_R11_SNORM_BLOCK
	DataFormat_EAC_R11G11_UNORM_BLOCK
	DataFormat_EAC_R11G11_SNORM_BLOCK
	DataFormat_ASTC_4x4_UNORM_BLOCK
	DataFormat_ASTC_4x4_SRGB_BLOCK
	DataFormat_ASTC_5x4_UNORM_BLOCK
	DataFormat_ASTC_5x4_SRGB_BLOCK
	DataFormat_ASTC_5x5_UNORM_BLOCK
	DataFormat_ASTC_5x5_SRGB_BLOCK
	DataFormat_ASTC_6x5_UNORM_BLOCK
	DataFormat_ASTC_6x5_SRGB_BLOCK
	DataFormat_ASTC_6x6_UNORM_BLOCK
	DataFormat_ASTC_6x6_SRGB_BLOCK
	DataFormat_ASTC_8x5_UNORM_BLOCK
	DataFormat_ASTC_8x5_SRGB_BLOCK
	DataFormat_ASTC_8x6_UNORM_BLOCK
	DataFormat_ASTC_8x6_SRGB_BLOCK
	DataFormat_ASTC_8x8_UNORM_BLOCK
	DataFormat_ASTC_8x8_SRGB_BLOCK
	DataFormat_ASTC_10x5_UNORM_BLOCK
	DataFormat_ASTC_10x5_SRGB_BLOCK
	DataFormat_ASTC_10x6_UNORM_BLOCK
	DataFormat_ASTC_10x6_SRGB_BLOCK
	DataFormat_ASTC_10x8_UNORM_BLOCK
	DataFormat_ASTC_10x8_SRGB_BLOCK
	DataFormat_ASTC_10x10_UNORM_BLOCK
	DataFormat_ASTC_10x10_SRGB_BLOCK
	DataFormat_ASTC_12x10_UNORM_BLOCK
	DataFormat_ASTC_12x10_SRGB_BLOCK
	DataFormat_ASTC_12x12_UNORM_BLOCK
	DataFormat_ASTC_12x12_SRGB_BLOCK
	DataFormat_G8B8G8R8_422_UNORM
	DataFormat_B8G8R8G8_422_UNORM
	DataFormat_G8_B8_R8_3PLANE_420_UNORM
	DataFormat_G8_B8R8_2PLANE_420_UNORM
	DataFormat_G8_B8_R8_3PLANE_422_UNORM
	DataFormat_G8_B8R8_2PLANE_422_UNORM
	DataFormat_G8_B8_R8_3PLANE_444_UNORM
	DataFormat_R10X6_UNORM_PACK16
	DataFormat_R10X6G10X6_UNORM_2PACK16
	DataFormat_R10X6G10X6B10X6A10X6_UNORM_4PACK16
	DataFormat_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16
	DataFormat_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16
	DataFormat_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16
	DataFormat_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16
	DataFormat_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16
	DataFormat_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16
	DataFormat_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16
	DataFormat_R12X4_UNORM_PACK16
	DataFormat_R12X4G12X4_UNORM_2PACK16
	DataFormat_R12X4G12X4B12X4A12X4_UNORM_4PACK16
	DataFormat_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16
	DataFormat_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16
	DataFormat_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16
	DataFormat_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16
	DataFormat_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16
	DataFormat_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16
	DataFormat_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16
	DataFormat_G16B16G16R16_422_UNORM
	DataFormat_B16G16R16G16_422_UNORM
	DataFormat_G16_B16_R16_3PLANE_420_UNORM
	DataFormat_G16_B16R16_2PLANE_420_UNORM
	DataFormat_G16_B16_R16_3PLANE_422_UNORM
	DataFormat_G16_B16R16_2PLANE_422_UNORM
	DataFormat_G16_B16_R16_3PLANE_444_UNORM
)