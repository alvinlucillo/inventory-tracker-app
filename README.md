# inventory-tracker-app

Build Docker image
docker build -f ./infra/dockerfiles/frontend.prod.dockerfile -t alvinlucillo/inventorytrackerapp_frontend .

kubectl apply -f ./infra/k8s/.
minikube start
