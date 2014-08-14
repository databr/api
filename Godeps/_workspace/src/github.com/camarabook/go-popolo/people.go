package popolo

import "time"

const DateSpec = "2006-01-02"

// Using spec from http://popoloproject.com/

type Person struct {
	Id               *string         `json:"id"`                // The person's unique identifier
	Name             *string         `json:"name"`              // A person's preferred full name
	OtherNames       []OtherNames    `json:"other_names"`       // Alternate or former names
	Identifiers      []Identifier    `json:"identifiers"`       // Issued identifiers
	FamilyName       *string         `json:"family_name"`       // One or more family names
	GivenName        *string         `json:"given_name"`        // One or more primary given names
	AdditionalName   *string         `json:"additional_name"`   // One or more secondary given names
	HonorificPrefix  *string         `json:"honorific_prefix"`  // One or more honorifics preceding a person's name
	HonorificSuffix  *string         `json:"honorific_suffix"`  // One or more honorifics following a person's name
	PatronymicName   *string         `json:"patronymic_name"`   // One or more patronymic names
	SortName         *string         `json:"sort_name"`         // A name to use in a lexicographically ordered list
	Email            *string         `json:"email"`             // A preferred email address
	Gender           *string         `json:"gender"`            // A gender
	BirthDate        Date            `json:"birth_date"`        // A date of birth
	DeathDate        Date            `json:"death_date"`        // A date of death
	Image            *string         `json:"image"`             // A URL of a head shot
	Summary          *string         `json:"summary"`           // A one-line account of a person's life
	Biography        *string         `json:"biography"`         // An extended account of a person's life
	NationalIdentify *string         `json:"national_identity"` // A national identity
	ContactDetails   []ContactDetail `json:"contact_details"`   // Means of contacting the person
	Links            []Link          `json:"link"`              // URLs to documents about the person
	Memberships      []Membership    `json:"memberships"`       // Memberships
	CreatedAt        time.Time       `json:"created_at"`        // The time at which the resource was created
	UpdatedAt        time.Time       `json:"updated_at"`        // The time at which the resource was last modified
	Sources          []Source        `json:"sources"`           // URLs to documents from which the person is derived
}

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

type OtherNames struct {
	Name            *string `json:"name"`             // [required] An alternate or former name
	FamilyName      *string `json:"family_name"`      // One or more family names
	GivenName       *string `json:"given_name"`       //One or more primary given names
	AdditionalName  *string `json:"additional_name"`  // One or more secondary given names
	HonorificPrefix *string `json:"honorific_prefix"` // One or more honorifics preceding a person's name
	HonorificSuffix *string `json:"honorific_suffix"` // One or more honorifics following a person's name
	PatronymicName  *string `json:"patronymic_name"`  // One or more patronymic names
	StartDate       Date    `json:"start_date"`       // The date on which the name was adopted
	EndDate         Date    `json:"end_date"`         // The date on which the name was abandoned
	Note            *string `json:"note"`             // A note, e.g. 'Birth name'
}

type Identifier struct {
	Identifier *string `json:"identifier"` // An issued identifier, e.g. a DUNS number
	Scheme     *string `json:"scheme"`     // An identifier scheme, e.g. DUNS
}

type ContactDetail struct {
	Label      *string   `json:"label"`       // A human-readable label for the contact detail
	Type       *string   `json:"type"`        //  [required] A type of medium, e.g. 'fax' or 'email'
	Value      *string   `json:"value"`       // [required] A value, e.g. a phone number or email address
	Note       *string   `json:"node"`        // A note, e.g. for grouping contact details by physical location
	ValidFrom  Date      `json:"valid_from"`  // The date from which the contact detail is valid",
	ValidUntil Date      `json:"valid_until"` // The date from which the contact detail is no longer valid",
	CreatedAt  time.Time `json:"created_at"`  // The time at which the resource was created
	UpdatedAt  time.Time `json:"updated_at"`  // The time at which the resource was last modified
	Sources    []Source  `json:"sources"`     // URLs to documents from which the person is derived
}

type Link struct {
	Url  *string `json:"url"`  // A URL
	Note *string `json:"node"` // A note, e.g. 'Wikipedia page'
}

type Membership struct{}

type Source Link

type Date struct {
	time.Time
}

type Area struct{} // TODO
type Post struct{} // TODO

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(d.Format(`"` + DateSpec + `"`)), nil
}
