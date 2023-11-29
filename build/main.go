package main

import (
	"github.com/buildbeaver/go-sdk/bb"
)

// TODO: Write a version of this fingerprinting that works with the Windows 'find' command
// TODO: (the two 'find' lines below don't work because they assume the syntax of the Unix find command)
var goJobFingerprint = []string{
	`@echo off`,
	`find build/scripts -type f | sort | xargs sha1sum`,
	`find backend/ -name '*.go' -not -path "*/vendor/*" -type f | sort | xargs sha1sum`,
	`sha1sum backend/go.mod`,
	`sha1sum backend/go.sum`,
}

var goEnvCmd = []string{
	"@echo off",
	`set BUILD_DIR=%cd%\build\output`,
	`set GODIR=%BUILD_DIR%\go`,
	`set GOBIN=%GODIR%\bin`,
	`set GOCACHE=%GODIR%\cache`,
}

func main() {
	bb.Workflows(
		bb.NewWorkflow().Name("base").Handler(submitBaseJobs),
		bb.NewWorkflow().Name("generate").Handler(submitGenerateJobs),
		bb.NewWorkflow().Name("unit-test").Handler(submitUnitTestJobs),
		bb.NewWorkflow().Name("integration-test").Handler(submitIntegrationTestJobs),
		bb.NewWorkflow().Name("build").Handler(submitBuildJobs),
		//		bb.NewWorkflow().Name("openapi").Handler(submitOpenAPIJobs),
	)
}

func submitBaseJobs(w *bb.Workflow) error {
	w.Job(bb.NewJob().
		Name("backend-preflight").
		Desc("Performs preflight checks on all backend code").
		Type(bb.JobTypeExec).
		//		Fingerprint(goJobFingerprint...).
		Step(bb.NewStep().
			Name("lint").
			//			Commands("bash build/scripts/ci/backend-preflight-lint.sh")))
			Commands("echo 'Skipping backend lint check...'")))

	return nil
}

func submitGenerateJobs(w *bb.Workflow) error {
	w.Job(bb.NewJob().
		Name("backend-generate").
		Desc("Generates all backend code (wire files, protobufs etc.)").
		Depends("base.backend-preflight").
		Type(bb.JobTypeExec).
		Fingerprint(goJobFingerprint...).
		Step(bb.NewStep().
			Name("wire").
			Commands(goEnvCmd...).
			Commands("bash build/scripts/ci/backend-generate-wire.sh")).
		Artifact(bb.NewArtifact().
			Name("wire").
			Paths("backend/*/app/wire_gen.go", "backend/*/app/*/wire_gen.go")).
		Artifact(bb.NewArtifact().
			Name("grpc").
			Paths("backend/api/grpc/*.pb.go")))

	return nil
}

func submitUnitTestJobs(w *bb.Workflow) error {
	w.Job(bb.NewJob().
		Name("backend-sqlite").
		Desc("Runs all backend unit tests on top of SQLite").
		Depends("generate.backend-generate.artifacts").
		Type(bb.JobTypeExec).
		Fingerprint(goJobFingerprint...).
		Env(bb.NewEnv().
			Name("TEST_DB_DRIVER").
			Value("sqlite3")).
		Step(bb.NewStep().
			Name("test").
			Commands(goEnvCmd...).
			Commands("cd backend",
				"go test -v -count=1 -mod=vendor -short ./...")))

	return nil
}

func submitIntegrationTestJobs(w *bb.Workflow) error {
	w.Job(bb.NewJob().
		Name("backend-sqlite").
		Desc("Runs all backend integration tests on top of SQLite").
		Depends("generate.backend-generate.artifacts").
		Type(bb.JobTypeExec).
		Fingerprint(goJobFingerprint...).
		Env(bb.NewEnv().
			Name("TEST_DB_DRIVER").
			Value("sqlite3")).
		Step(bb.NewStep().
			Name("test").
			Commands(goEnvCmd...).
			Commands(
				"cd backend",
				"go test -v -count=1 -mod=vendor -run Integration ./...")))

	return nil
}

func submitBuildJobs(w *bb.Workflow) error {
	w.Job(bb.NewJob().
		Name("backend-build").
		Desc("Builds all backend binaries").
		Depends(
			"generate.backend-generate.artifacts",
			"unit-test.backend-sqlite",
			"integration-test.backend-sqlite").
		Type(bb.JobTypeExec).
		Fingerprint(goJobFingerprint...).
		Step(bb.NewStep().
			Name("go").
			Commands(goEnvCmd...).
			Commands("bash build/scripts/ci/backend-build.sh")).
		Artifact(bb.NewArtifact().
			Name("go-binaries").
			Paths("build/output/go/bin/*")))
	return nil
}

func submitOpenAPIJobs(w *bb.Workflow) error {
	w.Job(bb.NewJob().
		Name("generate-code").
		Desc("Generates API Clients from our OpenAPI specs").
		Depends("base.backend-preflight").
		Type(bb.JobTypeExec).
		Fingerprint("sha1sum backend/server/api/rest/openapi/dynamic-openapi.yaml").
		Step(bb.NewStep().
			Name("dynamic-go").
			Commands("bash build/scripts/ci/backend-openapi.sh dynamic-openapi.yaml go sdk/go-sdk/bb/client")))
	return nil
}
