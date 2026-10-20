package task

import "errors"

var (
	ErrTaskNotFount   = errors.New("task not found")
	ErrTaskReachLimit = errors.New("tasks if limit")
)
