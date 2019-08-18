package domain

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dmgk/faker"
	"github.com/krakenlab/ternary"

	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
	"github.com/schweigert/mmosandbox/domain/presenters"
	"github.com/schweigert/mmosandbox/lib/bench"
)

// WorldRules struct
type WorldRules struct {
	Characters      []*entities.Character
	CharactersMutex sync.Mutex
	FieldOfVision   float64

	WorldName string
}

// NewWorldRules constructor
func NewWorldRules() (worldRules *WorldRules) {
	worldRules = &WorldRules{
		FieldOfVision: 25.0,
		WorldName:     presenters.WorldNamePresenter(faker.Address().City()),
	}

	go func() {
		for {
			worldRules.WorldLoop()
		}
	}()

	return
}

// WorldLoop :)
func (rule *WorldRules) WorldLoop() {
	rule.SendMetrics()
}

// SendMetrics to graphite
func (rule *WorldRules) SendMetrics() {
	bench.Metronome.Bip(fmt.Sprintf("worlds.%s.characters", rule.WorldName), fmt.Sprintf("%d", len(rule.Characters)))
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
	if in.HasBody() {
		character, err := rule.findCharacter(in.CharacterID)
		if err == nil {
			message := entities.NewMessageFromCharacter(character, in.Body)
			rule.spoke(character, message)
		}
	}

	return nil
}

// CharacterListen over this world
func (rule *WorldRules) CharacterListen(in *inputs.ChatInput, out *outputs.ListenMessagesOutput) error {
	character, err := rule.findCharacter(in.CharacterID)

	if err == nil {
		out.Messages = character.MessageBox.Read()
	}

	return rule.CharacterSpoke(in)
}

func (rule *WorldRules) spoke(character *entities.Character, message *entities.Message) {
	for _, otherCharacter := range rule.findCharactersInFieldOfVision(character) {
		otherCharacter.MessageBox.Append(message)
	}
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

func (rule *WorldRules) findCharactersInFieldOfVision(pivot *entities.Character) []*entities.Character {
	rule.CharactersMutex.Lock()
	defer rule.CharactersMutex.Unlock()

	ret := []*entities.Character{}

	for _, character := range rule.Characters {
		if pivot.ID != character.ID && character.DistanceTo(pivot.MapXPosition, pivot.MapYPosition) <= rule.FieldOfVision {
			ret = append(ret, character)
		}
	}

	return ret
}
