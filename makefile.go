//+build mage

//Build, test and more ... a developers everyday tool
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bclicn/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	service = "mail"
)

type Helper mg.Namespace
type Build mg.Namespace
type Clean mg.Namespace

// Remove the .temp folder
func (Clean) temp() error {

	// parallelize execution
	go func() {
		tempFolder := fmt.Sprintf("./%s/.temp", service)

		// clean any existing .temp folder
		if _, err := os.Stat(tempFolder); !os.IsNotExist(err) {
			err := sh.Run("rm", "-rf", tempFolder)
			if err != nil {
				fmt.Printf(color.Red("./.temp folder rm error: %s\n"), err)
			}
		}

		// create new .temp folder for future building and packing
		err := os.Mkdir(tempFolder, 0755)
		if err != nil {
			fmt.Printf(color.Red("mkdir error: %s\n"), err)
		}
	}()

	return nil
}

// Transpile protobuffer definitions
func (Build) Protoc() error {

	// find all *.protoc files
	cmdOut, err := sh.Output("bash", "-c", fmt.Sprintf("find %s -type f -name \"*.proto\" -exec ls {} \\;", "./"))
	if err != nil {
		fmt.Printf(color.Red("Build:Protoc find .*.proto error: %s\n"), err)
		return err
	}

	files := strings.Split(cmdOut, "\n")

	// iterate over all found protoc files
	for _, file := range files {
		// transpile .proto file
		err := sh.Run("protoc", "--proto_path=.", "--twirp_out=.", "--go_out=.", file)
		if err != nil {
			fmt.Printf(color.Red("Build:Protoc protoc transpile error: %s\n"), err)
			return err
		}
	}

	return nil
}

// Compile the service
func (Build) Build() error {

	mg.Deps(Clean.temp)

	servicePath := "./" + service + "/"

	err := sh.Run("env", "CGO_ENABLED=0", "GOOS=linux", "go", "build", "-o", servicePath+"/.temp/"+service, servicePath+"/server/"+service)
	if err != nil {
		fmt.Printf(color.Red("Build:Build go build error: %s\n"), err)
		return err
	}

	// if the service has a /templates folder, copy it to .temp image building source folder
	if _, err := os.Stat(servicePath + "/templates"); !os.IsNotExist(err) {
		err := sh.Run("cp", "-r", servicePath+"/templates", servicePath+"/.temp")
		if err != nil {
			fmt.Printf(color.Red("Build:Build cp templates folder error: %s\n"), err)
			return err
		}
	}

	return nil
}

// Running the service
func (Build) Run() error {
	mg.Deps(Build.Build)

	fmt.Printf("Running service %s...\n", service)

	err := sh.Run("./mail/.temp/mail")
	if err != nil {
		fmt.Printf("Error %s...\n", err)
		return err
	}

	return nil
}

// Testing the service
func (Build) Test() error {
	fmt.Printf("%s testing service %s...\n", color.Cyan("Microservice"), service)

	_, err := sh.Output("go", "run", ""+"./"+service+"/test/main.go")
	return err
}
