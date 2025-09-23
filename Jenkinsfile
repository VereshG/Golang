pipeline {
    agent any

    environment {
        SLACK_CHANNEL = '#your-channel'
        SLACK_COLOR = 'good'
        GITHUB_REPO = 'VereshG/Golang' // Replace with your actual repo
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building after PR merge...'
                // Your build steps here
            }
        }
    }

    post {
        success {
            script {
                // Extract PR number from merge commit message
                def prNumber = sh(script: "git log -1 --pretty=format:'%s' | grep -oE '#[0-9]+' | tr -d '#' || echo 'N/A'", returnStdout: true).trim()
                
                // Extract author of the merge commit
                def prAuthor = sh(script: "git log -1 --pretty=format:'%an'", returnStdout: true).trim()

                // Construct PR link
                def prLink = "https://github.com/${env.GITHUB_REPO}/pull/${prNumber}"

                slackSend (
                    channel: env.SLACK_CHANNEL,
                    color: env.SLACK_COLOR,
                    message: "✅ PR #${prNumber} was merged by *${prAuthor}*\n🔗 ${prLink}"
                )
            }
        }
        failure {
            slackSend (
                channel: env.SLACK_CHANNEL,
                color: 'danger',
                message: "❌ Build failed after PR merge on branch ${env.GIT_BRANCH}"
            )
        }
    }
}