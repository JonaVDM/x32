package osc

type Value struct {
	Type  byte
	Value []byte

	Float  float32
	Int    int32
	String string
}

type Message struct {
	Message string
	Values  []Value
}

const (
	MessageInfo            string = "/info"
	MessageXInfo                  = "/xinfo"
	MessageStatus                 = "/status"
	MessageSetNode                = "/"
	MessageGetNode                = "/node"
	MessageGetMeters              = "/meters"
	MessageSubscribe              = "/subscribe"
	MessageFormatSubscribe        = "/formatsubscribe"
	MessageBatchSubscribe         = "/batchsubscribe"
	MessageRenew                  = "/renew"
	MessageXRemote                = "/xremote"
)
