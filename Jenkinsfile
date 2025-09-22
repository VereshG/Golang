pipeline {
    agent any
    triggers {
        githubPullRequest()
    }
    stages {
        stage('Build') {
            steps {
                sh 'cd api && go build -v ./...'
            }
        }
        stage('Test') {
            steps {
                sh 'cd api && go test ./...'
            }
        }
    }
}
