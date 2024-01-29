package student

import (
	"fmt"
	"strings"
)

type Game struct {
	Person
}

func (g *Game) HandleCommand(command string) string {
	path := strings.Fields(command)
	switch {
	case len(path) == 1 && path[0] == "осмотреться":
		return fmt.Sprintf("%s. можно пройти - %s", g.Around(), strings.Join(g.Path(), ", "))
	case len(path) == 2 && path[0] == "идти":
		location := g.Find(path[1])
		if location != nil {
			g.Location = location
			return fmt.Sprintf("%s. можно пройти - %s", g.Enter(), strings.Join(g.Path(), ", "))
		}
		return fmt.Sprintf("нет пути в %s", path[1])
	case len(path) == 2 && path[0] == "взять":
		if g.Pop(path[1]) {
			return fmt.Sprintf("предмет добавлен в инвентарь: %s", path[1])
		}
		return "нет такого"
	}
	return "неизвестная команда"
}

type Location interface {
	Name() string
	Around() string
	Path() []string
	Find(string) Location
	Enter() string
	Pop(string) bool
}

type Person struct {
	Location
}

type Route []Location

func (r Route) Find(to string) Location {
	for _, location := range r {
		if to == location.Name() {
			return location
		}
	}
	return nil
}

func (r Route) Path() (paths []string) {
	for _, location := range r {
		paths = append(paths, location.Name())
	}
	return
}

type Corridor struct {
	Route
}

func (c Corridor) Name() string {
	return "коридор"
}

func (c Corridor) Around() string {
	return "не стой в коридоре"
}

func (c Corridor) Enter() string {
	return "ничего интересного"
}

func (c Corridor) Pop(string) bool { return false }

type Kitchen struct {
	Route
}

func (k Kitchen) Name() string {
	return "кухня"
}

func (k Kitchen) Around() string {
	return "ты находишься на кухне, надо собрать рюкзак и идти в универ"
}

func (k Kitchen) Enter() string {
	return "кухня, ничего интересного"
}

func (k Kitchen) Pop(string) bool { return false }

type Items map[string][]string

func (i Items) Pop(item string) bool {
	for on, items := range i {
		for j := range items {
			if item == items[j] {
				items = append(items[:j], items[j+1:]...)
				if len(items) > 0 {
					i[on] = items
				} else {
					delete(i, on)
				}
				return true
			}
		}
	}
	return false
}

func (i Items) String() string {
	var q bool
	var b strings.Builder
	for on, items := range i {
		if q {
			b.WriteString("; ")
		}
		b.WriteString(on)
		b.WriteString(": ")
		b.WriteString(strings.Join(items, ", "))
		q = true
	}
	return b.String()
}

func (i Items) Empty() bool {
	return len(i) == 0
}

type Room struct {
	Route
	Items
}

func (r Room) Name() string {
	return "комната"
}

func (r Room) Around() string {
	if !r.Empty() {
		return r.String()
	}
	return "пустая комната"
}

func (r Room) Enter() string {
	return "ты в своей комнате"
}

type Street struct {
	Route
}

func (s Street) Find(path string) Location {
	if path == "домой" {
		return s.Route.Find("коридор")
	}
	return s.Route.Find(path)
}

func (s Street) Path() []string {
	return []string{"домой"}
}

func (s Street) Name() string {
	return "улица"
}

func (s Street) Around() string {
	return "переходи дорогу на зеленый свет"
}

func (s Street) Enter() string {
	return "на улице весна"
}

func (s Street) Pop(string) bool { return false }

func New() *Game {
	var corridor Corridor
	var kitchen Kitchen
	var street Street
	var room Room
	room.Items = Items{"на столе": []string{"ключи", "конспекты", "рюкзак"}} //, "на полу": []string{"носки"}}
	room.Route = Route{&corridor}
	street.Route = Route{&corridor}
	corridor.Route = Route{&kitchen, &room, &street}
	kitchen.Route = Route{&corridor}
	return &Game{
		Person: Person{
			Location: &kitchen,
		},
	}
}
