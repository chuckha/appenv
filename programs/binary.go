package programs

import "os/exec"

type Binary struct {
	Command string
	Args    []string
}

func (b *Binary) Version() ([]byte, error) {
	return exec.Command(b.Command, b.Args...).CombinedOutput()
}
func (b *Binary) GetCommand() string {
	return b.Command
}

var (
	Python = &Binary{
		Command: "python",
		Args:    []string{"--version"},
	}
	Ruby = &Binary{
		Command: "ruby",
		Args:    []string{"--version"},
	}
	Bash = &Binary{
		Command: "bash",
		Args:    []string{"--version"},
	}
	Git = &Binary{
		Command: "git",
		Args:    []string{"--version"},
	}
	SSH = &Binary{
		Command: "ssh",
		Args:    []string{"-V"},
	}
	Docker = &Binary{
		Command: "docker",
		Args: []string{"version", "--format", `Docker: {{.Client.Version}}
Docker Engine: {{.Server.Version}}`},
	}
)
