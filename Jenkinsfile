pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/SakshiTathe/devoproject.git'
            }
        }
        stage('Build Docker Image') {
            steps {
                sh 'docker build -t sakshi2004docker/goapp .'
            }
        }
        stage('Push Image') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'USER', passwordVariable: 'PASS')]) {
                    sh 'echo $PASS | docker login -u $USER --password-stdin'
                    sh 'docker push sakshi2004docker/goapp:latest'
                }
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                sh 'kubectl set image deployment/go-app go-app=sakshi2004docker/goapp:latest'
            }
        }
    }
}

