package nc

import "gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"

type ShareParams struct {
	ShareId      int                   `json:"shareId"`
	Path         string                `json:"path"`
	ShareType    types.ShareType       `json:"shareType"`    //(int) 0 = user; 1 = group; 3 = public link; 4 = email; 6 = federated cloud share; 7 = circle; 10 = Talk conversation
	ShareWith    string                `json:"shareWith"`    //(string) user / group id / email address / circleID / conversation name with which the file should be shared
	PublicUpload bool                  `json:"publicUpload"` //(string) allow public upload to a public shared folder (true/false)
	Password     string                `json:"password"`     // (string) password to protect public link Share with
	Permissions  types.SharePermission `json:"permissions"`  // (int) 1 = read; 2 = update; 4 = create; 8 = delete; 16 = share; 31 = all (default: 31, for public shares: 1)
	ExpireDate   string                `json:"expireDate"`   // (string) set a expire date for public link shares. This argument expects a well formatted date string, e.g. ‘YYYY-MM-DD’
	Note         string                `json:"note"`         //(string) Adds a note for the share recipient.
}
