package main

import (
	"encoding/json"
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/external"
)

func main() {
	c, err := external.LoadDefaultUCloudConfig()
	if err != nil {
		panic(err)
	}

	cfg, cred := c.Config(), c.Credential()

	bs, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Printf("[Config] %s\n", string(bs))

	bs, _ = json.MarshalIndent(cred, "", "  ")
	fmt.Printf("[Credential] %s\n", string(bs))
}
