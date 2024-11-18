package schema

import "github.com/diginfra/diginfra/internal/config"

type Provider interface {
	Type() string
	DisplayType() string
	ProjectName() string
	RelativePath() string
	VarFiles() []string
	AddMetadata(*ProjectMetadata)
	LoadResources(UsageMap) ([]*Project, error)
	Context() *config.ProjectContext
}
