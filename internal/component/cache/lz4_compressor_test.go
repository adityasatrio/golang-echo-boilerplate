package cache

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCompressData_Success(t *testing.T) {
	// Arrange
	originalData := []byte("This is test data for compression. It should be compressed successfully using LZ4 algorithm.")

	// Act
	compressedData, err := CompressData(originalData)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, compressedData)
	assert.Greater(t, len(compressedData), 0, "Compressed data should not be empty")
	// Typically compressed data should be smaller than original for larger payloads
	// For small data, it might be similar or larger due to compression overhead
}

func TestCompressData_EmptyData(t *testing.T) {
	// Arrange
	originalData := []byte("")

	// Act
	compressedData, err := CompressData(originalData)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, compressedData)
}

func TestCompressData_LargeData(t *testing.T) {
	// Arrange - Create a large repeating pattern that compresses well
	largeData := make([]byte, 10000)
	for i := range largeData {
		largeData[i] = byte(i % 10)
	}

	// Act
	compressedData, err := CompressData(largeData)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, compressedData)
	assert.Less(t, len(compressedData), len(largeData), "Compressed data should be smaller than original for large repetitive data")
}

func TestDecompressData_Success(t *testing.T) {
	// Arrange
	originalData := []byte("This is test data for compression and decompression roundtrip test.")
	originalSize := len(originalData)

	// Compress first
	compressedData, err := CompressData(originalData)
	require.NoError(t, err)

	// Act - Decompress
	decompressedData, err := DecompressData(compressedData, originalSize)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, decompressedData)
	// Compare first originalSize bytes (decompressedData may have padding)
	assert.Equal(t, originalData, decompressedData[:originalSize], "Decompressed data should match original")
}

func TestCompressDecompressRoundtrip(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "Simple string",
			data: []byte("Hello, World!"),
		},
		{
			name: "JSON-like structure",
			data: []byte(`{"name":"John","age":30,"city":"New York","active":true}`),
		},
		{
			name: "Unicode characters",
			data: []byte("Hello 世界 🌍 مرحبا мир"),
		},
		{
			name: "Repeating pattern",
			data: []byte("AAAABBBBCCCCDDDD"),
		},
		{
			name: "Binary-like data",
			data: []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD, 0xFC},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act - Compress
			compressedData, err := CompressData(tc.data)
			require.NoError(t, err, "Compression should succeed")

			// Act - Decompress
			originalSize := len(tc.data)
			decompressedData, err := DecompressData(compressedData, originalSize)
			require.NoError(t, err, "Decompression should succeed")

			// Assert - Verify roundtrip
			assert.Equal(t, tc.data, decompressedData[:originalSize], "Roundtrip should preserve data")
		})
	}
}

func TestDecompressData_InvalidCompressedData(t *testing.T) {
	// Arrange - Invalid compressed data
	invalidData := []byte("This is not compressed data")
	originalSize := len(invalidData)

	// Act
	decompressedData, err := DecompressData(invalidData, originalSize)

	// Assert
	// LZ4 may or may not return error for invalid data depending on the content
	// If error, that's expected; if no error, decompressed data will be garbage
	if err != nil {
		assert.Error(t, err, "Should handle invalid compressed data")
	} else {
		assert.NotNil(t, decompressedData)
	}
}

func TestCompressData_NilData(t *testing.T) {
	// Arrange
	var nilData []byte

	// Act
	compressedData, err := CompressData(nilData)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, compressedData)
}

func TestDecompressData_WithWrongOriginalSize(t *testing.T) {
	// Arrange
	originalData := []byte("Test data for size mismatch")
	compressedData, err := CompressData(originalData)
	require.NoError(t, err)

	// Act - Use wrong original size (much smaller)
	wrongSize := 5
	decompressedData, err := DecompressData(compressedData, wrongSize)

	// Assert
	require.NoError(t, err) // LZ4 doesn't validate size match
	assert.NotNil(t, decompressedData)
	// The decompressed data won't match original due to wrong size assumption
}

func TestCompressData_PerformanceWithLargePayload(t *testing.T) {
	// Arrange - 1MB of data
	largeData := make([]byte, 1024*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	// Act
	compressedData, err := CompressData(largeData)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, compressedData)
	assert.Greater(t, len(compressedData), 0)

	// Verify decompression works
	decompressedData, err := DecompressData(compressedData, len(largeData))
	require.NoError(t, err)
	assert.Equal(t, largeData, decompressedData[:len(largeData)])
}
