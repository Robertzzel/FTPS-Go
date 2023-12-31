package commands

import (
	"FTP-NAS-SV/codes"
	"FTP-NAS-SV/utils"
	"os"
	"path"
	"strings"
)

type RMD struct {
	currentPath string
	parameters  []string
	user        *utils.User
}

func NewRMDCommand(params []string, currentPath string, user *utils.User) RMD {
	return RMD{parameters: params, currentPath: currentPath, user: user}
}

func (cmd RMD) Execute() (int, error) {
	if !cmd.user.IsLogenIn() {
		return codes.NotLoggedIn, nil
	}
	if len(cmd.parameters) != 2 {
		return codes.SyntaxErrorParametersArguments, nil
	}
	if !(strings.HasPrefix(cmd.parameters[1], "./") || strings.HasPrefix(cmd.parameters[1], "/")) {
		return codes.SyntaxErrorParametersArguments, nil
	}

	var dirPath string
	if path.IsAbs(cmd.parameters[1]) {
		dirPath = cmd.parameters[1]
	} else {
		dirPath = path.Join(cmd.currentPath, cmd.parameters[1])
	}

	if !strings.HasPrefix(dirPath, cmd.user.BasePath) {
		return codes.RequestedActionNotTaken, nil
	}
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return codes.RequestedActionNotTaken, nil
	}
	_ = os.RemoveAll(dirPath)

	return codes.RequestedFileActionOkayCompleted, nil
}
