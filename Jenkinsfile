pipeline {
    agent any

    environment {
        SLACK_CHANNEL_ID = 'C09F8HM77L6'
        SLACK_TOKEN = credentials('SLACK_BOT_TOKEN') // Jenkins credential ID for your xoxb token
        GITHUB_REPO = 'VereshG/Golang'
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
                        memberCoreChannel = "YOUR_MEMBERCORE_CHANNEL_ID"
                        memberFundsChannel = "YOUR_MEMBERFUNDS_CHANNEL_ID"
                        generalChannel = "YOUR_GENERAL_CHANNEL_ID" // <-- Add your general Slack channel ID here
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
                        // Send notification if PR is raised to release branch (not main)
                        if (env.GIT_BRANCH == 'release' || env.BRANCH_NAME == 'release') {
                    def changedFiles = env.CHANGED_FILES.split(',')
                    def prNumber = env.PR_NUMBER
                    def prAuthor = env.PR_AUTHOR
                    def prLink = env.PR_LINK
                    def sent = false
                    // Only notify relevant team if endpoint file changed
                    // Define onlyGetChanged and onlyPostChanged only once in this block
                    def onlyGetChanged = changedFiles.every { it == 'api/get_handler.go' }
                    def onlyPostChanged = changedFiles.every { it == 'api/post_handler.go' }
                    if (onlyGetChanged) {
                        def appName = 'GET endpoint'
                        def channelID = memberFundsChannel // Notify funds team when GET endpoint is changed
                        echo "App Name: ${appName}"
                        echo "Notification sent to channel: ${channelID} for app: ${appName}"
                        def message = "*✅ PR #${prNumber} merged by ${prAuthor}*\n" +
                            (prLink != '' ? "🔗 <${prLink}|View PR>\n" : '') +
                            "*Changed files:*\n" + changedFiles.join('\n') + "\n" +
                            "*API changed:* ${appName}\n" +
                            "*Note: This endpoint is owned by the core team. Funds team is being notified of changes.*\n" +
                            "Please review!"
                        echo "Sending Slack notification to ${channelID} with message: ${message}"
                        sh """
                        curl -X POST \
                            -H "Authorization: Bearer ${env.SLACK_TOKEN}" \
                            -H "Content-Type: application/json" \
                            -d '{"channel": "${channelID}", "text": "${message}"}' \
                            https://slack.com/api/chat.postMessage || echo "Slack notification failed"
                        """
                    } else if (onlyPostChanged) {
                        def appName = 'POST endpoint'
                        def channelID = memberCoreChannel // Notify core team when POST endpoint is changed
                        def message = """
*✅ PR #${prNumber} merged by ${prAuthor}*
${prLink != '' ? "🔗 <${prLink}|View PR>\n" : ''}
*Changed files:*
${changedFiles.join('\n')}
*API changed:* ${appName}
*Note: This endpoint is owned by the funds team. Core team is being notified of changes.*
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
                    } else if (changedFiles.size() > 0) {
                        // If any other file (or both endpoints) changed, notify both teams with full details
                        def apiChanged = ''
                        def teamNote = ''
                        if (changedFiles.contains('api/get_handler.go') && changedFiles.contains('api/post_handler.go')) {
                            apiChanged = 'Both endpoints changed'
                            teamNote = 'Both teams are being notified of changes.'
                        } else if (changedFiles.contains('api/get_handler.go')) {
                            apiChanged = 'GET endpoint changed'
                            teamNote = 'Funds team and core team are being notified of changes.'
                        } else if (changedFiles.contains('api/post_handler.go')) {
                            apiChanged = 'POST endpoint changed'
                            teamNote = 'Core team and funds team are being notified of changes.'
                        } else {
                            apiChanged = 'Non-endpoint files changed'
                            teamNote = 'Both teams are being notified of changes.'
                        }
                        def message = """
*✅ PR #${prNumber} merged by ${prAuthor}*
                            def endpointInfo = changedFiles.contains('api/get_handler.go') ? 'GET endpoint (owned by Core Team)' : 'No endpoint files changed'
                            endpointInfo += changedFiles.contains('api/post_handler.go') ? ', POST endpoint (owned by Funds Team)' : ''
                            def teamNames = 'Core Team (C09G161KD0Q) and Funds Team (C09F8HM77L6)'
                            def message = """
*Changed files:*
${changedFiles.join('\n')}
*API changed:* ${apiChanged}
*Note: ${teamNote}*
Please review!
"""
                        echo "Sending Slack notification to ${memberCoreChannel} and ${memberFundsChannel} for non-endpoint or multiple endpoint file changes: ${changedFiles}"
                        // Notify memberCoreChannel
                        sh """
                        curl -X POST \
                            -H "Authorization: Bearer ${env.SLACK_TOKEN}" \
                            -H "Content-Type: application/json" \
                            -d '{"channel": "${memberCoreChannel}", "text": "${message}"}' \
                            https://slack.com/api/chat.postMessage || echo "Slack notification failed"
                        """
                        // Notify memberFundsChannel
                        sh """
                        curl -X POST \
                            -H "Authorization: Bearer ${env.SLACK_TOKEN}" \
                            -H "Content-Type: application/json" \
                            -d '{"channel": "${memberFundsChannel}", "text": "${message}"}' \
                            https://slack.com/api/chat.postMessage || echo "Slack notification failed"
                        """
                    } else {
                        echo "No relevant file changed. No notification sent."
                    }
                        } else {
                            echo "PR was not merged to release branch. No notifications sent."
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