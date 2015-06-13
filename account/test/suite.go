package test

import (
	"testing"

	"github.com/backstage/backstage/account"
	"github.com/backstage/backstage/errors"
	. "gopkg.in/check.v1"
)

var user account.User
var team account.Team
var service account.Service
var token account.TokenInfo

//Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type StorableSuite struct {
	Storage account.Storable
}

func (s *StorableSuite) SetUpTest(c *C) {
	user = account.User{Name: "Alice", Email: "alice@example.org", Password: "123456"}
	token = account.TokenInfo{Token: "secret-token", Expires: 10, Type: "Token", User: &user}
	team = account.Team{Name: "Backstage Team", Alias: "backstage", Users: []string{user.Email}, Owner: user.Email}
	service = account.Service{Endpoint: "http://example.org/api", Subdomain: "backstage", Team: team.Alias, Owner: user.Email, Transformers: []string{}}
}

func (s *StorableSuite) TestUpsertUser(c *C) {
	defer s.Storage.DeleteUser(user)
	err := s.Storage.UpsertUser(user)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestUpdateUser(c *C) {
	s.Storage.UpsertUser(user)
	user.Name = "Bob"
	defer s.Storage.DeleteUser(user)
	err := s.Storage.UpsertUser(user)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestTeams(c *C) {
	defer s.Storage.DeleteTeam(team)
	s.Storage.UpsertTeam(team)
	teams, err := s.Storage.UserTeams(account.User{Email: team.Owner})
	c.Check(err, IsNil)
	c.Assert(teams, DeepEquals, []account.Team{team})
}

func (s *StorableSuite) TestTeamsNotFound(c *C) {
	teams, err := s.Storage.UserTeams(account.User{})
	c.Check(err, IsNil)
	c.Assert(teams, DeepEquals, []account.Team{})
}

func (s *StorableSuite) TestDeleteUser(c *C) {
	s.Storage.UpsertUser(user)
	err := s.Storage.DeleteUser(user)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestDeleteUserNotFound(c *C) {
	err := s.Storage.DeleteUser(user)
	_, ok := err.(errors.NotFoundErrorNEW)
	c.Assert(ok, Equals, true)
}

func (s *StorableSuite) TestFindUserByEmail(c *C) {
	defer s.Storage.DeleteUser(user)
	s.Storage.UpsertUser(user)
	u, err := s.Storage.FindUserByEmail(user.Email)
	c.Assert(u, Equals, user)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestFindUserByEmailNotFound(c *C) {
	_, err := s.Storage.FindUserByEmail("not-found")
	_, ok := err.(errors.NotFoundErrorNEW)
	c.Assert(ok, Equals, true)
}

func (s *StorableSuite) TestUpsertTeam(c *C) {
	defer s.Storage.DeleteTeam(team)
	err := s.Storage.UpsertTeam(team)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestDeleteTeam(c *C) {
	s.Storage.UpsertTeam(team)
	err := s.Storage.DeleteTeam(team)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestDeleteTeamNotFound(c *C) {
	err := s.Storage.DeleteTeam(team)
	_, ok := err.(errors.NotFoundErrorNEW)
	c.Assert(ok, Equals, true)
}

func (s *StorableSuite) TestDeleteTeamByAlias(c *C) {
	s.Storage.UpsertTeam(team)
	err := s.Storage.DeleteTeamByAlias(team.Alias)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestDeleteTeamByAliasNotFound(c *C) {
	err := s.Storage.DeleteTeamByAlias(team.Alias)
	_, ok := err.(errors.NotFoundErrorNEW)
	c.Assert(ok, Equals, true)
}

func (s *StorableSuite) TestFindTeamByAlias(c *C) {
	defer s.Storage.DeleteTeam(team)
	s.Storage.UpsertTeam(team)
	u, err := s.Storage.FindTeamByAlias(team.Alias)
	c.Assert(u, DeepEquals, team)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestFindTeamByAliasNotFound(c *C) {
	_, err := s.Storage.FindTeamByAlias("not-found")
	_, ok := err.(errors.NotFoundErrorNEW)
	c.Assert(ok, Equals, true)
}

func (s *StorableSuite) TestTeamServices(c *C) {
	s.Storage.UpsertService(service)
	defer s.Storage.DeleteService(service)

	services, err := s.Storage.TeamServices(team)
	c.Assert(err, IsNil)
	c.Assert(services, DeepEquals, []account.Service{service})
}

func (s *StorableSuite) TestTeamServiceNotFound(c *C) {
	services, err := s.Storage.TeamServices(team)
	c.Assert(err, IsNil)
	c.Assert(services, DeepEquals, []account.Service{})
}

func (s *StorableSuite) TestCreateToken(c *C) {
	defer s.Storage.DeleteToken(token.Token)
	err := s.Storage.CreateToken(token)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestDeleteToken(c *C) {
	err := s.Storage.CreateToken(token)
	c.Check(err, IsNil)
	err = s.Storage.DeleteToken(token.Token)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestDecodeToken(c *C) {
	s.Storage.CreateToken(token)
	var u account.User
	s.Storage.DecodeToken(token.Token, &u)
	c.Assert(u, DeepEquals, user)
}

func (s *StorableSuite) TestUpsertService(c *C) {
	defer s.Storage.DeleteService(service)
	err := s.Storage.UpsertService(service)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestDeleteService(c *C) {
	s.Storage.UpsertService(service)
	err := s.Storage.DeleteService(service)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestDeleteServiceNotFound(c *C) {
	err := s.Storage.DeleteService(service)
	_, ok := err.(errors.NotFoundErrorNEW)
	c.Assert(ok, Equals, true)
}

func (s *StorableSuite) TestFindServiceBySubdomain(c *C) {
	defer s.Storage.DeleteService(service)
	s.Storage.UpsertService(service)
	serv, err := s.Storage.FindServiceBySubdomain(service.Subdomain)
	c.Assert(serv, DeepEquals, service)
	c.Check(err, IsNil)
}

func (s *StorableSuite) TestFindServiceBySubdomainNotFound(c *C) {
	_, err := s.Storage.FindServiceBySubdomain("not-found")
	_, ok := err.(errors.NotFoundErrorNEW)
	c.Assert(ok, Equals, true)
}

func (s *StorableSuite) TestUserServices(c *C) {
	s.Storage.UpsertTeam(team)
	defer s.Storage.DeleteTeam(team)
	s.Storage.UpsertService(service)
	defer s.Storage.DeleteService(service)
	another_service := account.Service{Endpoint: "http://example.org/api", Subdomain: "example", Team: team.Alias, Owner: user.Email, Transformers: []string{}}
	s.Storage.UpsertService(another_service)
	defer s.Storage.DeleteService(another_service)

	services, err := s.Storage.UserServices(account.User{Email: team.Owner})
	c.Check(err, IsNil)
	c.Assert(len(services), Equals, 2)
	c.Assert(services, DeepEquals, []account.Service{service, another_service})
}

func (s *StorableSuite) TestUserServicesNotFound(c *C) {
	services, err := s.Storage.UserServices(user)
	c.Assert(err, IsNil)
	c.Assert(services, DeepEquals, []account.Service{})
}
