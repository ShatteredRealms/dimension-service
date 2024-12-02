package game

import (
	"fmt"
	"regexp"
	"time"

	"github.com/ShatteredRealms/gameserver-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/model"
	"github.com/google/uuid"
)

const (
	MinMapNameLength = 1
	MaxMapNameLength = 64
)

var (
	MapNameRegex = "^[a-zA-Z0-9_- ]+$"

	// ErrMapNameToShort thrown when a map name is too short
	ErrMapNameToShort = fmt.Errorf("%w: name must be at least %d characters", ErrValidation, MinMapNameLength)

	// ErrMapNameToLong thrown when a map name is too long
	ErrMapNameToLong = fmt.Errorf("%w: name can be at most %d characters", ErrValidation, MaxMapNameLength)

	// ErrNameToLong thrown when a map name is too long
	ErrMapRegex = fmt.Errorf("%w: name can be alphanumeric with spaces, dashes and underscores", ErrValidation)
)

type Map struct {
	model.Model
	Name      string `gorm:"not null" json:"name"`
	MapPath   string `gorm:"not null" json:"map_path"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Maps []*Map

func (m *Map) Validate() error {
	return m.ValidateName()
}

func (m *Map) ValidateName() error {
	if len(m.Name) < MinMapNameLength {
		return ErrMapNameToShort
	}

	if len(m.Name) > MaxMapNameLength {
		return ErrMapNameToLong
	}

	ok, err := regexp.MatchString(MapNameRegex, m.Name)
	if !ok {
		return ErrMapRegex
	}
	if err != nil {
		return fmt.Errorf("%w: %w", ErrRegex, err)
	}

	return nil
}

func (m *Maps) HasMap(id *uuid.UUID) (int, bool) {
	for idx, mapItem := range *m {
		if mapItem.Id == id {
			return idx, true
		}
	}

	return -1, false
}

func (m *Maps) RemoveElement(idx int) {
	(*m)[idx] = (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
}

func (m *Map) ToPb() *pb.Map {
	return &pb.Map{
		Id:      m.Id.String(),
		Name:    m.Name,
		MapPath: m.MapPath,
	}
}

func (maps Maps) ToPb() *pb.Maps {
	out := make([]*pb.Map, len(maps))
	for idx, m := range maps {
		out[idx] = m.ToPb()
	}

	return &pb.Maps{
		Maps: out,
	}
}
