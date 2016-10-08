import (
    "github.com/go-macaron/i18n"
    "gopkg.in/macaron.v1"
)
func main() {
      m := macaron.Classic()
      m.Use(i18n.I18n(i18n.Options{
        Langs:    []string{"en-US", "zh-CN"},
        Names:    []string{"English", "简体中文"},
    }))

    m.Get("/", func(locale i18n.Locale) string {
        return "current language is" + locale.Lang
    })
    // Use in handler.
    m.Get("/trans", func(ctx *macaron.Context) string {
        return ctx.Tr("hello %s", "world")
    })
    m.Run()
}
// função disponivel na view
<h2>{{.i18n.Tr "hello %s" "world"}}!</h2>
