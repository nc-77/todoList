package response

import "time"

type TaskResp struct {
	ID        int    `json:"id"`
	Text      string ` json:"text"`
	DDL       string `json:"ddl" `
	Finished  bool   `json:"isFinish"`
	CreatedAt time.Time `json:"createdAt"`
}

