package near_api_types

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
)

func ConvertError(err *jsonrpc.RPCError) error {
	errorType := ConvertErrorType(err.Name)
	switch ErrorCause(err.Cause.Name) {
	case UnknownBlockCause:
		return &ErrorUnknownBlock{
			ErrorType:  errorType,
			ErrorCause: UnknownBlockCause,
		}
	case InvalidAcountCause:
		return &ErrorInvalidAccount{
			ErrorType:          errorType,
			ErrorCause:         InvalidAcountCause,
		}
	case UnknownAccountCause:
		return &ErrorUnknownAccount{
			ErrorType:  errorType,
			ErrorCause: UnknownAccountCause,
		}
	case UnavailableShardCause:
		return &ErrorUnavailableShard{
			ErrorType:  errorType,
			ErrorCause: UnavailableShardCause,
		}
	case NoSyncedBlocksCause:
		return &ErrorNoSyncedBlocks{
			ErrorType:  errorType,
			ErrorCause: NoSyncedBlocksCause,
		}
	case ParseErrorCause:
		return &ErrorParseError{
			ErrorType:  errorType,
			ErrorCause: ParseErrorCause,
		}
	case GarbageCollectedBlockCause:
		return &ErrorGarbageCollectedBlock{
			ErrorType:  errorType,
			ErrorCause: GarbageCollectedBlockCause,
		}
	case InternalErrorCause:
		return &ErrorInternalError{
			ErrorType:  errorType,
			ErrorCause: InternalErrorCause,
		}
	default:
		return &ErrorUnknownCause{
			ErrorType:  errorType,
			ErrorCause: UnknownCause,
		}
	}
}

type ErrorUnknownCause struct {
	ErrorType
	ErrorCause
    RPCError *jsonrpc.RPCError
}

func (e *ErrorUnknownCause) Error() string {
    var info []string
    // for key, value := range e.Info {
    //     info = append(info, fmt.Sprintf("%s: %v", key, value))
    // }
	return fmt.Sprintf("%s:%s, info [%v]", e.Type(), e.Name(), strings.Join(info, ", "))
}

type ErrorUnknownBlock struct {
	ErrorType
	ErrorCause
}

func (e *ErrorUnknownBlock) Error() string {
	return fmt.Sprintf(
		"%s:%s",
		e.Type(),
		e.Name())
}

type ErrorInvalidAccount struct {
	ErrorType
	ErrorCause
}

func (e *ErrorInvalidAccount) Error() string {
	return fmt.Sprintf(
		"%s:%s,",
		e.Type(),
		e.Name())
}

type ErrorUnknownAccount struct {
	ErrorType
	ErrorCause
}

func (e *ErrorUnknownAccount) Error() string {
	return fmt.Sprintf("%s:%s", e.Type(), e.Name())
}

type ErrorUnavailableShard struct {
	ErrorType
	ErrorCause
}

func (e *ErrorUnavailableShard) Error() string {
	return fmt.Sprintf("%s:%s", e.Type(), e.Name())
}

type ErrorNoSyncedBlocks struct {
	ErrorType
	ErrorCause
}

func (e *ErrorNoSyncedBlocks) Error() string {
	return fmt.Sprintf("%s:%s", e.Type(), e.Name())
}

type ErrorParseError struct {
	ErrorType
	ErrorCause
}

func (e *ErrorParseError) Error() string {
	return fmt.Sprintf("%s:%s", e.Type(), e.Name())
}

type ErrorGarbageCollectedBlock struct {
	ErrorType
	ErrorCause
}

func (e *ErrorGarbageCollectedBlock) Error() string {
	return fmt.Sprintf("%s:%s", e.Type(), e.Name())
}

type ErrorInternalError struct {
	ErrorType
	ErrorCause
}

func (e *ErrorInternalError) Error() string {
	return fmt.Sprintf("%s:%s", e.Type(), e.Name())
}
