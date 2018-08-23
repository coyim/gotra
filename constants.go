package gotra

const (
	version = uint16(4)
)

var otrPrefix = []byte("?OTR:")

const (
	messageTypeDataMessage     = uint8(0x03)
	messageTypeIdentityMessage = uint8(0x35)
	messageTypeAuthRMessage    = uint8(0x36)
	messageTypeAuthIMessage    = uint8(0x37)
)

const (
	usageFingerprint                  = byte(0x00)
	usageThirdBraceKey                = byte(0x01)
	usageBraceKey                     = byte(0x02)
	usageSharedSecret                 = byte(0x03)
	usageSSID                         = byte(0x04)
	usageAuthRBobClientProfile        = byte(0x05)
	usageAuthRAliceClientProfile      = byte(0x06)
	usageAuthRPhi                     = byte(0x07)
	usageAuthIBobClientProfile        = byte(0x08)
	usageAuthIAliceClientProfile      = byte(0x09)
	usageAuthIPhi                     = byte(0x0A)
	usageTmpKey                       = byte(0x0B)
	usageAuthMACKey                   = byte(0x0C)
	usageNonIntAuthBobClientProfile   = byte(0x0D)
	usageNonIntAuthAliceClientProfile = byte(0x0E)
	usageNonIntAuthPhi                = byte(0x0F)
	usageAuthMAC                      = byte(0x10)
	usageECDHFirstEphemeral           = byte(0x11)
	usageDHFirstEphemeral             = byte(0x12)
	usageRootKey                      = byte(0x13)
	usageChainKey                     = byte(0x14)
	usageNextChainKey                 = byte(0x15)
	usageMessageKey                   = byte(0x16)
	usageMACKey                       = byte(0x17)
	usageExtraSymmKey                 = byte(0x18)
	usageDataMessageSections          = byte(0x19)
	usageAuthenticator                = byte(0x1A)
	usageSMPSecret                    = byte(0x1B)
	usageAuth                         = byte(0x1C)
)
