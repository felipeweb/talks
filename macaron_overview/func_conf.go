m.Use(jade.Renderer(jade.Options{
  Funcs: []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
}))
