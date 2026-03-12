package registry

import (
	"encoding/json"
	"io"
)

// ApplyPatchBody 用于在 Gin/其他 HTTP 框架里接收 registry 推送的 patch
func ApplyPatchBody(body []byte) error {
	var p patch
	if err := json.Unmarshal(body, &p); err != nil {
		return err
	}
	return prov.Update(p)
}

// ApplyPatchReader 备用：如果以后想直接从 io.Reader 读 patch
func ApplyPatchReader(r io.Reader) error {
	var p patch
	if err := json.NewDecoder(r).Decode(&p); err != nil {
		return err
	}
	return prov.Update(p)
}
