package crypto

var (
	_encryptKey     = []byte{0x90, 0x51, 0x26, 0x25, 0x04, 0xBF, 0xCF, 0x4C, 0x92, 0x02, 0x52, 0x7A, 0x70, 0x1A, 0x41, 0x88, 0x8C, 0xC2, 0xCE, 0xB8, 0xF6, 0x57, 0x7E, 0xBA, 0x83, 0x63, 0x2C, 0x24, 0x9A, 0x67, 0x86, 0x0C, 0xBE, 0x72, 0xFD, 0xB6, 0x7B, 0x79, 0xB0, 0x22, 0x5A, 0x60, 0x5C, 0x4F, 0x49, 0xE2, 0x0E, 0xF5, 0x3A, 0x81, 0xAE, 0x11, 0x6B, 0xF0, 0xA1, 0x01, 0xE8, 0x65, 0x8D, 0x5B, 0xDC, 0xCC, 0x93, 0x18, 0xB3, 0xAB, 0x77, 0xF7, 0x8E, 0xEC, 0xEF, 0x05, 0x00, 0xCA, 0x4E, 0xA7, 0xBC, 0xB5, 0x10, 0xC6, 0x6C, 0xC0, 0xC4, 0xE5, 0x87, 0x3F, 0xC1, 0x82, 0x29, 0x96, 0x45, 0x73, 0x07, 0xCB, 0x43, 0xF9, 0xF3, 0x08, 0x89, 0xD0, 0x99, 0x6A, 0x3B, 0x37, 0x19, 0xD4, 0x40, 0xEA, 0xD7, 0x85, 0x16, 0x66, 0x1E, 0x9C, 0x39, 0xBB, 0xEE, 0x4A, 0x03, 0x8A, 0x36, 0x2D, 0x13, 0x1D, 0x56, 0x48, 0xC7, 0x0D, 0x59, 0xB2, 0x44, 0xA3, 0xFE, 0x8B, 0x32, 0x1B, 0x84, 0xA0, 0x2E, 0x62, 0x17, 0x42, 0xB9, 0x9B, 0x2B, 0x75, 0xD8, 0x1C, 0x3C, 0x4D, 0x76, 0x27, 0x6E, 0x28, 0xD3, 0x33, 0xC3, 0x21, 0xAF, 0x34, 0x23, 0xDD, 0x68, 0x9F, 0xF1, 0xAD, 0xE1, 0xB4, 0xE7, 0xA6, 0x74, 0x15, 0x4B, 0xFA, 0x3D, 0x5F, 0x7C, 0xDA, 0x2F, 0x0A, 0xE3, 0x7D, 0xC8, 0xB7, 0x12, 0x6F, 0x9E, 0xA9, 0x14, 0x53, 0x97, 0x8F, 0x64, 0xF4, 0xF8, 0xA2, 0xA4, 0x2A, 0xD2, 0x47, 0x9D, 0x71, 0xC5, 0xE9, 0x06, 0x98, 0x20, 0x54, 0x80, 0xAA, 0xF2, 0xAC, 0x50, 0xD6, 0x7F, 0xD9, 0xC9, 0xCD, 0x69, 0x46, 0x6D, 0x30, 0xB1, 0x58, 0x0B, 0x55, 0xD1, 0x5D, 0xD5, 0xBD, 0x31, 0xDE, 0xA5, 0xE4, 0x91, 0x0F, 0x61, 0x38, 0xDF, 0xA8, 0xE6, 0x3E, 0x1F, 0x35, 0xED, 0xDB, 0x94, 0xEB, 0x09, 0x5E, 0x95, 0xFB, 0xFC, 0xE0, 0x78, 0xFF}
	_decryptKey     = []byte{0x48, 0x37, 0x09, 0x76, 0x04, 0x47, 0xCC, 0x5C, 0x61, 0xF8, 0xB3, 0xE0, 0x1F, 0x7F, 0x2E, 0xEB, 0x4E, 0x33, 0xB8, 0x7A, 0xBC, 0xAB, 0x6E, 0x8C, 0x3F, 0x68, 0x0D, 0x87, 0x93, 0x7B, 0x70, 0xF2, 0xCE, 0x9D, 0x27, 0xA0, 0x1B, 0x03, 0x02, 0x97, 0x99, 0x58, 0xC5, 0x90, 0x1A, 0x79, 0x8A, 0xB2, 0xDD, 0xE6, 0x86, 0x9B, 0x9F, 0xF3, 0x78, 0x67, 0xED, 0x72, 0x30, 0x66, 0x94, 0xAE, 0xF1, 0x55, 0x6A, 0x0E, 0x8D, 0x5E, 0x82, 0x5A, 0xDB, 0xC7, 0x7D, 0x2C, 0x75, 0xAC, 0x07, 0x95, 0x4A, 0x2B, 0xD4, 0x01, 0x0A, 0xBD, 0xCF, 0xE1, 0x7C, 0x15, 0xDF, 0x80, 0x28, 0x3B, 0x2A, 0xE3, 0xF9, 0xAF, 0x29, 0xEC, 0x8B, 0x19, 0xC0, 0x39, 0x6F, 0x1D, 0xA2, 0xDA, 0x65, 0x34, 0x50, 0xDC, 0x98, 0xB9, 0x0C, 0xC9, 0x21, 0x5B, 0xAA, 0x91, 0x96, 0x42, 0xFE, 0x25, 0x0B, 0x24, 0xB0, 0xB5, 0x16, 0xD6, 0xD0, 0x31, 0x57, 0x18, 0x88, 0x6D, 0x1E, 0x54, 0x0F, 0x62, 0x77, 0x85, 0x10, 0x3A, 0x44, 0xBF, 0x00, 0xEA, 0x08, 0x3E, 0xF6, 0xFA, 0x59, 0xBE, 0xCD, 0x64, 0x1C, 0x8F, 0x71, 0xC8, 0xBA, 0xA3, 0x89, 0x36, 0xC3, 0x83, 0xC4, 0xE8, 0xA9, 0x4B, 0xEF, 0xBB, 0xD1, 0x41, 0xD3, 0xA5, 0x32, 0x9E, 0x26, 0xDE, 0x81, 0x40, 0xA7, 0x4D, 0x23, 0xB7, 0x13, 0x8E, 0x17, 0x73, 0x4C, 0xE5, 0x20, 0x05, 0x51, 0x56, 0x11, 0x9C, 0x52, 0xCA, 0x4F, 0x7E, 0xB6, 0xD8, 0x49, 0x5D, 0x3D, 0xD9, 0x12, 0x06, 0x63, 0xE2, 0xC6, 0x9A, 0x69, 0xE4, 0xD5, 0x6C, 0x92, 0xD7, 0xB1, 0xF5, 0x3C, 0xA1, 0xE7, 0xEE, 0xFD, 0xA6, 0x2D, 0xB4, 0xE9, 0x53, 0xF0, 0xA8, 0x38, 0xCB, 0x6B, 0xF7, 0x45, 0xF4, 0x74, 0x46, 0x35, 0xA4, 0xD2, 0x60, 0xC1, 0x2F, 0x14, 0x43, 0xC2, 0x5F, 0xAD, 0xFB, 0xFC, 0x22, 0x84, 0xFF}
	_sharedCryptKey = []byte{0xDD, 0xA8, 0x5F, 0x1E, 0x57, 0xAF, 0xC0, 0xCC, 0x43, 0x35, 0x8F, 0xBB, 0x6F, 0xE6, 0xA1, 0xD6, 0x60, 0xB9, 0x1A, 0xAE, 0x20, 0x49, 0x24, 0x81, 0x21, 0xFE, 0x86, 0x2B, 0x98, 0xB7, 0xB3, 0xD2, 0x91, 0x01, 0x3A, 0x4C, 0x65, 0x92, 0x1C, 0xF4, 0xBE, 0xDD, 0xD9, 0x08, 0xE6, 0x81, 0x98, 0x1B, 0x8D, 0x60, 0xF3, 0x6F, 0xA1, 0x47, 0x24, 0xF1, 0x53, 0x45, 0xC8, 0x7B, 0x88, 0x80, 0x4E, 0x36, 0xC3, 0x0D, 0xC9, 0xD6, 0x8B, 0x08, 0x19, 0x0B, 0xA5, 0xC1, 0x11, 0x4C, 0x60, 0xF8, 0x5D, 0xFC, 0x15, 0x68, 0x7E, 0x32, 0xC0, 0x50, 0xAB, 0x64, 0x1F, 0x8A, 0xD4, 0x08, 0x39, 0x7F, 0xC2, 0xFB, 0xBA, 0x6C, 0xF0, 0xE6, 0xB0, 0x31, 0x10, 0xC1, 0xBF, 0x75, 0x43, 0xBB, 0x18, 0x04, 0x0D, 0xD1, 0x97, 0xF7, 0x23, 0x21, 0x83, 0x8B, 0xCA, 0x25, 0x2B, 0xA3, 0x03, 0x13, 0xEA, 0xAE, 0xFE, 0xF0, 0xEB, 0xFD, 0x85, 0x57, 0x53, 0x65, 0x41, 0x2A, 0x40, 0x99, 0xC0, 0x94, 0x65, 0x7E, 0x7C, 0x93, 0x82, 0xB0, 0xB3, 0xE5, 0xC0, 0x21, 0x09, 0x84, 0xD5, 0xEF, 0x9F, 0xD1, 0x7E, 0xDC, 0x4D, 0xF5, 0x7E, 0xCD, 0x45, 0x3C, 0x7F, 0xF5, 0x59, 0x98, 0xC6, 0x55, 0xFC, 0x9F, 0xA3, 0xB7, 0x74, 0xEE, 0x31, 0x98, 0xE6, 0xB7, 0xBE, 0x26, 0xF4, 0x3C, 0x76, 0xF1, 0x23, 0x7E, 0x02, 0x4E, 0x3C, 0xD1, 0xC7, 0x28, 0x23, 0x73, 0xC4, 0xD9, 0x5E, 0x0D, 0xA1, 0x80, 0xA5, 0xAA, 0x26, 0x0A, 0xA3, 0x44, 0x82, 0x74, 0xE6, 0x3C, 0x44, 0x27, 0x51, 0x0D, 0x5F, 0xC7, 0x9C, 0xD6, 0x63, 0x67, 0xA5, 0x27, 0x97, 0x38, 0xFB, 0x2D, 0xD3, 0xD6, 0x60, 0x25, 0x83, 0x4D, 0x37, 0x5B, 0x40, 0x59, 0x11, 0x77, 0x51, 0x11, 0x14, 0x18, 0x07, 0x63, 0xB1, 0x34, 0x3D, 0xB8, 0x60, 0x13, 0xC2, 0xE8, 0x13, 0x82}
)

// Encrypt encrypts the given data using MHF's custom encryption+checksum method.
// if a overrideByteKey value is supplied (!= nil), it will be used to override the derived/truncated key byte.
func Encrypt(data []byte, key uint32, overrideByteKey *byte) (outputData []byte, combinedCheck uint16, check0 uint16, check1 uint16, check2 uint16) {
	return _generalCrypt(data, key, 0, overrideByteKey)
}

// Decrypt decrypts the given data using MHF's custom decryption+checksum method.
// if a overrideByteKey value is supplied (!= nil), it will be used to override the derived/truncated key byte.
func Decrypt(data []byte, key uint32, overrideByteKey *byte) (outputData []byte, combinedCheck uint16, check0 uint16, check1 uint16, check2 uint16) {
	return _generalCrypt(data, key, 1, overrideByteKey)
}

// _generalCrypt is a generalized MHF crypto function that can perform both encryption and decryption,
// these two crypto operations are combined into a single function because they shared most of their logic.
//		encrypt: cryptType==0
//		decrypt: cryptType==1
func _generalCrypt(data []byte, rotKey uint32, cryptType int, overrideByteKey *byte) ([]byte, uint16, uint16, uint16, uint16) {
	cryptKeyTruncByte := byte(((rotKey >> 1) % 999983) & 0xFF)
	if overrideByteKey != nil {
		cryptKeyTruncByte = *overrideByteKey
	}

	derivedCryptKey := int32((uint32(len(data)) * uint32(cryptKeyTruncByte+1)) & 0xFFFFFFFF)
	sharedBufIdx := byte(1)
	accumulator0 := uint32(0)
	accumulator1 := uint32(0)
	accumulator2 := uint32(0)

	var outputData []byte
	if cryptType == 0 {
		for i := 0; i < len(data); i++ {
			// Do the encryption for this iteration
			encKeyIdx := int32(((uint32(derivedCryptKey) >> 10) ^ uint32(data[i])) & 0xFF)
			derivedCryptKey = (0x4FD * (derivedCryptKey + 1))
			encKeyByte := _encryptKey[encKeyIdx]

			// Update the checksum accumulators.
			accumulator2 = uint32((accumulator2 + (uint32(sharedBufIdx) * uint32(data[i]))) & 0xFFFFFFFF)
			accumulator1 = uint32((accumulator1 + uint32(encKeyIdx)) & 0xFFFFFFFF)
			accumulator0 = uint32((accumulator0 + (uint32(encKeyByte)<<(i&7))&0xFFFFFFFF) & 0xFFFFFFFF)

			// Append the output.
			outputData = append(outputData, _sharedCryptKey[sharedBufIdx]^encKeyByte)

			// Update the sharedBufIdx for the next iteration.
			sharedBufIdx = data[i]
		}

	} else if cryptType == 1 {
		for i := 0; i < len(data); i++ {
			// Do the decryption for this iteration
			oldSharedBufIdx := sharedBufIdx
			tIdx := data[i] ^ _sharedCryptKey[sharedBufIdx]
			decKeyByte := _decryptKey[tIdx]
			sharedBufIdx = byte(((uint32(derivedCryptKey) >> 10) ^ uint32(decKeyByte)) & 0xFF)

			// Update the checksum accumulators.
			accumulator0 = (accumulator0 + ((uint32(tIdx) << (i & 7)) & 0xFFFFFFFF))
			accumulator1 = (accumulator1 + uint32(decKeyByte)) & 0xFFFFFFFF
			accumulator2 = (accumulator2 + ((uint32(oldSharedBufIdx) * uint32(sharedBufIdx)) & 0xFFFFFFFF)) & 0xFFFFFFFF

			// Append the output.
			outputData = append(outputData, sharedBufIdx)

			// Update the key pos for next iteration.
			derivedCryptKey = (0x4FD * (derivedCryptKey + 1))
		}
	}

	combinedCheck := uint16((accumulator1 + (accumulator0 >> 1) + (accumulator2 >> 2)) & 0xFFFF)
	check0 := uint16((accumulator0 ^ ((accumulator0 & 0xFFFF0000) >> 16)) & 0xFFFF)
	check1 := uint16((accumulator1 ^ ((accumulator1 & 0xFFFF0000) >> 16)) & 0xFFFF)
	check2 := uint16((accumulator2 ^ ((accumulator2 & 0xFFFF0000) >> 16)) & 0xFFFF)

	return outputData, combinedCheck, check0, check1, check2
}
