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
                 ansiblePlaybook credentialsId: 'toobox-vagrant-key', inventory: 'hosts.ini', playbook: 'playbook.yml'
          }
      }
   }
}

