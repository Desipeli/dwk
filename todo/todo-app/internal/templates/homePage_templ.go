// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func HomePage(backendPath string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><script src=\"https://unpkg.com/htmx.org@2.0.2\"></script><meta name=\"htmx-config\" content=\"{&#34;selfRequestsOnly&#34;: false}\"><title>Todo App</title></head><body style=\"margin: auto; display: block; width: 400px;\"><img src=\"/public/image.jpg\" alt=\"A random image\" width=\"400px\"><form id=\"todo-form\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(backendPath)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/homePage.templ`, Line: 18, Col: 24}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#todo-list\" hx-swap=\"innerHTML\" hx-on::after-request=\"clearTodoInput()\"><input type=\"text\" name=\"todo\" id=\"todo\" maxlength=\"140\"> <input type=\"submit\" value=\"Create TODO\"></form><br><p><b>Todos</b></p><ul hx-trigger=\"load\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(backendPath)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/homePage.templ`, Line: 28, Col: 44}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"innerHTML\" id=\"todo-list\"></ul><script>\n\n\t\t\t\tdocument.addEventListener(\"DOMContentLoaded\", function() {\n\t\t\t\t\tconst backendPath = \"{{backendPath}}\"\n\t\t\t\t\tconst baseUrl = window.location.origin\n\n\t\t\t\t\tdocument.getElementById(\"todo-form\").setAttribute(\"hx-post\", `${baseUrl}${backendPath}`)\n\t\t\t\t\tdocument.getElementById(\"todo-list\").setAttribute(\"hx-get\", `${baseUrl}${backendPath}`)\n\t\t\t\t})\n\n\t\t\t\tfunction clearTodoInput() {\n\t\t\t\t\ttodoInput = document.getElementById(\"todo\")\n\t\t\t\t\ttodoInput.value = \"\"\n\t\t\t\t}\n\t\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
