package fixtures

import atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"

func Event() *atlas.Event {
	return &atlas.Event{
		Created:       "2018-06-19T15:06:15Z",
		EventTypeName: "JOINED_ORG",
		ID:            "5b48f4d2d7e33a1c0c60597e",
		IsGlobalAdmin: false,
		Links: []*atlas.Link{
			{
				Rel:  "http://mms.mongodb.com/org",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5b478b3afc4625789ce616a3",
			},
			{
				Rel:  "http://mms.mongodb.com/org",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/users/6b610e1087d9d66b272f0c86",
			},
			{
				Rel:  "http://mms.mongodb.com/org",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5b478b3afc4625789ce616a3/events/5b48f4d2d7e33a1c0c60597e",
			},
		},
		OrgID:          "5b478b3afc4625789ce616a3",
		RemoteAddress:  "198.51.100.64",
		TargetUsername: "j.doe@example.com",
		UserID:         "6b610e1087d9d66b272f0c86",
		Username:       "j.doe@example.com",
	}
}

func Events() *atlas.EventResponse {
	return &atlas.EventResponse{
		Links:      []*atlas.Link{},
		Results:    []*atlas.Event{Event()},
		TotalCount: 1,
	}
}
