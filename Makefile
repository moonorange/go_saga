# Define variables for image names and paths
BFF_IMAGE := bff:latest
QUERY_IMAGE := inventory_service:latest
COMMAND_IMAGE := payment_service:latest
DOCKERHUB_REPO := keigokida/gomicroservices

build:
	docker build --no-cache -t $(BFF_IMAGE) ./bff
	docker build --no-cache -t $(QUERY_IMAGE) ./microservices/inventory_service
	docker build --no-cache -t $(COMMAND_IMAGE) ./microservices/payment_service
	docker tag $(BFF_IMAGE) $(DOCKERHUB_REPO):bff
	docker tag $(QUERY_IMAGE) $(DOCKERHUB_REPO):inventory_service
	docker tag $(COMMAND_IMAGE) $(DOCKERHUB_REPO):payment_service

push_images:
	docker push $(DOCKERHUB_REPO):bff
	docker push $(DOCKERHUB_REPO):inventory_service
	docker push $(DOCKERHUB_REPO):payment_service

build_push: build push_images

minikube:
	minikube start

# Need to start minikube before deploying services
helm_install:
	helm install -f k8s/bff.yaml bff ./k8s/microservice
	helm install -f k8s/payment-service.yaml payment-service ./k8s/microservice
	helm install -f k8s/inventory-service.yaml inventory-service ./k8s/microservice

helm_uninstall:
	helm uninstall bff
	helm uninstall payment-service
	helm uninstall inventory-service

helm_update: helm_uninstall helm_install

.PHONY: build push_images build_push minikube helm_install helm_uninstall helm_update
