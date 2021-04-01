package nilbot

type Event struct {
	Time     int64 `json:"time"`
	Duration int64 `json:"duration"`

	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	HonorType   string `json:"honor_type"`
	RequestType string `json:"request_type"`
	NoticeType  string `json:"notice_type"`

	MessageID  int64 `json:"message_id"`
	GroupID    int64 `json:"group_id"`
	UserID     int64 `json:"user_id"`
	TargetID   int64 `json:"target_id"`
	SelfID     int64 `json:"self_id"`
	OperatorID int64 `json:"operator_id"`

	Anonymous *Anonymous `json:"anonymous"`
	File      *File      `json:"file"`
	Sender    *Sender    `json:"sender"`
	Client    *[]Device  `json:"client"`

	Flag       string `json:"flag"`
	Comment    string `json:"comment"`
	Message    string `json:"message"`
	RawMessage string `json:"raw_message"`

	Online bool `json:"online"`
}
type Sender struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Card     string `json:"card"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	Area     string `json:"area"`
	Level    string `json:"string"`
	Role     string `json:"role"`
	Title    string `json:"title"`
}
type Anonymous struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}
type File struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	BusID int64  `json:"busid"`
}
type Device struct {
	APPID int64  `json:"app_id"`
	Name  string `json:"device_name"`
	Kind  string `json:"device_kind"`
}
