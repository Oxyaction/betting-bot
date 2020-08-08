.PHONY: start

cmd_dir = cmd

start:
	cd $(cmd_dir) && API_TOKEN=$(API_TOKEN) go run .