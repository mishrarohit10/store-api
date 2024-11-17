pipeline {
    agent any

    environment {
        REGISTRY = "docker.io" 
        IMAGE_NAME = "brucexwayne/libman" 
        DOCKERHUB_CREDENTIAL = 'docker-credentials-id'
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    echo "Checking out code from the repository..."
                }
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    echo "Building Docker image: ${REGISTRY}/${IMAGE_NAME}:${env.BUILD_ID}"
                }
                script {
                    docker.build("${REGISTRY}/${IMAGE_NAME}:${env.BUILD_ID}")
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    echo "Pushing Docker image: ${REGISTRY}/${IMAGE_NAME}:${env.BUILD_ID}"
                    docker.withRegistry("https://${REGISTRY}", DOCKERHUB_CREDENTIAL) {
                        echo "Authenticating with Docker Hub..."
                        docker.image("${REGISTRY}/${IMAGE_NAME}:${env.BUILD_ID}").push()
                        echo "Image pushed successfully to Docker Hub."
                    }
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                script {
                    echo "Deploying to Kubernetes..."
                }
                sh """
                kubectl set image deployment/myapp-deployment myapp=${REGISTRY}/${IMAGE_NAME}:${env.BUILD_ID} -n default
                kubectl rollout status deployment/myapp-deployment -n default
                """
                script {
                    echo "Deployment completed successfully."
                }
            }
        }
    }
    
    post {
        always {
            echo "Cleaning up workspace..."
            cleanWs()
            echo "Workspace cleaned."
        }
    }
}
