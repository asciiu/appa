downgrade-test:
	goose postgres "user=postgres password=postgres dbname=appa_test sslmode=disable" down 
downgrade-dev:
	goose postgres "user=postgres dbname=appa_dev sslmode=disable" down

redo-test:
	goose postgres "user=postgres dbname=appa_test sslmode=disable" reset 
	goose postgres "user=postgres dbname=appa_test sslmode=disable" up
redo-dev:
	goose postgres "user=postgres dbname=appa_dev sslmode=disable" reset 
	goose postgres "user=postgres dbname=appa_dev sslmode=disable" up 

upgrade-dev:
	goose postgres "user=postgres dbname=appa_dev sslmode=disable" up
upgrade-test:
	goose postgres "user=postgres password=postgres dbname=appa_test sslmode=disable" up