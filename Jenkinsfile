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
                echo 'Fetching latest main branch...'
                sh 'git fetch origin main'
                echo 'Getting previous and current commit SHAs...'
                script {
                    def previousCommit = sh(script: "git rev-parse HEAD^1", returnStdout: true).trim()
                    def currentCommit = sh(script: "git rev-parse HEAD", returnStdout: true).trim()
                    def changedFiles = sh(script: "git diff --name-only ${previousCommit} ${currentCommit}", returnStdout: true).trim().split('\n')
                    echo "Changed files: ${changedFiles}"
                }
            }
        }
    }

    post {
        always {
            script {
                // Shared logic for both success and failure
                def previousCommit = sh(script: "git rev-parse HEAD^1", returnStdout: true).trim()
                def currentCommit = sh(script: "git rev-parse HEAD", returnStdout: true).trim()
                def changedFiles = sh(script: "git diff --name-only ${previousCommit} ${currentCommit}", returnStdout: true).trim().split('\n')
                def prNumber = sh(script: "git log -1 --pretty=format:'%s' | grep -oE 'pull request #[0-9]+' | grep -oE '[0-9]+' || echo 'N/A'", returnStdout: true).trim()
                def prAuthor = sh(script: "git log -1 --pretty=format:'%an'", returnStdout: true).trim()
                def prLink = prNumber != 'N/A' ? "https://github.com/${env.GITHUB_REPO}/pull/${prNumber}" : ''
                def memberCoreChannel = "C09G161KD0Q"
                def memberFundsChannel = "C09F8HM77L6"

                // Expose for later blocks
                env.CHANGED_FILES = changedFiles.join(',')
                env.PR_NUMBER = prNumber
                env.PR_AUTHOR = prAuthor
                env.PR_LINK = prLink
            }
        }
        success {
            script {
                def memberCoreChannel = "C09G161KD0Q"
                def memberFundsChannel = "C09F8HM77L6"
                if (env.GIT_BRANCH == 'main' || env.BRANCH_NAME == 'main') {
                    def changedFiles = env.CHANGED_FILES.split(',')
                    def prNumber = env.PR_NUMBER
                    def prAuthor = env.PR_AUTHOR
                    def prLink = env.PR_LINK
                    def sent = false
                    if (changedFiles.contains('api/get_handler.go')) {
                        def appName = 'membercore'
                        def channelID = memberCoreChannel
                        echo "App Name: ${appName}"
                        echo "Notification sent to channel: ${channelID} for app: ${appName}"
                        def message = """
*✅ PR #${prNumber} merged by ${prAuthor}*
${prLink != '' ? "🔗 <${prLink}|View PR>\n" : ''}
*Changed files:*
${changedFiles.join('\n')}
*API changed:* ${appName}
Please review!
"""
                        echo "Sending Slack notification to ${channelID} with message: ${message}"
                        sh """
                        curl -X POST \
                            -H "Authorization: Bearer ${env.SLACK_TOKEN}" \
                            -H "Content-Type: application/json" \
                            -d '{"channel": "${channelID}", "text": "${message}"}' \
                            https://slack.com/api/chat.postMessage || echo "Slack notification failed"
                        """
                        sent = true
                    }
                    if (changedFiles.contains('api/post_handler.go')) {
                        def appName = 'member funds'
                        def channelID = memberFundsChannel
                        echo "App Name: ${appName}"
                        echo "Notification sent to channel: ${channelID} for app: ${appName}"
                        def message = """
*✅ PR #${prNumber} merged by ${prAuthor}*
${prLink != '' ? "🔗 <${prLink}|View PR>\n" : ''}
*Changed files:*
${changedFiles.join('\n')}
*API changed:* ${appName}
Please review!
"""
                        echo "Sending Slack notification to ${channelID} with message: ${message}"
                        sh """
                        curl -X POST \
                            -H "Authorization: Bearer ${env.SLACK_TOKEN}" \
                            -H "Content-Type: application/json" \
                            -d '{"channel": "${channelID}", "text": "${message}"}' \
                            https://slack.com/api/chat.postMessage || echo "Slack notification failed"
                        """
                        sent = true
                    }
                    if (!sent) {
                        echo "No relevant API file changed. No notification sent."
                    }
                } else {
                    echo "PR was not merged to main branch. No notifications sent."
                }
            }
        }
        failure {
            script {
                def memberCoreChannel = "C09G161KD0Q"
                def memberFundsChannel = "C09F8HM77L6"
                if (env.GIT_BRANCH == 'main' || env.BRANCH_NAME == 'main') {
                    def message = "❌ Build failed after PR merge on branch ${env.GIT_BRANCH}"
                    echo "Build failed. Notification sent to both channels."
                    echo "Notification sent to channel: ${memberCoreChannel}"
                    echo "Notification sent to channel: ${memberFundsChannel}"
                    // Notify member core team
                    echo "Sending Slack notification to ${memberCoreChannel} with message: ${message}"
                    sh """
                    curl -X POST \
                        -H "Authorization: Bearer ${env.SLACK_TOKEN}" \
                        -H "Content-Type: application/json" \
                        -d '{"channel": "${memberCoreChannel}", "text": "${message}"}' \
                        https://slack.com/api/chat.postMessage || echo "Slack notification failed"
                    """
                    // Notify member funds team
                    echo "Sending Slack notification to ${memberFundsChannel} with message: ${message}"
                    sh """
                    curl -X POST \
                        -H "Authorization: Bearer ${env.SLACK_TOKEN}" \
                        -H "Content-Type: application/json" \
                        -d '{"channel": "${memberFundsChannel}", "text": "${message}"}' \
                        https://slack.com/api/chat.postMessage || echo "Slack notification failed"
                    """
                } else {
                    echo "Build failed, but not on main branch. No notifications sent."
                }
            }
        }
    }
}