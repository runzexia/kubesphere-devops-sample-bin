pipeline {
  agent {
    node {
      label 'base'
    }
  }


  stages {
    stage('checkout scm') {
      steps {
        checkout(scm)

      }
    }
    stage('get dependencies') {
      steps {
        container('base') {
         input(id: 'input', message: 'input?')
        }

      }
    }
  }

}
