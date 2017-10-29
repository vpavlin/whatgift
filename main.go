package main

import (
	"fmt"
	"net/url"

	cn "github.com/vpavlin/whatgift/conceptnet"

	log "github.com/sirupsen/logrus"
	ig "github.com/yanatan16/golang-instagram/instagram"
)

func main() {
	cn := cn.NewConceptNet("http://api.conceptnet.io")
	n := cn.GetNode("example")
	fmt.Printf("%+v", n)

	for _, e := range n.Edges {
		log.Info(e.Rel)
	}

	tryIG(cn)
}

func tryIG(cn *cn.ConceptNet) {
	clientID := "8f02adbce03348798348b18221661051"
	clientSecret := "85fcbb41e06149bdbf2d2bc16969e039"
	userID := "30621586"
	fmt.Println("Hello")

	api := ig.New(clientID, clientSecret, "30621586.8f02adb.62abdac8dd2c4faaa9d81ff06c864518", false)

	params := url.Values{}
	//params.Set("count", "30")
	if resp, err := api.GetUserRecentMedia(userID, params); err != nil {
		panic(err)
	} else {
		for _, m := range resp.Medias {
			fmt.Println(m.Tags)
			for _, t := range m.Tags {
				n := cn.GetNodeIsA(t)
				for _, e := range n.Edges {
					if e.Rel.Label == "IsA" && e.Start.Label == t {
						if len(e.SurfaceText) == 0 {
							log.Info(e.End.Label)
						} else {
							log.Info(e.SurfaceText)
						}
					}
				}
			}
		}
		fmt.Println(resp.Pagination)
	}

	fmt.Println("ok")
}
