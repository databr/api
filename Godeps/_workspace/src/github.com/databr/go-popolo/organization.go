package popolo

import "time"

type Organization struct {
	Id               *string         `json:"id"`               // The organization's unique identifier
	Name             *string         `json:"name"`             // A primary name, e.g. a legally recognized name
	OtherNames       []OtherNames    `json:"other_names"`      // Alternate or former names
	Identifiers      []Identifier    `json:"identifiers"`      // Issued identifiers
	Classification   *string         `json:"classification"`   // An organization category, e.g. committee
	ParentId         *string         `json:"parent_id"`        //The ID of the organization that contains this organization
	Parent           *Organization   `json:"parent"`           // The organization that contains this organization
	AreaId           *string         `json:"area"`             // The ID of the geographic area to which this organization is related
	Area             *Area           `json:"area"`             // The geographic area to which this organization is related
	FoundingDate     *string         `json:"founding_date"`    // A date of founding
	DissoulutionDate *string         `json:"dissolution_date"` // A date of dissolution
	Image            *string         `json:"image"`            // A URL of a head shot
	ContactDetails   []ContactDetail `json:"contact_details"`  // Means of contacting the person
	Links            []Link          `json:"link"`             // URLs to documents about the person
	Memberships      []Membership    `json:"memberships"`      // Memberships
	Posts            []Post          `json:"posts"`            // Posts within the organization
	CreatedAt        time.Time       `json:"created_at"`       // The time at which the resource was created
	UpdatedAt        time.Time       `json:"updated_at"`       // The time at which the resource was last modified
	Sources          []Source        `json:"sources"`          // URLs to documents from which the person is derived
}
