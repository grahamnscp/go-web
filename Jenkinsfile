# Jenkinsfile for go-web

env.DOCKERHUB_USERNAME = 'grahamnscp'
  
// this will start an executor on a Jenkins agent with label docker-test

node("docker-test") {

  // Setup variables
  String APP_NAME = "go-web"
  //String THIS_BUILD_NUMBER = "0.1.${env.BUILD_NUMBER}"
  //String goPath = "/go/src/${DOCKERHUB_USERNAME}/${APP_NAME}"

  checkout scm

  stage("Test") {

    try {

      # build image
      sh "docker build -t ${APP_NAME} ."

      # clean local source
      sh "docker rm -f ${APP_NAME} || true"

      # run container instance of app
      sh "docker run -d -p 8080:8080 --name=${APP_NAME} ${DOCKERHUB_USERNAME}/${APP_NAME}:${BUILD_NUMBER}"

      # test
      sh "docker run --rm -v ${WORKSPACE}:/go/src/${APP_NAME} --link=${APP_NAME} -e SERVER=${APP_NAME} golang go test ${APP_NAME} -v --run Unit"

    } catch(e) {

      error "Test failed"

    } finally {

      # test done, clean up
      sh "docker rm -f ${APP_NAME} || true"
      sh "docker ps -aq | xargs docker rm || true"
      sh "docker images -aq -f dangling=true | xargs docker rmi || true"
    }
  }
}

