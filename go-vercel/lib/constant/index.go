package constant

const messageBitSize = 24
const messageBlockBitSize = 128
const keyBitSize = 128
// key bit size has to be larger than message block bit size

const MessageByteSize = messageBitSize/8
const MessageBlockByteSize = messageBlockBitSize/8
const KeyByteSize = keyBitSize/8
const Rounds = 10