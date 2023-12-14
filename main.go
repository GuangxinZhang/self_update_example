package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const version = "1.0.0"
const repository = "GuangxinZhang/self_update_example"

func selfUpdate(repository string) error {
	selfupdate.EnableLog()

	previous := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(previous, repository)
	if err != nil {
		return err
	}

	if previous.Equals(latest.Version) {
		fmt.Println("Current binary is the latest version", version)
	} else {
		fmt.Println("Update successfully done to version", latest.Version)
		fmt.Println("Release note:\n", latest.ReleaseNotes)
	}
	return nil
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: self_update_example [flags]")
	flag.PrintDefaults()
}

func main() {
	help := flag.Bool("help", false, "Show this help")
	ver := flag.Bool("version", false, "Show version")

	flag.Usage = usage
	flag.Parse()

	if *help {
		usage()
		os.Exit(0)
	}

	if *ver {
		fmt.Println(version)
		os.Exit(0)
	}

	service()
}

func service() {
	log.Println("版本号：", version)
	if err := selfUpdate(repository); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	log.Println("版本号：", version)
	restart()
}

func restart() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal("Error locating executable path:", err)
	}
	cmd := exec.Command(exe, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		log.Fatal("Error occurred while restarting:", err)
	}
	fmt.Println("Restarting...")
	os.Exit(0)
}
