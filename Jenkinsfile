// Jenkinsfile for go-web

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

      // build image
      sh "/usr/local/bin/docker build -t ${DOCKERHUB_USERNAME}/${APP_NAME}:${BUILD_NUMBER} ."

      // clean any existing running container
      //sh "/usr/local/bin/docker rm -f ${APP_NAME} || true"

      // run container instance of app
      sh "/usr/local/bin/docker run -d -p 8080:8080 --name=${APP_NAME} ${DOCKERHUB_USERNAME}/${APP_NAME}:${BUILD_NUMBER}"

      // test
      sh "/usr/local/bin/docker run --rm --env HOSTIP=`/sbin/ip route|awk '/default/ { print  $3}'` -v ${WORKSPACE}/src/${APP_NAME}:/go/src/${APP_NAME} --link=${APP_NAME} -e SERVER=${APP_NAME} golang sed -i 's/HOSTIP/$HOSTIP/' src/go-web/unit_test.go ; go test ${APP_NAME} -v --run Unit"

    } catch(e) {

      error "Test failed"

    } finally {

      // test done, clean up
      sh "/usr/local/bin/docker rm -f ${APP_NAME} || true"
      sh "/usr/local/bin/docker rmi $(/usr/local/bin/docker images -a  | grep go-web | awk '{print $3}')

      //sh "/usr/local/bin/docker ps -aq | xargs /usr/local/bin/docker rm || true"
      //sh "/usr/local/bin/docker images -aq -f dangling=true | xargs /usr/local/bin/docker rmi || true"
    }
  }
}

