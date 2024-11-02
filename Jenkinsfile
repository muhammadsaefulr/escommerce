pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/muhammadsaefulr/escommerce.git' 
            }
        }
        stage('Builds Docker Compose') {
            steps {
                script {
                    sh 'docker-compose -f docker-compose.yml build'
                }
            }
        }
        stage('Stop Existing Service') {
            steps {
                script {
                    sh '''
                    docker-compose -f docker-compose.yml down
                    '''
                }
            }
        }
        stage('Deploy New Service') {
            steps {
                script {
                    sh '''
                    docker-compose -f docker-compose.yml up -d
                    '''
                }
            }
        }
    }
    post {
        success {
            echo 'Build and Deployment Successful!'
        }
        failure {
            echo 'Build or Deployment Failed!'
        }
    }
}
