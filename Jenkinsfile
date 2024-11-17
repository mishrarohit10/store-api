pipeline {
    agent any

    environment {
        REGISTRY = "docker.io" 
        IMAGE_NAME = "brucexwayne/libman" 
        // KUBECONFIG_CREDENTIALS = 'kubeconfig-credentials-id'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${REGISTRY}/${IMAGE_NAME}:${env.BUILD_ID}")
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    docker.withRegistry("https://${REGISTRY}", "docker-credentials-id") {
                        docker.image("${REGISTRY}/${IMAGE_NAME}:${env.BUILD_ID}").push()
                    }
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh """
                kubectl set image deployment/myapp-deployment myapp=${REGISTRY}/${IMAGE_NAME}:${env.BUILD_ID} -n default
                kubectl rollout status deployment/myapp-deployment -n default
                """
            }
        }

    }
    
    post {
        always {
            cleanWs()
        }
    }
}
