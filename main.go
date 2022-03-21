package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	fmt.Printf("Start %v\n", os.Args[1:])
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Bad command")
	}
}

func run() {
	fmt.Printf("R: Running %v as %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	// Conectando file descriptors e criando um fluxo de output
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("C: Running %v as %d\n", os.Args[2:], os.Getpid())

	// cg()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	// Conectando file descriptors e criando um fluxo de output
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))
	// must(syscall.Chroot("/ROOT_FOR_CONTAINER"))
	// must(syscall.Chdir("/"))
	// must(syscall.Mount("proc", "proc", "proc", 0, ""))
	// must(syscall.Mount("thing", "mytemp", "tmpfs", 0, ""))

	must(cmd.Run())

	// must(syscall.Unmount("proc", 0))

	fmt.Print("finalizaed\n")
}

func cg() {
	cgroups := "/sys/fs/cgroup"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "liz"), 0755)
	must(ioutil.WriteFile(filepath.Join(pids, "liz/pids.max"), []byte("20"), 0700))
	// Remove o novo grupo depois que o container finalizar
	must(ioutil.WriteFile(filepath.Join(pids, "liz/notify_on_release"), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "liz/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
