package generator

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/goombaio/namegenerator"
	"github.com/yaminmhd/go-kafka-producer-protobuf/generatedProtos/person"
	"math/rand"
	"time"
)

func NewPerson() *person.PersonMessage {
	timestamp, _ := ptypes.TimestampProto(time.Now().UTC())
	person := &person.PersonMessage{
		Name:  randomNameGenerator(),
		Id:    int32(randomId()),
		Email: randomEmail(),
		Phones: []*person.PhoneNumber{
			{Number: randomPhoneNumber(), Type: randomPhoneType()},
		},
		LastUpdated: timestamp,
	}
	return person
}

func randomNameGenerator() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	return name
}

func randomId() int {
	rand.Seed(time.Now().UnixNano())
	return random(90000000,99999999)

}

func randomPhoneNumber() string {
	return fmt.Sprintf("%v", randomId())
}

func randomEmail() string {
	return fmt.Sprintf("%v@gmail.com", randomNameGenerator())
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func randomPhoneType() person.PhoneType {
	randomSource := rand.NewSource(time.Now().UnixNano())
	source := rand.New(randomSource)
	switch source.Intn(3) {
	case 1:
		return 0
	case 2:
		return 1
	default:
		return 2
	}
}