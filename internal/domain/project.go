package domain

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

const MaxProjectNameLength = 100

var ErrProjectNotFound = errors.New("project not found")
var ErrProjectNameRequired = errors.New("project name is required")
var ErrProjectNameTooLong = errors.New("project name is too long")

type ProjectID int64

type Project struct {
	ID          ProjectID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProject(name string, description string, now time.Time) (*Project, error) {
	name = strings.TrimSpace(name)

	if err := validateName(name); err != nil {
		return nil, err
	}

	return &Project{
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (p *Project) ChangeName(name string, now time.Time) error {
	name = strings.TrimSpace(name)

	if err := validateName(name); err != nil {
		return err
	}

	p.Name = name
	p.UpdatedAt = now

	return nil
}

func (p *Project) ChangeDescription(description string, now time.Time) error {
	p.Description = description
	p.UpdatedAt = now

	return nil
}

func validateName(name string) error {
	if name == "" {
		return ErrProjectNameRequired
	}

	if utf8.RuneCountInString(name) > MaxProjectNameLength {
		return ErrProjectNameTooLong
	}

	return nil
}
