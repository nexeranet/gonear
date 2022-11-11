package types

type DeleteKey struct {
	PublicKey PublicKey
}
func DeleteKeyAction(key PublicKey) Action {
    return Action {
        Enum: DeleteKeyEnum,
        DeleteKey: DeleteKey {
            PublicKey: key,
        },
    }
}
