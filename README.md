# HubSpot

[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/retgits/hubspot)

## Why build this?

[HubSpot](https://app.hubspot.com) has a great [API](https://developers.hubspot.com/docs/overview) to allow you to create a functional application or integration quickly and easily. The APIs they expose are the same that power the HubSpot application. There isn't a good abstraction to be able to use those in Go apps, unless you want to marshal and unmarshal things yourself and have code with tons of HTTP requests.

## Usage

To use the HubSpot module, you'll need to create an [API key](https://developers.hubspot.com/docs/methods/auth/oauth-overview). That key can be used to create the _hubspot client_.

```go
import (
	"fmt"

	"github.com/retgits/hubspot/client"
)

func main() {
    hubspot := client.NewClient().WithAPIKey("<myAccessToken>")
	fmt.Println(hubspot.APIKey)
}
```

Depending on which type of resource you want to access, you'll need to import one of the services

```go
import (
    "github.com/retgits/hubspot/client/contacts" // If you want to use the contacts API
    "github.com/retgits/hubspot/client/crmassociations" // If you want to use the crm associations API
    "github.com/retgits/hubspot/client/deals" // If you want to use the deals API
    "github.com/retgits/hubspot/client/engagement" // If you want to use the engagements API
    "github.com/retgits/hubspot/client/tickets" // If you want to use the tickets API
)
```

For example, getting the first name of the most recently updated contact in HubSpot would be like:

```go
package main

import (
	"fmt"

	"github.com/retgits/hubspot/client"
	"github.com/retgits/hubspot/client/contacts"
)

func main() {
    // Create a new HubSpot client
    hubspot := client.NewClient().WithAPIKey("<myAccessToken>")
    // Create a new Contacts service
    contactsSvc := contacts.New(hubspot).WithProperties([]string{"firstname"})
	// Get all the recentlu updated contacts
    contacts, _err_ := contactsSvc.GetRecentlyUpdatedContacts()
    // Print the first in the array
    fmt.Printf("%s", contacts[0].Properties["firstname"]["value"])
}
```

## Contributing

Currently the methods I use regularly are implemented, so chances are that something you might need is missing. If something is missing, or if you'd like to suggest new features feel free to [create an issue](https://github.com/retgits/hubspot/issues/new) or a [PR](https://github.com/retgits/hubspot/compare)!

## License

See the [LICENSE](./LICENSE) file in the repository

## Acknowledgements

A most sincere thanks to the team of [HubSpot](https://hubspot.com), for building a great CRM!

_This package is not endorsed by HubSpot