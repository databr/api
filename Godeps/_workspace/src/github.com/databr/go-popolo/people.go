package popolo

import "time"

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
	DeathDate        DateTime        `json:"death_date"`        // A date of death
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
	Note       *string   `json:"note"`        // A note, e.g. for grouping contact details by physical location
	ValidFrom  Date      `json:"valid_from"`  // The date from which the contact detail is valid",
	ValidUntil Date      `json:"valid_until"` // The date from which the contact detail is no longer valid",
	CreatedAt  time.Time `json:"created_at"`  // The time at which the resource was created
	UpdatedAt  time.Time `json:"updated_at"`  // The time at which the resource was last modified
	Sources    []Source  `json:"sources"`     // URLs to documents from which the person is derived
}

type Link struct {
	Url  *string `json:"url"`  // A URL
	Note *string `json:"note"` // A note, e.g. 'Wikipedia page'
}

type Membership struct{}

type Source Link

type Area struct{} // TODO
type Post struct{} // TODO
