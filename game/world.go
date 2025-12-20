package game

import "snake/entity"

type World struct {
	entities []entity.Entity
}

func NewWorld() *World {
	return &World{
		entities: []entity.Entity{},
	}
}

func (w *World) AddEntity(e entity.Entity) {
	w.entities = append(w.entities, e)
}

func (w World) Entities() []entity.Entity {
	return w.entities
}

func (w World) GetEntities(tag string) []entity.Entity {
	var result []entity.Entity

	for _, e := range w.entities {
		if e.Tag() == tag {
			result = append(result, e)
		}
	}
	return result
}

func (w World) GetFirstEntity(tag string) (entity.Entity, bool) {
	for _, e := range w.entities {
		if e.Tag() == tag {
			return e, true
		}
	}
	return nil, false
}

func (w *World) ClearAllEntity() {
	w.entities = []entity.Entity{}
}
