// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package fixtures

import (
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

var Project1 = &atlas.Project{
	ClusterCount: 2,
	Created:      "2016-07-14T14:19:33Z",
	ID:           "5a0a1e7e0f2912c554080ae6",
	Links: []*atlas.Link{
		{
			Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
			Rel:  "self",
		},
	},
	Name:  "ProjectBar",
	OrgID: "5a0a1e7e0f2912c554080adc",
}

var Project2 = &atlas.Project{
	ClusterCount: 0,
	Created:      "2017-10-16T15:24:01Z",
	ID:           "5a0a1e7e0f2912c554080ae7",
	Links: []*atlas.Link{
		{
			Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae7",
			Rel:  "self",
		},
	},
	Name:  "Project Foo",
	OrgID: "5a0a1e7e0f2912c554080adc",
}

func Projects() *atlas.Projects {
	return &atlas.Projects{
		Links: []*atlas.Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups",
				Rel:  "self",
			},
		},
		Results:    []*atlas.Project{Project1, Project2},
		TotalCount: 2,
	}
}
