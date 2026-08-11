package models

type GetActionGroupPart struct {
	RowID     string `json:"row_id"`
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timestamp"`
}

type GetActionGroupCursor struct {
	Timestamp int64  `json:"timestamp"`
	RowID     string `json:"row_id"`
}

type GetActionsGroupOutput struct {
	Actions []GetActionGroupPart  `json:"actions"`
	Cursor  *GetActionGroupCursor `json:"cursor"`
}

type GetActionGroupInput struct {
	Timestamp *int64  `json:"timestamp"`
	RowID     *string `json:"row_id"`
}
