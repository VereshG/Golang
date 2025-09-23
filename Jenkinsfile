pipeline {
    agent any

    environment {
        SLACK_CHANNEL_ID = 'C09F8HM77L6'
        SLACK_TOKEN = credentials('Slack_Token') // Jenkins credential ID for your xoxb token
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
                def prNumber = sh(script: "git log -1 --pretty=format:'%s' | grep -oE 'pull request #[0-9]+' | grep -oE '[0-9]+' || echo 'N/A'", returnStdout: true).trim()
                def prAuthor = sh(script: "git log -1 --pretty=format:'%an'", returnStdout: true).trim()
                def prLink = "https://github.com/${env.GITHUB_REPO}/pull/${prNumber}"
                def message = "✅ PR #${prNumber} was merged by *${prAuthor}*\n🔗 ${prLink}"

                sh """
                curl -X POST \
                  -H "Authorization: Bearer ${SLACK_TOKEN}" \
                  -H "Content-Type: application/json" \
                  -d '{
                        "channel": "${SLACK_CHANNEL_ID}",
                        "text": "${message}"
                      }' \
                  https://slack.com/api/chat.postMessage
                """
            }
        }

        failure {
            script {
                def message = "❌ Build failed after PR merge on branch ${env.GIT_BRANCH}"

                sh """
                curl -X POST \
                  -H "Authorization: Bearer ${SLACK_TOKEN}" \
                  -H "Content-Type: application/json" \
                  -d '{
                        "channel": "${SLACK_CHANNEL_ID}",
                        "text": "${message}"
                      }' \
                  https://slack.com/api/chat.postMessage
                """
            }
        }
    }
}