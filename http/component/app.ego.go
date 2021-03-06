// Generated by ego.
// DO NOT EDIT

//line component/app.ego:1
package component

import "fmt"
import "html"
import "io"
import "context"

type App struct {
	Title string
	Yield func()
}

func (c *App) Render(ctx context.Context, w io.Writer) {

//line component/app.ego:11
	_, _ = io.WriteString(w, "<!doctype html>\n<html>\n\t<head>\n\t\t<meta charset=\"utf-8\">\n\t\t<meta http-equiv=\"x-ua-compatible\" content=\"ie=edge\">\n\t\t<title>\n\t\t\t")
//line component/app.ego:17
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(c.Title)))
//line component/app.ego:18
	_, _ = io.WriteString(w, "\n\t\t</title>\n\t\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n\n\t\t<link rel=\"stylesheet\" href=\"/styles/main.css\">\n\t</head>\n\t<body>\n\t\t")
//line component/app.ego:24
	if c.Yield != nil {
//line component/app.ego:24
		c.Yield()
//line component/app.ego:24
	}
//line component/app.ego:25
	_, _ = io.WriteString(w, "\n\t</body>\n</html>\n")
//line component/app.ego:27
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
