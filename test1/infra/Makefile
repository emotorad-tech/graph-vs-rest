GRAPHQL_TEST1_SERVER=graphqlTest1Server
REST_TEST1_SERVER=restTest1Server

build: build_graphql build_rest

deploy: deploy_graphql deploy_rest

build_graphql:
	@echo "Building graphql test 1 server binary..."
	cd ../graphql && env GOOS=linux CGO_ENABLED=0 go build -o ${GRAPHQL_TEST1_SERVER} server.go
	@echo "Building Docker image..."
	cd ../graphql && docker build -t 179172241320.dkr.ecr.ap-south-1.amazonaws.com/graphql-test1 -t 179172241320.dkr.ecr.ap-south-1.amazonaws.com/graphql-test1:latest .
	@echo "Done!"

deploy_graphql: build_graphql
	@echo "Logging in to AWS ECR"
	aws ecr get-login-password --region ap-south-1 | docker login --username AWS --password-stdin 179172241320.dkr.ecr.ap-south-1.amazonaws.com
	@echo "Pushing graphql test 1 to AWS ECR"
	docker push -a 179172241320.dkr.ecr.ap-south-1.amazonaws.com/graphql-test1
	@echo "Removing graphql test 1 deployment"
	kubectl delete deployment graphql-test1
	@echo "Deploying graphql test 1"
	kubectl apply -f k8s/graphql-test1.yaml

build_rest:
	@echo "Building rest test 1 server binary..."
	cd ../rest && env GOOS=linux CGO_ENABLED=0 go build -o ${REST_TEST1_SERVER} ./main
	@echo "Building Docker image..."
	cd ../rest && docker build -t 179172241320.dkr.ecr.ap-south-1.amazonaws.com/rest-test1 -t 179172241320.dkr.ecr.ap-south-1.amazonaws.com/rest-test1:latest .
	@echo "Done!"

deploy_rest: build_rest
	@echo "Logging in to AWS ECR"
	aws ecr get-login-password --region ap-south-1 | docker login --username AWS --password-stdin 179172241320.dkr.ecr.ap-south-1.amazonaws.com
	@echo "Pushing rest test 1 to AWS ECR"
	docker push -a 179172241320.dkr.ecr.ap-south-1.amazonaws.com/rest-test1
	@echo "Removing rest test 1 deployment"
	kubectl delete deployment rest-test1
	@echo "Deploying rest test 1"
	kubectl apply -f k8s/rest-test1.yaml