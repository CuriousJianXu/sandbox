.PHONY: sqlc-gen

sqlc-gen:
	# sadly not work for windows path. Just paste it in command line.
	`docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate`

tmp:
	docker run --rm -v ${PWD}:`"/src"` -w `"/src"` sqlc/sqlc generate