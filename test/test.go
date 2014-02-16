import (
	"github.com/azer/url-router"
)

func User () {}
func Fruit () {}
func FruitPage () {}
func Home () {}

Desc("url-router", func (it It) {
	routes := urlrouter.New("/user/:name", "/fruits/:fruit", "/fruits/:fruit/:page", "/")

	it("returns the matching pattern", func (expect Expect) {
		expect(routes.Match("/user/azer").Pattern).Equal("/user/:name")
		expect(routes.Match("/fruits/watermelon").Pattern).Equal("/fruits/:fruit")
		expect(routes.Match("/fruits/cherry/452").Pattern).Equal("/fruits/:fruit/:page")
		expect(routes.Match("/").Pattern).Equal("/")
	})

	it("returns matching parameters", func (expect Expect) {
		expect(routes.Match("/user/azer").Params["name"]).Equal("azer")
		expect(routes.Match("/fruits/watermelon").Params["fruit"]).Equal("watermelon")
		expect(routes.Match("/fruits/cherry/452").Params["fruit"]).Equal("cherry")
		expect(routes.Match("/fruits/cherry/452").Params["page"]).Equal("452")
	})

	it("returns nil when it can not match", func (expect Expect) {
		expect(routes.Match("/nonexisting") == nil).Equal(true)
	})

})
