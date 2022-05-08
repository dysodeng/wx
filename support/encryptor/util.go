package encryptor

import "bytes"

// PadLength calculates padding length, from github.com/vgorin/cryptogo
func PadLength(sliceLength, blockSize int) (padLen int) {
	padLen = blockSize - sliceLength%blockSize
	if padLen == 0 {
		padLen = blockSize
	}
	return padLen
}

func PKCS7Pad(message []byte, blockSize int) (padded []byte) {
	// block size must be bigger or equal 2
	if blockSize < 1<<1 {
		panic("block size is too small (minimum is 2 bytes)")
	}
	// block size up to 255 requires 1 byte padding
	if blockSize < 1<<8 {
		// calculate padding length
		padLen := PadLength(len(message), blockSize)

		// define PKCS7 padding block
		padding := bytes.Repeat([]byte{byte(padLen)}, padLen)

		// apply padding
		padded = append(message, padding...)
		return padded
	}
	// block size bigger or equal 256 is not currently supported
	panic("unsupported block size")
}

// value2CDATA 值转换为CDATA
func value2CDATA(value string) CDATAText {
	return CDATAText{"<![CDATA[" + value + "]]>"}
}
