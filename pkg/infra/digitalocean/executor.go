package infra

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

type InfraExecutor struct {
}

func InitInfraExecutor() *InfraExecutor {
	return &InfraExecutor{}
}

func (infra *InfraExecutor) Apply(execPath string) (publicIP string, err error) {
	fmt.Println("[INFO] Terraform init in progress ...")

	cmd := exec.Command("terraform", "init")
	cmd.Dir = execPath
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running 'terraform init': %w for %s", err, execPath)
	}

	fmt.Println("[OK] Terraform init done!")

	fmt.Println("[INFO] Terraform apply in progress ...")

	cmd = exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = execPath

	var out bytes.Buffer
	var stdErr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stdErr

	err = cmd.Run()
	if err != nil {
		fmt.Println(stdErr.String())
		return "", fmt.Errorf("error running 'terraform apply': %w for %s", err, execPath)
	}

	fmt.Println("[OK] Terraform init done!")

	// Extract the public IP from the Terraform output
	re := regexp.MustCompile(`public_ip\s+=\s+\"(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\"`)
	matches := re.FindStringSubmatch(out.String())
	if len(matches) != 2 {
		return "", fmt.Errorf("failed to extract public IP from Terraform output")
	}

	publicIP = matches[1]

	fmt.Println("[OK] VPN public IP:", publicIP)

	return
}

func (infra *InfraExecutor) Destroy(execPath string) error {
	fmt.Println("[INFO] Terraform destroy in progress ...")

	cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd.Dir = execPath
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error running 'terraform destroy': %w for %s", err, execPath)
	}

	fmt.Println("[OK] Terraform destroy done!")

	return nil
}