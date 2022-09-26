package types

type SignedTransaction[T Action] struct {
	Transaction Transaction[T]
	Signature   Signature
}

type Action interface {
	FunctionCallAction | TransferAction
}

type Actions[T Action] []T

type Transaction[T Action] struct {
	SignerId   string
	PublicKey  PublicKey
	Nonce      uint64
	ReceiverId string
	BlockHash  [32]byte
	Actions    Actions[T]
}
