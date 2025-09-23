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
                echo 'Changed files in this PR:'
                sh 'git diff --name-only origin/main...HEAD'
            }
        }
    }

    post {
        success {
            script {
                if (env.GIT_BRANCH == 'main' || env.BRANCH_NAME == 'main') {
                    def changedFiles = sh(script: "git diff --name-only origin/main...HEAD", returnStdout: true).trim().split('\n')
                    def prNumber = sh(script: "git log -1 --pretty=format:'%s' | grep -oE 'pull request #[0-9]+' | grep -oE '[0-9]+' || echo 'N/A'", returnStdout: true).trim()
                    def prAuthor = sh(script: "git log -1 --pretty=format:'%an'", returnStdout: true).trim()
                    def prLink = "https://github.com/${env.GITHUB_REPO}/pull/${prNumber}"
                    def memberCoreChannel = "C09G161KD0Q"   // Channel ID for member core
                    def memberFundsChannel = "C09F8HM77L6" // Channel ID for member funds

                    echo "Changed files: ${changedFiles}"
                    echo "PR Link: ${prLink}"
                    echo "PR Author: ${prAuthor}"

                    // Call Go script to get app name (simulate output)
                    def appName = ''
                    if (changedFiles.contains('api/get_handler.go')) {
                        appName = 'membercore'
                        echo "App Name: ${appName}"
                        echo "Notification sent to channel: ${memberCoreChannel} for app: ${appName}"
                        def message = "✅ PR #${prNumber} was merged by *${prAuthor}*\n🔗 ${prLink}\nChanged files: ${changedFiles}\nGET API changed, member core team please review!"
                        sh """
                        curl -X POST \
                            -H "Authorization: Bearer ${SLACK_TOKEN}" \
                            -H "Content-Type: application/json" \
                            -d '{"channel": "${memberCoreChannel}", "text": "${message}"}' \
                            https://slack.com/api/chat.postMessage
                        """
                    }
                    if (changedFiles.contains('api/post_handler.go')) {
                        appName = 'member funds'
                        echo "App Name: ${appName}"
                        echo "Notification sent to channel: ${memberFundsChannel} for app: ${appName}"
                        def message = "✅ PR #${prNumber} was merged by *${prAuthor}*\n🔗 ${prLink}\nChanged files: ${changedFiles}\nPOST API changed, member funds team please review!"
                        sh """
                        curl -X POST \
                            -H "Authorization: Bearer ${SLACK_TOKEN}" \
                            -H "Content-Type: application/json" \
                            -d '{"channel": "${memberFundsChannel}", "text": "${message}"}' \
                            https://slack.com/api/chat.postMessage
                        """
                    }
                } else {
                    echo "PR was not merged to main branch. No notifications sent."
                }
            }
        }

        failure {
            script {
                if (env.GIT_BRANCH == 'main' || env.BRANCH_NAME == 'main') {
                    def message = "❌ Build failed after PR merge on branch ${env.GIT_BRANCH}"
                    def memberCoreChannel = "C09G161KD0Q"
                    def memberFundsChannel = "C09F8HM77L6"

                    echo "Build failed. Notification sent to both channels."
                    echo "Notification sent to channel: ${memberCoreChannel}"
                    echo "Notification sent to channel: ${memberFundsChannel}"

                    // Notify member core team
                    sh """
                    curl -X POST \
                        -H "Authorization: Bearer ${SLACK_TOKEN}" \
                        -H "Content-Type: application/json" \
                        -d '{"channel": "${memberCoreChannel}", "text": "${message}"}' \
                        https://slack.com/api/chat.postMessage
                    """

                    // Notify member funds team
                    sh """
                    curl -X POST \
                        -H "Authorization: Bearer ${SLACK_TOKEN}" \
                        -H "Content-Type: application/json" \
                        -d '{"channel": "${memberFundsChannel}", "text": "${message}"}' \
                        https://slack.com/api/chat.postMessage
                    """
                } else {
                    echo "Build failed, but not on main branch. No notifications sent."
                }
            }
        }
    }
}