package handler

import (
	"fmt"
	"mime/multipart"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
}

// PublishScan
type PublishScan struct {
	Product    string                `form:"product"`
	Engagement string                `form:"engagement"`
	Test       string                `form:"test"`
	Branch     string                `form:"branch"`
	RepoUrl    string                `form:"repoUrl"`
	Origin     string                `form:"origin"`
	File       *multipart.FileHeader `form:"file" binding:"required"`
}

func (r *PublishScan) Validate() error {
	if r.Product == "" {
		return errParamIsRequired("product", "string")
	}
	if r.Engagement == "" {
		return errParamIsRequired("engagement", "string")
	}
	if r.Test == "" {
		return errParamIsRequired("test", "string")
	}
	if r.Branch == "" {
		return errParamIsRequired("branch", "string")
	}
	if r.RepoUrl == "" {
		return errParamIsRequired("repoUrl", "string")
	}
	if r.Origin == "" {
		return errParamIsRequired("origin", "string")
	}
	if r.File == nil {
		return errParamIsRequired("report file", "file")
	}
	return nil
}
