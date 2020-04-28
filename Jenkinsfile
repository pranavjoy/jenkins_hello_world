def output

pipeline {
   agent any

   tools {
        go {'gc-1.14'}
   }

   stages {
      stage('Build') {
          steps {
              sh 'go build'
          }
      }
   }
}

