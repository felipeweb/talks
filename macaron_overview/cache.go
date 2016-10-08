import (
    "github.com/go-macaron/cache"
    "gopkg.in/macaron.v1"
)

func main() {
    m := macaron.Classic()
    m.Use(cache.Cacher())

    m.Get("/", func(c cache.Cache) string {
        c.Put("cache", "cache middleware", 120)
        return c.Get("cache")
    })

    m.Run()
}
