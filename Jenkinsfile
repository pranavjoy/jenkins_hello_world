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
             sh 'sudo docker run -t postman/newman:latest run "https://www.getpostman.com/collections/0fb9d09f41531188decc"'
          }
      }
   }
}