package atlas

// ListOptions specifies the optional parameters to List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	PageNum int `url:"pageNum,omitempty"`

	// For paginated result sets, the number of results to include per page.
	ItemsPerPage int `url:"itemsPerPage,omitempty"`

	// Flag that indicates whether Atlas returns the totalCount parameter in the response body.
	IncludeCount bool `url:"includeCount,omitempty"`
}
