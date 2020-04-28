def output

pipeline {
   agent any

   tools {
        go {'gc-1.14'}
   }

   stages {
      stage('Hello') {
         steps {
             script {
                 output = sh label: '', returnStdout: true, script: 'echo "Hello, World"'
             }
         }
      }
      stage('Step two') {
         steps {
            echo output
         }
      }
      stage('Build') {
               steps {
                  sh 'go build'
              }
      }
   }
}

