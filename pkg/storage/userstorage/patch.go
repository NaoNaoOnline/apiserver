package userstorage

import "github.com/xh3b4sd/tracer"

type Patch struct {
	// Ope is the RFC6902 compliant operation for this JSON-Patch.
	Ope string `json:"op"`
	// Pat is the RFC6902 compliant path for this JSON-Patch.
	Pat string `json:"path"`
	// Val is the RFC6902 compliant value for this JSON-Patch.
	Val string `json:"value"`
}

func (p *Patch) Verify() error {
	if p.Ope == "" {
		return tracer.Mask(jsonPatchOperationEmptyError)
	}
	if p.Ope != "replace" {
		return tracer.Maskf(jsonPatchOperationInvalidError, p.Ope)
	}

	if p.Pat == "" {
		return tracer.Maskf(jsonPatchPathEmptyError, p.Pat)
	}
	if p.Pat != "/home/data" && p.Pat != "/name/data" {
		return tracer.Maskf(jsonPatchPathInvalidError, p.Pat)
	}

	return nil
}
