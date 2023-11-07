package documents

import (
	"github.com/buildbeaver/buildbeaver/common/models"
	"github.com/buildbeaver/buildbeaver/server/api/rest/routes"
)

type LegalEntity struct {
	baseResourceDocument
	RepoSearchURL   string `json:"repo_search_url"`
	RunnerSearchURL string `json:"runner_search_url"`
	BuildSummaryURL string `json:"build_summary_url"`
	// TODO: Include all model fields directly in this object
	*models.LegalEntity
}

func MakeLegalEntity(rctx routes.RequestContext, legalEntity *models.LegalEntity) *LegalEntity {
	return &LegalEntity{
		baseResourceDocument: baseResourceDocument{
			URL: routes.MakeLegalEntityLink(rctx, legalEntity.ID),
		},
		RepoSearchURL:   routes.MakeRepoSearchLink(rctx, legalEntity.ID),
		RunnerSearchURL: routes.MakeRunnerSearchLink(rctx, legalEntity.ID),
		BuildSummaryURL: routes.MakeBuildSummaryLink(rctx, legalEntity.ID),
		LegalEntity:     legalEntity,
	}
}

func MakeLegalEntities(rctx routes.RequestContext, list []*models.LegalEntity) []*LegalEntity {
	var docs []*LegalEntity
	for _, model := range list {
		docs = append(docs, MakeLegalEntity(rctx, model))
	}
	return docs
}
