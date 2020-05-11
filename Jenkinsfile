pipeline {
   agent any

   tools {
        go {'gc-1.14'}
   }

   stages {
      stage('Unit Test') {
          steps {
              sh 'go test'
          }
      }
      
      stage('Build') {
          steps {
              sh 'go build -o example1'
          }
      }
      
      stage('Deliver') {
          steps {
                 ansiblePlaybook disableHostKeyChecking: true, credentialsId: 'toobox-vagrant-key', inventory: 'hosts.ini', playbook: 'playbook.yml'
          }
      }

      stage('Integration Test') {
          steps {
             sh 'docker run -t postman/newman:latest run "https://www.getpostman.com/collections/7a43d59097dbfeac55f8"'
          }
      }
   }
}

