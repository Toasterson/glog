pipeline {
  agent {
    node {
      label 'buildserver'
    }

  }
  stages {
    stage('Build') {
      steps {
        echo 'Testing the Pipeline editor'
        isUnix()
        sh 'go test -v ./...'
      }
    }

  }
}