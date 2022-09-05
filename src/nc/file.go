package nc

type FileParams struct {
	Type      string `json:"type"`
	OldPath   string `json:"old_path"`
	NewPath   string `json:"new_path"`
	Path      string `json:"path"`
	Filename  string `json:"filename"`
	Overwrite bool   `json:"overwrite"`
}
