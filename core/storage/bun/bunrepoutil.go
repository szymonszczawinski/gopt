package bun

import (
	"context"
	"gosi/domain/lifecycle"
	"gosi/domain/project"
	"log"

	"github.com/uptrace/bun"
)

func mustInitDatabase(db *bun.DB, ctx context.Context) {
	if err := initStates(db, ctx); err != nil {
		return
	}
	if err := initLifecycles(db, ctx); err != nil {
		return
	}

	if err := initProjects(db, ctx); err != nil {
		return
	}

	initTransitions(db, ctx)
}

func initTransitions(db *bun.DB, ctx context.Context) {
	_, err := db.NewCreateTable().
		Model((*lifecycle.StateTransition)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Println(err)
	}
	//Project :: DRAFT -> NEW -> ANALISYS -> DESIGN -> DEV -> CLOSED
	transitions := []lifecycle.StateTransition{
		{LifecycleId: 1, FromStateId: 1, ToStateId: 2},
		{LifecycleId: 1, FromStateId: 2, ToStateId: 4},
		{LifecycleId: 1, FromStateId: 4, ToStateId: 5},
		{LifecycleId: 1, FromStateId: 5, ToStateId: 6},
		{LifecycleId: 1, FromStateId: 6, ToStateId: 10},
	}
	_, err = db.NewInsert().Model(&transitions).Exec(ctx)
	if err != nil {
		log.Println(err)
	}

}

func initProjects(db *bun.DB, ctx context.Context) error {
	_, err := db.NewCreateTable().
		Model((*project.ProjectRow)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	project := project.ProjectRow{
		Id:          1,
		Name:        "COSMOS",
		ProjectKey:  "COSMOS",
		Description: "Super COSMOS Project",
		StateId:     1,
		LifecycleId: 1,
		CreatedById: 1,
		OwnerId:     1,
	}
	_, err = db.NewInsert().Model(&project).Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func initLifecycles(db *bun.DB, ctx context.Context) error {
	_, err := db.NewCreateTable().
		Model((*lifecycle.LifecycleRow)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	lifecycle := lifecycle.LifecycleRow{Id: 1, Name: "Project", StartStateId: 1}
	_, err = db.NewInsert().Model(&lifecycle).Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func initStates(db *bun.DB, ctx context.Context) error {
	_, err := db.NewCreateTable().
		Model((*lifecycle.LifecycleStateRow)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	states := []lifecycle.LifecycleStateRow{
		{Id: 1, Name: "Draft"},
		{Id: 2, Name: "New"},
		{Id: 3, Name: "Open"},
		{Id: 4, Name: "Analysis"},
		{Id: 5, Name: "Design"},
		{Id: 6, Name: "Development"},
		{Id: 7, Name: "Integration"},
		{Id: 8, Name: "Verification"},
		{Id: 9, Name: "Fixed"},
		{Id: 10, Name: "Closed"},
		{Id: 11, Name: "Retest"},
		{Id: 12, Name: "Rejected"},
		{Id: 13, Name: "Assigned"},
	}

	_, err = db.NewInsert().Model(&states).Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
