package types

type DeployContract struct {
	Code []byte
}

func DeployContractAction(code []byte) Action{
    return Action{
        Enum: DeployContractEnum,
        DeployContract: DeployContract {
            Code: code,
        },
    }
}
