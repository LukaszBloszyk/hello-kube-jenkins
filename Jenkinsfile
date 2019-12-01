pipeline {
  agent any
  stages {
    stage('pull') {
      steps {
        echo 'hello from hello-kube-jenkins'
      }
    }

    stage('build') {
      steps {
        sh 'docker build -t lukaszbloszyk/hello-go-demo:latest -t lukaszbloszyk/hello-go-demo:v1.1 .'
      }
    }

    stage('push') {
      steps {
        sh 'docker login -u $DOCKER_HUB_USERNAME -p $DOCKER_HUB_PASSWORD'
        sh '#docker push lukaszbloszyk/hello-go-demo'
        sh '''curl -ik \\
     -H "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \\
     https://kubernetes.default.svc.cluster.local/api/v1/namespaces/default/pods'''
        sh 'kubectl get svc'
      }
    }

  }
}