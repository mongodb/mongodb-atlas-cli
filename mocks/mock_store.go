package mocks

import atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"

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

func ProjectsMock() *atlas.Projects {
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
