package domain

import (
	"errors"
	"log"
	"sync"

	"github.com/krakenlab/ternary"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/inputs"
)

// WorldRules struct
type WorldRules struct {
	Characters      []*entities.Character
	CharactersMutex sync.Mutex
}

// NewWorldRules constructor
func NewWorldRules() *WorldRules {
	return &WorldRules{}
}

// SpawnCharacter in character list
func (rule *WorldRules) SpawnCharacter(characterID int) {
	character := CharacterRepository.LoadCharacter(characterID)
	rule.appendCharacter(character)
}

// MoveCharacter in your world
func (rule *WorldRules) MoveCharacter(in *inputs.MoveCharacterInput) error {
	character, err := rule.findCharacter(in.CharacterID)

	if err == nil {
		in.DeltaX = ternary.Int(in.DeltaX > 1, 1, in.DeltaX)
		in.DeltaX = ternary.Int(in.DeltaX < -1, -1, in.DeltaX)
		in.DeltaY = ternary.Int(in.DeltaY > 1, 1, in.DeltaY)
		in.DeltaY = ternary.Int(in.DeltaY < -1, -1, in.DeltaY)

		character.MapXPosition += in.DeltaX
		character.MapYPosition += in.DeltaY
	}

	return err
}

// CharacterSpoke in this world
func (rule *WorldRules) CharacterSpoke(in *inputs.ChatInput) error {
	character, err := rule.findCharacter(in.CharacterID)

	if err == nil {
		log.Println(character.Name, "|>", in.Body)
	}

	return err
}

func (rule *WorldRules) appendCharacter(character *entities.Character) {
	rule.CharactersMutex.Lock()
	defer rule.CharactersMutex.Unlock()

	rule.Characters = append(rule.Characters, character)
}

func (rule *WorldRules) findCharacter(characterID int) (*entities.Character, error) {
	rule.CharactersMutex.Lock()
	defer rule.CharactersMutex.Unlock()

	for _, character := range rule.Characters {
		if int(character.ID) == characterID {
			return character, nil
		}
	}

	return nil, errors.New("character not found")
}

func (rule *WorldRules) findCharactersInDistanceFromPoint(x int, y int, maxDistance float32) []*entities.Character {
	rule.CharactersMutex.Lock()
	defer rule.CharactersMutex.Unlock()

	ret := []*entities.Character{}

	for _, character := range rule.Characters {
		ret = append(ret, character)
	}

	return ret
}
