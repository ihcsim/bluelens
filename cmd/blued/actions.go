package main

import "context"

// invoke executes the task function provided by the user controller. Both c and e are receiving channels used to send the results back to the user controller.
func invoke(ctx context.Context, task func()) {
	if ctx.Err() != nil {
		return
	}

	task()
}
