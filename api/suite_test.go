package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/backstage/maestro/account"
	"github.com/backstage/maestro/account/mem"
	"github.com/backstage/maestro/account/mongore"
	"github.com/backstage/maestro/api"
	"github.com/backstage/maestro/db"
	. "github.com/backstage/maestro/log"
	"github.com/backstage/maestro/requests"
	. "gopkg.in/check.v1"
)

var httpClient requests.HTTPClient

func Test(t *testing.T) { TestingT(t) }

var app account.App
var pluginConfig account.Plugin
var service account.Service
var team account.Team
var user account.User

type S struct {
	api        *api.Api
	authHeader string
	store      account.Storable
	server     *httptest.Server
	pubsub     account.PubSub
}

func (s *S) SetUpSuite(c *C) {
	Logger.Disable()
	// FIXME: add memory
	s.pubsub = account.NewEtcdSubscription("/api_test", &db.EtcdConfig{Machines: []string{"http://localhost:2379"}})
}

func (s *S) SetUpTest(c *C) {
	// setUpMongoreTest(s)
	setUpMemoryTest(s)

	s.api = api.NewApi(s.store, s.pubsub)
	s.server = httptest.NewServer(s.api.Handler())
	httpClient = requests.NewHTTPClient(s.server.URL)

	team = account.Team{Name: "Backstage Team", Alias: "backstage"}
	service = account.Service{Endpoint: "http://example.org/api", Subdomain: "backstage"}
	user = account.User{Name: "Bob", Email: "bob@bar.example.org", Password: "secret"}
	app = account.App{ClientId: "ios", ClientSecret: "secret", Name: "Ios App", Team: team.Alias, Owner: user.Email, RedirectUris: []string{"http://www.example.org/auth"}}
	pluginConfig = account.Plugin{Name: "Plugin Config", Service: service.Subdomain, Config: map[string]interface{}{"version": 1}}

	user.Create()
	token, err := s.api.Login(user.Email, "secret")
	if err != nil {
		panic(err)
	}
	s.authHeader = fmt.Sprintf("%s %s", token.Type, token.AccessToken)
}

func (s *S) TearDownTest(c *C) {
	user.Delete()
	s.server.Close()
}

var _ = Suite(&S{})

// Run the tests in memory
func setUpMemoryTest(s *S) {
	s.store = mem.New()
}

// Run the tests using MongoRe
func setUpMongoreTest(s *S) {
	s.store = mongore.New(mongore.Config{
		Host:         "127.0.0.1:27017",
		DatabaseName: "backstage_api_test",
	})
}

func testWithoutSignIn(reqArgs requests.Args, c *C) {
	headers, code, body, err := httpClient.MakeRequest(reqArgs)

	c.Check(err, IsNil)
	c.Assert(code, Equals, http.StatusUnauthorized)
	c.Assert(headers.Get("Content-Type"), Equals, "application/json")
	c.Assert(string(body), Equals, `{"error":"unauthorized_access","error_description":"Invalid or expired token. Please log in with your Backstage credentials."}`)
}
