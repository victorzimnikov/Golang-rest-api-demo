package domain

import (
	"errors"
	"strings"
	"time"
)

var ErrProjectNameRequired = errors.New("project name is required")

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
	if name == "" {
		return nil, ErrProjectNameRequired
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
	if name == "" {
		return ErrProjectNameRequired
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
