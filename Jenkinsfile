pipeline {
    agent any
    tools {
        go 'go1.20.4'
        git 'Default'
    }
    environment{
        STAGING_ID = "192.168.105.3"
        AWS_ID = "ec2-52-59-114-177.eu-central-1.compute.amazonaws.com"
    }

    stages {
        stage('Checkout') {
            steps {
                echo 'Cloning repo'
                git branch: 'main', url: 'https://github.com/dimka3553/go-app.git'
            }
        }
    

        stage('Test') {
            steps {
                echo 'Running tests'
                sh 'go test'
            }
        }
        
            stage('Build') {
            steps {
                echo 'Building'
                sh 'go build -o main'
                sh 'ls'
            }
        }
        
         stage('Deploy Staging') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'aws', keyFileVariable: 'key', usernameVariable: 'name')]) {
                    echo 'Deploying Staging...'
                    sh "ssh-keyscan ${STAGING_ID} >> ~/.ssh/known_hosts"
                    sh "ssh vagrant@${STAGING_ID} 'if systemctl --all --state=running | grep -q main; then sudo systemctl stop main; fi'"
                    sh "scp main vagrant@${STAGING_ID}:/home/vagrant"
                    sh "ssh vagrant@${STAGING_ID} 'chmod +x /home/vagrant/main'"
                    sh 'ssh vagrant@${STAGING_ID} "echo \'[Unit]\nDescription=My Go App\nAfter=network.target\n\n[Service]\nExecStart=/home/vagrant/main\nUser=vagrant\nRestart=always\nRestartSec=5\n\n[Install]\nWantedBy=multi-user.target\' > main.service"'
                    sh "ssh vagrant@${STAGING_ID} 'sudo mv main.service /etc/systemd/system/main.service'"
                    sh "ssh vagrant@${STAGING_ID} 'sudo systemctl daemon-reload'"
                    sh "ssh vagrant@${STAGING_ID} 'sudo systemctl enable main'"
                    sh "ssh vagrant@${STAGING_ID} 'sudo systemctl start main'"
                }
            }
        }

        stage('Deploy Prod') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'aws', keyFileVariable: 'key', usernameVariable: 'name')]) {
                    echo 'Deploying Prod...'
                    sh "ssh-keyscan ${AWS_ID} >> ~/.ssh/known_hosts"
                    sh "ssh -i ${key} ${name}@${AWS_ID} 'if systemctl --all --state=running | grep -q main; then sudo systemctl stop main; fi'"
                    sh "scp -i ${key} main ${name}@${AWS_ID}:/home/${name}"
                    sh "ssh -i ${key} ${name}@${AWS_ID} 'chmod +x /home/${name}/main'"
                    sh 'ssh -i ${key} ${name}@${AWS_ID} "echo \'[Unit]\nDescription=My Go App\nAfter=network.target\n\n[Service]\nExecStart=/home/${name}/main\nUser=${name}\nRestart=always\nRestartSec=5\n\n[Install]\nWantedBy=multi-user.target\' > main.service"'
                    sh "ssh -i ${key} ${name}@${AWS_ID} 'sudo mv main.service /etc/systemd/system/main.service'"
                    sh "ssh -i ${key} ${name}@${AWS_ID} 'sudo systemctl daemon-reload'"
                    sh "ssh -i ${key} ${name}@${AWS_ID} 'sudo systemctl enable main'"
                    sh "ssh -i ${key} ${name}@${AWS_ID} 'sudo systemctl start main'"
                }
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

