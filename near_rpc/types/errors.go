package near_rpc_types

import (
	"fmt"
	"strings"

	"github.com/nexeranet/gonear/jsonrpc"
)

type ErrorType string

func (e ErrorType) Type() string {
	return string(e)
}

const (
	UnknowErrorType        ErrorType = "UNKNOW_ERROR" // INFO: package error type
	HandlerError           ErrorType = "HANDLER_ERROR"
	RequestValidationError ErrorType = "REQUEST_VALIDATION_ERROR"
	InternalError          ErrorType = "INTERNAL_ERROR"
)

func ConvertErrorType(name string) ErrorType {
	switch ErrorType(name) {
	case HandlerError:
		return HandlerError
	case RequestValidationError:
		return RequestValidationError
	case InternalError:
		return InternalError
	default:
		return UnknowErrorType
	}
}

type ErrorCause string

func (e ErrorCause) Name() string {
	return string(e)
}

const (
	UnknownCause               ErrorCause = "UNKNOWN_CAUSE" // INFO: package error cause
	UnknownBlockCause          ErrorCause = "UNKNOWN_BLOCK"
	InvalidAcountCause         ErrorCause = "INVALID_ACCOUNT"
	UnknownAccountCause        ErrorCause = "UNKNOWN_ACCOUNT"
	UnavailableShardCause      ErrorCause = "UNAVAILABLE_SHARD"
	NoSyncedBlocksCause        ErrorCause = "NO_SYNCED_BLOCKS"
	ParseErrorCause            ErrorCause = "PARSE_ERROR"
	GarbageCollectedBlockCause ErrorCause = "GARBAGE_COLLECTED_BLOCK"
	InternalErrorCause         ErrorCause = "INTERNAL_ERROR"
	NotSyncedYetCause          ErrorCause = "NOT_SYNCED_YET"
	UnknownEpochCause          ErrorCause = "UNKNOWN_EPOCH"
	UnknownReceiptCause        ErrorCause = "UNKNOWN_RECEIPT"
	UnknownTransactionCause    ErrorCause = "UNKNOWN_TRANSACTION"
	InvalidTransactionCause    ErrorCause = "INVALID_TRANSACTION"
	TimeoutErrorCause          ErrorCause = "TIMEOUT_ERROR"
	UnknownChunkCause          ErrorCause = "UNKNOWN_CHUNK"
	InvalidShardIdCause        ErrorCause = "INVALID_SHARD_ID"
)

func ConvertError(err *jsonrpc.RPCError) error {
	errorType := ConvertErrorType(err.Name)
	switch ErrorCause(err.Cause.Name) {
    case InvalidShardIdCause:
        return &ErrorInvalidShardId{
			NearError: NewNearError(errorType, InvalidShardIdCause, err),
        }
	case UnknownChunkCause:
		return &ErrorUnknownChunk{
			NearError: NewNearError(errorType, UnknownChunkCause, err),
		}
	case TimeoutErrorCause:
		return &ErrorTimeoutError{
			NearError: NewNearError(errorType, TimeoutErrorCause, err),
		}
	case InvalidTransactionCause:
		return &ErrorInvalidTransaction{
			NearError: NewNearError(errorType, InvalidTransactionCause, err),
		}
	case UnknownTransactionCause:
		return &ErrorUnknownTransaction{
			NearError: NewNearError(errorType, UnknownTransactionCause, err),
		}
	case UnknownBlockCause:
		return &ErrorUnknownBlock{
			NearError: NewNearError(errorType, UnknownBlockCause, err),
		}
	case UnknownReceiptCause:
		return &ErrorUnknownReceipt{
			NearError: NewNearError(errorType, UnknownReceiptCause, err),
		}
	case InvalidAcountCause:
		return &ErrorInvalidAccount{
			NearError: NewNearError(errorType, InvalidAcountCause, err),
		}
	case UnknownAccountCause:
		return &ErrorUnknownAccount{
			NearError: NewNearError(errorType, UnknownAccountCause, err),
		}
	case UnavailableShardCause:
		return &ErrorUnavailableShard{
			NearError: NewNearError(errorType, UnavailableShardCause, err),
		}
	case NoSyncedBlocksCause:
		return &ErrorNoSyncedBlocks{
			NearError: NewNearError(errorType, NoSyncedBlocksCause, err),
		}
	case ParseErrorCause:
		return &ErrorParseError{
			NearError: NewNearError(errorType, ParseErrorCause, err),
		}
	case GarbageCollectedBlockCause:
		return &ErrorGarbageCollectedBlock{
			NearError: NewNearError(errorType, GarbageCollectedBlockCause, err),
		}
	case InternalErrorCause:
		return &ErrorInternalError{
			NewNearError(errorType, InternalErrorCause, err),
		}
	case NotSyncedYetCause:
		return &ErrorNotSyncedYet{
			NewNearError(errorType, NotSyncedYetCause, err),
		}
	case UnknownEpochCause:
		return &ErrorUnknownEpoch{
			NewNearError(errorType, UnknownEpochCause, err),
		}
	default:
		return &ErrorUnknownCause{
			NewNearError(errorType, UnknownCause, err),
		}
	}
}
func NewNearError(errorType ErrorType, cause ErrorCause, rpcErr *jsonrpc.RPCError) *NearError {
	return &NearError{
		ErrorType:  errorType,
		ErrorCause: cause,
		RPCError:   rpcErr,
	}
}

type NearError struct {
	ErrorType
	ErrorCause
	RPCError *jsonrpc.RPCError
}

func (e *NearError) Cause() ErrorCause {
	return e.ErrorCause
}

func (e *NearError) Info() map[string]interface{} {
	return e.RPCError.Cause.Info
}

func (e *NearError) Error() string {
	var info []string
	for key, value := range e.Info() {
		info = append(info, fmt.Sprintf("%s: %v", key, value))
	}
	return fmt.Sprintf("%s:%s, info [%v]", e.Type(), e.Name(), strings.Join(info, ", "))
}

type ErrorUnknownCause struct {
	*NearError
}

type ErrorUnknownBlock struct {
	*NearError
}

type ErrorInvalidAccount struct {
	*NearError
}

type ErrorUnknownAccount struct {
	*NearError
}

type ErrorUnavailableShard struct {
	*NearError
}

type ErrorNoSyncedBlocks struct {
	*NearError
}

type ErrorParseError struct {
	*NearError
}

type ErrorGarbageCollectedBlock struct {
	*NearError
}

type ErrorInternalError struct {
	*NearError
}

type ErrorNotSyncedYet struct {
	*NearError
}

type ErrorUnknownEpoch struct {
	*NearError
}
type ErrorUnknownReceipt struct {
	*NearError
}

type ErrorUnknownTransaction struct {
	*NearError
}

type ErrorInvalidTransaction struct {
	*NearError
}

type ErrorTimeoutError struct {
	*NearError
}
type ErrorUnknownChunk struct {
	*NearError
}

type ErrorInvalidShardId struct {
	*NearError
}
