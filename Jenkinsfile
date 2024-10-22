pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'escommerce-api-web'
    }

    stages {
        stage('Inject Env Variables') {
            steps {
                script {
                    def props = readProperties file: '.env'
                    env.DB_HOST = props['DB_HOST']
                    env.DB_USER = props['DB_USER']
                    env.DB_PASSWORD = props['DB_PASSWORD']
                    env.DB_NAME = props['DB_NAME']
                    env.DB_PORT = props['DB_PORT']
                }
            }
        }

        stage('Checkout') {
            steps {
                git 'https://github.com/muhammadsaefulr/escommerce-api'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build(DOCKER_IMAGE)
                }
            }
        }

        stage('Run Docker Container') {
            steps {
                script {
                    docker.image(DOCKER_IMAGE).run("-e DB_HOST=${env.DB_HOST} -e DB_USER=${env.DB_USER} -e DB_PASSWORD=${env.DB_PASSWORD} -e DB_NAME=${env.DB_NAME} -e DB_PORT=${env.DB_PORT} -p 3000:3000")
                }
            }
        }

        stage('Test') {
            steps {
                echo 'Unitest Not Ready For Now, Skipping...'
            }
        }

        stage('Deploy') {
            steps {
                script {
                    docker.withRegistry(DOCKER_REGISTRY, 'docker-credentials') {
                        docker.image(DOCKER_IMAGE).push('latest')
                    }
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
