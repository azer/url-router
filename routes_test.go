package urlrouter

import (
	"testing"
	"fmt"
)

var (
	routes = New("/user/:name", "/fruits/:fruit", "/fruits/:fruit/:page", "/")
)

func TestMatching (t *testing.T) {
	equal(t, routes.Match("/user/azer").Pattern, "/user/:name")
  equal(t, routes.Match("/fruits/watermelon").Pattern, "/fruits/:fruit")
	equal(t, routes.Match("/fruits/cherry/452").Pattern, "/fruits/:fruit/:page")
	equal(t, routes.Match("/").Pattern, "/")
}

func TestParams (t *testing.T) {
	equal(t, routes.Match("/user/azer").Params["name"], "azer")
	equal(t, routes.Match("/fruits/watermelon").Params["fruit"], "watermelon")
	equal(t, routes.Match("/fruits/cherry/452").Params["fruit"], "cherry")
	equal(t, routes.Match("/fruits/cherry/452").Params["page"], "452")
}

func equal (t *testing.T, a string, b string) {
  if a != b {
		t.Error(fmt.Sprintf("%s and %s aren't equal", a, b))
  }
}

func User () {}
func Fruit () {}
func FruitPage () {}
func Home () {}
