// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/api"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	grpchandler "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/grpc"
	planhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
	"google.golang.org/grpc"
)

type Plan struct {
	Handler     *planhttp.PlanHandler
	GRPCHandler api.PlanServiceServer
	Store       store.Plan
}

func NewPlan(*config.Plans) *Plan {
	store := memory.NewPlanStore()
	return &Plan{
		Handler:     planhttp.NewPlanHandler(store),
		GRPCHandler: grpchandler.NewPlanServer(store),
		Store:       store,
	}
}

func (a *Plan) RegisterRoutes(mux *http.ServeMux, grpcSrv *grpc.Server) {
	mux.HandleFunc("GET /plans", a.Handler.List)
	mux.HandleFunc("POST /plans", a.Handler.Create)
	mux.HandleFunc("GET /plans/{id}", a.Handler.Get)
	mux.HandleFunc("PUT /plans/{id}", a.Handler.Update)
	mux.HandleFunc("DELETE /plans/{id}", a.Handler.Delete)

	api.RegisterPlanServiceServer(grpcSrv, a.GRPCHandler)
}
