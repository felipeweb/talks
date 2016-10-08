var ctx *Context
type Context struct {
	*macaron.Context
	render  jade.Render
	Session session.Store
	Flash   *session.Flash
}
func Contexter() macaron.Handler {
	return func(c *macaron.Context, r jade.Render, session session.Store, flash *session.Flash) {
		ctx = &Context{
			Context: c,
			render:  r,
			Session: session,
			Flash:   flash,
		}
		c.Map(ctx)
	}
}
func (ctx *Context) HTML(status int, name string) {
	ctx.render.HTML(status, name, ctx.Data)
}
func I18n(key string) string {
	return ctx.Tr(key)

}
