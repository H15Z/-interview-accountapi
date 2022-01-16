# Form3 Take Home Exercise 
# Przemyslaw Sargeant 16/01/2022

## Instructions

run `docker-compose up`

*tests are executed after 5 secs to allow for the API to initialize*

## Instructions - Usage inside another application

`go get  github.com/H15Z/-interview-accountapi/client`

Example:

```golang
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/H15Z/-interview-accountapi/client"
)

func main() {
	c := client.New()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	resp, err := c.Accounts.FetchByResource(ctx, "/v1/organisation/accounts/771d850a-f6b3-4c16-b544-bbc8a05b740d")

	fmt.Println(resp, err)
}

```

### Checklist
- Be written in Go. **done**
- Use the `docker-compose.yaml` of this repository. **done**
- Be a client library suitable for use in another software project. **done and tested**
- Implement the `Create`, `Fetch`, and `Delete` operations on the `accounts` resource. **done**
- Be well tested to the level you would expect in a commercial environment. Note that tests are expected to run against the provided fake account API. **done**
- Be simple and concise. **subjective :)**
- Have tests that run from `docker-compose up` - our reviewers will run `docker-compose up` to assess if your tests pass. **done**

