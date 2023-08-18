package model

import (
	"strconv"
	"strings"
)

type BatriumIdentifier struct {
	Id string
}

func NewBatriumIdentifier(hubId uint16, systemId uint16) BatriumIdentifier {
	return BatriumIdentifier{Id: strconv.Itoa(int(hubId)) + "-" + strconv.Itoa(int(systemId))}
}

func (batriumIdentifier *BatriumIdentifier) GetHubId() string {
	return strings.Split(batriumIdentifier.Id, "-")[0]
}

func (batriumIdentifier *BatriumIdentifier) GetSystemId() string {
	return strings.Split(batriumIdentifier.Id, "-")[1]
}
