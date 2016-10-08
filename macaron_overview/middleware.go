func SetupMiddlewares(app *macaron.Macaron) {
	app.Use(i18n.I18n(i18n.Options{
		Directory: "locale",
		Langs:     []string{"pt-BR", "en-US"},
		Names:     []string{"Protuguês do Brasil", "American English"},
	}))
	app.Use(jade.Renderer(jade.Options{
		Directory: "public/templates",
	}))
}
//handler
m.Get("/", func(r jade.Render, ctx *macaron.Context) {
    r.HTML(http.StatusOK, "login", ctx.Data)
})
// Set language properties.
		locale := Locale{i18n.Locale{lang}}
		ctx.Map(locale)
		ctx.Locale = locale
		ctx.Data[opt.TmplName] = locale
		ctx.Data["Tr"] = i18n.Tr
		ctx.Data["Lang"] = locale.Lang
		ctx.Data["LangName"] = curLang.Name
		ctx.Data["AllLangs"] = append([]LangType{curLang}, restLangs...)
		ctx.Data["RestLangs"] = restLangs
