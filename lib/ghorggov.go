package ghogov

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Legislation defines the rules governing the organization
type Legislation struct {
	Doc    string // eg. rotblauer/gov
	Owners []string
}

func mustGetTokenFromConfig() (token string) {
	token = os.Getenv("GH_ORG_GOV_TOKEN")
	if token == "" {
		fmt.Println("missing token")
		os.Exit(1)
	}
	return
}

// WebhookServiceOpts configures the github webhook endpoint service.
type WebhookServiceOpts struct {
	Host      string
	Port      int
	CertsPath string
	DevMode   bool
	secret    string
}

func startGithubWebhookService(opts *WebhookServiceOpts) {
	// Build HTTPS-ready server

	// Listen for github webhook POST

	// Verify request authorization header (with secret+sha1(body)
	// https://developer.github.com/webhooks/#delivery-headers

	// If payload is PING reponse
	// https://developer.github.com/webhooks/#ping-event

	// Else if payload corresponds to a valid law (event type)
	// https://developer.github.com/webhooks/#events
	// https://developer.github.com/v3/activity/events/types/#memberevent
	// https://developer.github.com/v3/activity/events/types/#membershipevent
	// https://developer.github.com/v3/activity/events/types/#organizationevent
	// https://developer.github.com/v3/activity/events/types/#teamevent
	// https://developer.github.com/v3/activity/events/types/#teamaddevent

	//
	// Handle judiciary duties
	//  According to received webhook, did the webhook.sender follow the rules described in Legislation ?
	//  Legislation reference implementation ideas:
	//    = keep a 'cached' local copy of the designated Legislation document, retrieved from a specified url (eg. https://github.com/rotblauer/gov/tree/master/blob/LEGISLATION.yaml)
	//    = listen for another webhook that would broadcast when anything in this repo changes, and use git to keep an up-to-date copy of the repo, then use relative paths to specify the legistlation doc
	// Delegate executive duties based on legislation
}

func generateSecret() string {
	return ""
}

// Start starts the gov.
func Start() {
	// Setup Github client.
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: mustGetTokenFromConfig()},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	ghc := github.NewClient(tc)

	// Get LIST of organization hooks.
	// https://developer.github.com/v3/orgs/hooks/
	// If our host is NOT registered (no webhook set up for this client yet), then CREATE hook

	//

	// Set up server to listen for webhook events.
	startGithubWebhookService(&WebhookServiceOpts{
		Host:      "192.168.0.1",
		Port:      8080,
		CertsPath: filepath.Join(os.Getenv("HOME"), ".githubgov", "certs"),
		DevMode:   true,
		secret:    generateSecret(),
	})
}
