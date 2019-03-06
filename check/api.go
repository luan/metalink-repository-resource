package main

import (
	"fmt"

	"github.com/dpb587/metalink-repository-resource/api"
	"github.com/dpb587/metalink/repository/filter/and"
	"github.com/dpb587/metalink/repository/filter/fileversion"
	"github.com/dpb587/metalink/repository/utility"
)

type Request struct {
	Source  api.Source   `json:"source"`
	Version *api.Version `json:"version"`
}

func (r Request) ApplyFilter(filter *and.Filter) error {
	err := r.Source.ApplyFilter(filter)
	if err != nil {
		return err
	} else if r.Version == nil {
		return nil
	}

	if r.Version != nil {
		addFilter, err := fileversion.CreateFilter(fmt.Sprintf("> %s", utility.RewriteSemiSemVer(r.Version.Version)))
		if err != nil {
			return err
		}

		filter.Add(addFilter)
	}

	if r.Source.MetalinkGlob != "" {
		addFilter, err := api.CreateFilePathFilter(r.Source.MetalinkGlob)
		if err != nil {
			return err
		}

		filter.Add(addFilter)
	}

	return nil
}

type Response []api.Version
