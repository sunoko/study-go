package repository

import (
	"github.com/gemcook/study-go/20181003/api_example/model"
	"github.com/go-xorm/xorm"
)

type Gems struct {
	engine *xorm.Engine
}

func NewGems(engine *xorm.Engine) *Gems {
	g := Gems{
		engine: engine,
	}

	return &g
}

func (g *Gems) Create(input *model.Gem) error {
	_, err := g.engine.Insert(input)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gems) GetAll() ([]*model.Gem, error) {

}
