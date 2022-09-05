package nc

import "gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"

type UserAction struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserDetail struct {
	Username string `json:"username"`
	Password string `json:"password"`
	*types.UserDetails
}
