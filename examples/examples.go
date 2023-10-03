package main

import (
	"fmt"
	"os"

	"github.com/souleb/dag"
)

// Job is a job to be executed
type Job struct {
	cmd       string
	DependsOn []string
	workDir   string
	arch      string
	os        string
}

func main() {
	build := jobToVertex("build", Job{cmd: "make build", DependsOn: []string{"unit-test"}, workDir: "/tmp", arch: "amd64", os: "linux"})
	unitTest := jobToVertex("unit-test", Job{cmd: "make unit-test", DependsOn: []string{}, workDir: "/tmp", arch: "amd64", os: "linux"})
	deploy := jobToVertex("deploy", Job{cmd: "make deploy", DependsOn: []string{"build"}, workDir: "/tmp", arch: "amd64", os: "linux"})
	g := dag.New[Job]()
	g.Add(deploy)
	g.Add(unitTest)
	g.Add(build)
	g.AddEdge(unitTest, build, 1)
	g.AddEdge(build, deploy, 1)

	sort, err := g.TopologicalSort()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%v", *sort)
}

func jobToVertex(name string, job Job) dag.Vertex[Job] {
	return dag.NewVertex[Job](name, job)
}
