GRAPHQL_TEST2_SERVER=graphqlTest2Server
REST_TEST2_SERVER=restTest2Server

build: build_graphql build_rest

deploy: deploy_graphql deploy_rest

build_graphql:
	@echo "Building graphql test 2 server binary..."
	cd ../graphql && env GOOS=linux CGO_ENABLED=0 go build -o ${GRAPHQL_TEST2_SERVER} ./main
	@echo "Building Docker image..."
	cd ../graphql && docker build -t 179172241320.dkr.ecr.ap-south-1.amazonaws.com/graphql-test2 -t 179172241320.dkr.ecr.ap-south-1.amazonaws.com/graphql-test2:latest .
	@echo "Done!"

deploy_graphql: build_graphql
	@echo "Logging in to AWS ECR"
	aws ecr get-login-password --region ap-south-1 | docker login --username AWS --password-stdin 179172241320.dkr.ecr.ap-south-1.amazonaws.com
	@echo "Pushing graphql test 2 to AWS ECR"
	docker push -a 179172241320.dkr.ecr.ap-south-1.amazonaws.com/graphql-test2
	@echo "Removing graphql test 2 deployment"
	kubectl delete deployment graphql-test2
	@echo "Deploying graphql test 2"
	kubectl apply -f k8s/graphql-test2.yaml

build_rest:
	@echo "Building rest test 2 server binary..."
	cd ../rest && env GOOS=linux CGO_ENABLED=0 go build -o ${REST_TEST2_SERVER} ./main
	@echo "Building Docker image..."
	cd ../rest && docker build -t 179172241320.dkr.ecr.ap-south-1.amazonaws.com/rest-test2 -t 179172241320.dkr.ecr.ap-south-1.amazonaws.com/rest-test2:latest .
	@echo "Done!"

deploy_rest: build_rest
	@echo "Logging in to AWS ECR"
	aws ecr get-login-password --region ap-south-1 | docker login --username AWS --password-stdin 179172241320.dkr.ecr.ap-south-1.amazonaws.com
	@echo "Pushing rest test 2 to AWS ECR"
	docker push -a 179172241320.dkr.ecr.ap-south-1.amazonaws.com/rest-test2
	@echo "Removing rest test 2 deployment"
	kubectl delete deployment rest-test2
	@echo "Deploying rest test 2"
	kubectl apply -f k8s/rest-test2.yaml