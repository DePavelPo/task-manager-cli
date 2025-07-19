# task-manager-cli
A simple task manager with CLI

There are few commands available now:
 - add [task] : add new task (example: add "check forecast")
 - list : get a list of created tasks. 
 - - Flags: --pending for getting only non-completed tasks, --completed for getting only completed tasks (example: list --pending)
 - done [id] : mark task as completed (example: done 1)
 - delete [id] : delete task by id (example: delete 1)
