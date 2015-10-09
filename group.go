package zd




type Group struct {
	Url       *string  `json:"url,omitempty"`
	Id *float64 `json:"id,omitempty"`
	
	
	Name      *string  `json:"name,omitempty"`
	Deleted   *bool    `json:"deleted,omitempty"`      
	CreatedAt *string  `json:"created_at,omitempty"`     
	UpdatedAt *string  `json:"updated_at,omitempty"`     
}
