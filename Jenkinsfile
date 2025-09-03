pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                git  branch: 'main', url:'https://github.com/SakshiTathe/devoproject.git'
            }
        }
        stage('Build Docker Image') {
            steps {
                sh 'docker build -t sakshi2004docker/goapp .'
            }
        }
        stage('Push Image') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-creds', usernameVariable: 'USER', passwordVariable: 'PASS')]) {
                    sh 'echo "Testing login with user: $USER"'
                    sh 'echo $PASS | docker login -u $USER --password-stdin'
                    sh 'docker push sakshi2004docker/goapp:latest'
                }
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                withCredentials([file(credentialsId: 'jenkins-kubeconfig', variable: 'KUBECONFIG')]) {
                sh '''
                 export KUBECONFIG=$KUBECONFIG
                helm upgrade --install goapp ./goapp-chart \
                  --namespace default --create-namespace \
                  --set image.repository=sakshi2004docker/goapp \
                  --set image.tag=latest
                kubectl rollout status deployment/goapp -n default
                '''
                }
            }
        }
    }
}

