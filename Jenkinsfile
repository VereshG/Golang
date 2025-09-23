pipeline {
    agent any

    environment {
        SLACK_CHANNEL = '#gitping'
        SLACK_COLOR = 'good'
        GITHUB_REPO = 'VereshG/Golang'
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
                def prNumber = sh(script: "git log -1 --pretty=format:'%s' | grep -oE 'pull request #[0-9]+' | grep -oE '[0-9]+' || echo 'N/A'", returnStdout: true).trim()
                
                // Extract author of the merge commit
                def prAuthor = sh(script: "git log -1 --pretty=format:'%an'", returnStdout: true).trim()

                // Construct PR link
                def prLink = "https://github.com/${env.GITHUB_REPO}/pull/${prNumber}"

                slackSend (
                    channel: env.SLACK_CHANNEL,
                    color: env.SLACK_COLOR,
                    tokenCredentialId: 'Slack_Token',
                    message: "✅ PR #${prNumber} was merged by *${prAuthor}*\n🔗 ${prLink}"
                )
            }
        }
        failure {
            slackSend (
                channel: env.SLACK_CHANNEL,
                color: 'danger',
                tokenCredentialId: 'Slack_Token',
                message: "❌ Build failed after PR merge on branch ${env.GIT_BRANCH}"
            )
        }
    }
}