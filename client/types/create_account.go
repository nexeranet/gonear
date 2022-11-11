package types

type CreateAccount struct{}

func CreateAccountAction() Action{
    return Action{
        Enum: CreateAccountEnum,
        CreateAccount: CreateAccount{},
    }
}
