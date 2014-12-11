package controllers

import "github.com/revel/revel"
import "time"
import "net/http"

type App struct {
	*revel.Controller
}

type Todo struct {
  Text string
  PostedAt time.Time
}

var todos []Todo

func init() {
  revel.TemplateFuncs["succ"] = func(a int) int { return a + 1 }
  revel.TemplateFuncs["date"] = func(t time.Time) string {
    const layout = "2006/01/02 15:04"
    jst := time.FixedZone("Asia/Tokyo", 9*60*60)
    
    return t.In(jst).Format(layout)
  }
}

func (c App) Index() revel.Result {
	return c.Render(todos)
}

func (c App) List() revel.Result {
  response := map[string]interface{}{
    "Status": "succeeded",
    "Todos": todos,
  }
  
  return c.RenderJson(response)
}

func (c App) Add(text string) revel.Result {
  if len(text) == 0 {
    c.Response.Status = http.StatusBadRequest
    return c.RenderJson(map[string]string{ "Status": "failed" })
  }
  
  todos = append(todos, Todo { Text: text, PostedAt: time.Now() })
  
  response := map[string]interface{}{
    "Status": "succeeded",
    "Todos": todos,
  }
  
  return c.RenderJson(response)
}

func (c App) Remove(index int) revel.Result {
  todos = append(todos[:index], todos[index+1:]...)
  
  response := map[string]interface{}{
    "Status": "succeeded",
    "Todos": todos,
  }
  
  return c.RenderJson(response)
}

func (c App) RemoveAll() revel.Result {
  todos = nil
  
  response := map[string]interface{}{
    "Status": "succeeded",
  }
  
  return c.RenderJson(response)
}