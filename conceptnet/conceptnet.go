package conceptnet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ConceptNet struct {
	APIURL string
}

type Rel struct {
	ID    string `json:"@id"`
	Label string `json:"label"`
}

type Node struct {
	ID       string `json:"@id"`
	Edges    []Edge `json:"edges"`
	Label    string `json:"label"`
	Language string `json:"language"`
	Term     string `json:"term"`
}

type Edge struct {
	ID          string  `json:"@id"`
	Rel         Rel     `json:"rel"`
	End         Node    `json:"end"`
	Start       Node    `json:"start"`
	SurfaceText string  `json:"surfaceText"`
	Weight      float32 `json:"weight"`
}

func NewConceptNet(apiURL string) *ConceptNet {
	return &ConceptNet{
		APIURL: apiURL,
	}
}

func (c *ConceptNet) GetNode(name string) *Node {
	q := fmt.Sprintf("%s/c/en/%s", c.APIURL, name)
	log.Info(q)
	resp, err := http.Get(q)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic("Wrong return code: " + resp.Status)
	}

	n := Node{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &n)
	if err != nil {
		panic(err)
	}

	return &n
}

func (c *ConceptNet) GetNodeIsA(name string) *Node {
	q := fmt.Sprintf("%s/query?node=/c/en/%s&start=/c/en/%s&rel=/r/IsA", c.APIURL, name, name)
	log.Info(q)
	resp, err := http.Get(q)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic("Wrong return code: " + resp.Status)
	}

	n := Node{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &n)
	if err != nil {
		panic(err)
	}

	return &n
}
