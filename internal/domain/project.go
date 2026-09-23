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
var ErrProjectNameAlreadyExists = errors.New("project name already exists")

type ProjectID int64

type Project struct {
	ID          ProjectID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProject(name string, description string) (*Project, error) {
	normalizedName, err := NormalizeProjectName(name)
	if err != nil {
		return nil, err
	}

	return &Project{
		Name:        normalizedName,
		Description: description,
	}, nil
}

func (p *Project) ChangeName(name string) error {
	normalizedName, err := NormalizeProjectName(name)
	if err != nil {
		return err
	}

	p.Name = normalizedName

	return nil
}

func (p *Project) ChangeDescription(description string) error {
	p.Description = description

	return nil
}

func NormalizeProjectName(name string) (string, error) {
	normalizedName := strings.TrimSpace(name)

	if err := validateName(normalizedName); err != nil {
		return "", err
	}

	return normalizedName, nil
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
