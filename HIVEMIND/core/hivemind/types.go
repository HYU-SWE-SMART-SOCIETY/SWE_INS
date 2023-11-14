package hivemind

type Status int8

type Instance int8

type PingRequest struct {
	userId int
}

type InstanceResponse struct {
	Ok    Status
	error error
}

type RESTSyncResponse struct {
	Result []InstanceResponse
}

type RESTPingResponse struct {
	Result []bool
}

type InstanceFrame struct {
	User    string            `json:"user"`
	Request []InstanceRequest `json:"request"`
}

type InstanceRequest struct {
	InstType int    `json:"instance"`
	Payload  string `json:"payload"`
}

type HandleFrameFunction func(string) InstanceResponse

const (
	SUCCESSFUL Status = iota
	NO_DEVICE
	ERROR
)

const (
	LIGHT_BULB Instance = iota
	CURTAIN
	AC
	REFRIGERATOR
	WATER_DISPENSER
	TELEVISION
	SOUNDBAR
	MASSAGE_CHAIR
	AI_SPEAKER
)
