package types
type DeleteAccount struct {
	BeneficiaryID string
}

func DeleteAccountAction(id string) Action {
    return Action {
        Enum: DeleteAccountEnum,
        DeleteAccount: DeleteAccount{
            BeneficiaryID: id,
        },
    }
}
