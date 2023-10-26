//go:build containers
// +build containers

package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"ru.mts.teta.tests_and_docs/internal/adapters/http"
)

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(ExampleSuit))
}

func (s *ExampleSuit) TestServerVerify() {

}

type ExampleSuit struct {
	suite.Suite

	server *http.Server
}

func (s *ExampleSuit) SetupSuite() {
	server := http.Server{}
	go func() {
		server.Start()
	}()
	s.server = &server

	s.T().Log("Suite setup is done")
}

func (s *ExampleSuit) TearDownSuite() {
	s.T().Log("Suite stop is done")
}

func (s *ExampleSuit) BeforeTest(suiteName, testName string) {}

func (s *ExampleSuit) AfterTest(suiteName, testName string) {}
