package cache

import "github.com/pierrec/lz4/v4"

// CompressData this method for compress data with LZ4:
func CompressData(data []byte) ([]byte, error) {
	compressedSize := lz4.CompressBlockBound(len(data))
	compressedData := make([]byte, compressedSize)
	compressedSize, err := lz4.CompressBlock(data, compressedData, nil)
	if err != nil {
		return nil, err
	}

	return compressedData[:compressedSize], nil
}

// DecompressData this method for decompress data from LZ4:
func DecompressData(compressedData []byte, originalSize int) ([]byte, error) {
	decompressedData := make([]byte, originalSize*10)
	_, err := lz4.UncompressBlock(compressedData, decompressedData)
	if err != nil {
		return nil, err
	}

	return decompressedData, nil
}
