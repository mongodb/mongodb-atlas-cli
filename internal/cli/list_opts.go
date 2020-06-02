package cli

import "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"

type ListOpts struct {
	PageNum      int
	ItemsPerPage int
}

func (opts *ListOpts) NewListOptions() *mongodbatlas.ListOptions {
	return &mongodbatlas.ListOptions{
		PageNum:      opts.PageNum,
		ItemsPerPage: opts.ItemsPerPage,
	}
}
