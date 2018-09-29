package models

import (
	"database/sql"

	slugify "github.com/gosimple/slug"
	"github.com/metapods/metapods/data/marshal"
)

// Group is a collection of Organizations
// Refers to the https://www.w3.org/TR/activitystreams-vocabulary/#dfn-group
// Also refers to the Groups table in the database
type Group struct {
	Slug      string                  `json:"slug"`
	Name      string                  `json:"name"`
	Note      string                  `json:"note"`
	CreatedAt marshal.MarshalableTime `json:"created_at"`
	UpdatedAt marshal.MarshalableTime `json:"updated_at"`
}

// GetGroup returns a single Group object or nil
func GetGroup(db *sql.DB, slug string) (*Group, error) {
	row := db.QueryRow(`
		select slug, name, note, created_at, updated_at
		from groups;
	`)

	var group Group
	err := row.Scan(&group.Slug,
		&group.Name, &group.Note, &group.CreatedAt, &group.UpdatedAt)

	// This is not an error from the user's perspective
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// PutGroup creates a group with this name and note
func PutGroup(db *sql.DB, name string, note string) (string, error) {
	slug := slugify.MakeLang(name, "en")

	query := `
		INSERT INTO groups (slug, name, note)
		VALUES ($1, $2, $3)
	`

	_, err := db.Exec(query, slug, name, note)
	return slug, err
}

// Organization is someone who owns a episodes of podcasts
// Refers to the https://www.w3.org/TR/activitystreams-vocabulary/#dfn-organization
// Also refers to the Organizations table in the database
type Organization struct {
	Slug      string                  `json:"slug"`
	Name      string                  `json:"name"`
	Note      string                  `json:"note"`
	CreatedAt marshal.MarshalableTime `json:"created_at"`
	UpdatedAt marshal.MarshalableTime `json:"updated_at"`
}

// Podcast is a something with an audio link, a name, and a note
// Refers to the Podcasts table in the database
type Podcast struct {
	Slug         string                  `json:"slug"`
	Name         string                  `json:"name"`
	Note         string                  `json:"note"`
	ThumbnailURL string                  `json:"thumbnail_url"`
	AudioURL     string                  `json:"audio_url"`
	MediaType    string                  `json:"media_type"`
	PostedAt     marshal.MarshalableTime `json:"posted_at"`
	CreatedAt    marshal.MarshalableTime `json:"created_at"`
	UpdatedAt    marshal.MarshalableTime `json:"updated_at"`
}
