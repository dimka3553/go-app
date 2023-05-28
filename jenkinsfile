pipeline {
    agent any
    tools {
        go 'go1.20.4'
        git 'Default'
    }

    stages {
        stage('Checkout') {
            steps {
                echo 'Cloning repo'
                git branch: 'main', url: 'https://github.com/dimka3553/go-app.git'
            }
        }
        stage('Build') {
            steps {
                echo 'Building'
                sh 'go build -o main'
                sh 'ls'
            }
        }

        stage('Test') {
            steps {
                echo 'Running tests'
                sh 'go test'
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying...'
                sh "ssh -i ~/.ssh/id_rsa vagrant@192.168.105.3 'if systemctl --all --state=running | grep -q main; then sudo systemctl stop main; fi'"
                sh "scp -i ~/.ssh/id_rsa main vagrant@192.168.105.3:/home/vagrant"
                sh "ssh -i ~/.ssh/id_rsa vagrant@192.168.105.3 'chmod +x /home/vagrant/main'"
                sh 'ssh -i ~/.ssh/id_rsa vagrant@192.168.105.3 "echo \'[Unit]\nDescription=My Go App\nAfter=network.target\n\n[Service]\nExecStart=/home/vagrant/main\nUser=vagrant\nRestart=always\nRestartSec=5\n\n[Install]\nWantedBy=multi-user.target\' > main.service"'
                sh "ssh -i ~/.ssh/id_rsa vagrant@192.168.105.3 'sudo mv main.service /etc/systemd/system/main.service'"
                sh "ssh -i ~/.ssh/id_rsa vagrant@192.168.105.3 'sudo systemctl daemon-reload'"
                sh "ssh -i ~/.ssh/id_rsa vagrant@192.168.105.3 'sudo systemctl enable main'"
                sh "ssh -i ~/.ssh/id_rsa vagrant@192.168.105.3 'sudo systemctl start main'"
            }
        }
    }

    post {
        failure {
            echo 'There were some failures...'
        }
        success {
            echo 'All stages completed successfully!'
            archiveArtifacts artifacts: 'main', fingerprint: true
        }
    }
}
