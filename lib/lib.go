package lib

import (
	"os/exec"
)

// Handleurl ...
func Handleurl(name string, outputdir string) {
	gitExecPath, _ := exec.LookPath("git")

	cmdGit := &exec.Cmd{
		Path:   gitExecPath,
		Args:   []string{gitExecPath, "clone", name, outputdir},
	}

  if err:= cmdGit.Run(); err != nil {

    println("oh no! we encounterd an error \\_(*-*)_/ \n")
    println("")
  }


  println("To add this package to your Cmake project add:")
  println("")

  println("add_subdirectory( \""+ outputdir +"\" )")
  println("target_link_libaries(\"${PROJECT_NAME} PUBLIC" +" "+outputdir+" "+ ") \n")


  println("")
  println("to your CMakeLists.txt")

}

// Handlename ...
func Handlename(name string, outputdir string) {
	gitExecPath, _ := exec.LookPath("git")

	cmdGit := &exec.Cmd{
		Path:   gitExecPath,
		Args:   []string{gitExecPath, "clone", "https://github.com/" + name + ".git", outputdir},
	}

  if err:= cmdGit.Run(); err != nil {

    println("oh no! we encounterd an error \\_(*-*)_/")
    println("")
  }


  println("To add this package to your Cmake project add:")
  println("")

  println("add_subdirectory( \""+ outputdir +"\" )")
  println("target_link_libaries(\"${PROJECT_NAME} PUBLIC" +" \""+outputdir+"\" "+ ")")

  println("")
  println("to your CMakeLists.txt")

}
