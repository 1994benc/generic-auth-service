pipeline {
    agent any
    environment {
            ENV_FILE = credentials('env')
    }
    stages {
        stage("build & run") {
            steps {
                echo "Spinning up the docker containers..."
                sh '/usr/local/bin/docker-compose --env-file $ENV_FILE up -d --build'
            }
        }
    }
}