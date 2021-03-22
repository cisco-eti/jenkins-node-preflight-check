@Library(['srePipeline']) _

// --------------------------------------------
// Refer to Pipeline docs for options used in mysettings
// https://wwwin-github.cisco.com/pages/eti/sre-pipeline-library
// --------------------------------------------

// TODO: Enable after testing
// blackduck: 1
//  apiDocs: 1,
// sonarQube: [ [ name: "sonar-sjc" ] ],
// sa: [ [lang: "go", find: "*.go"] ],

def pipelinesettings = [
  deploy: [
    [name: "sre-go-helloworld" ]    // Containers to publish
  ],

  tagversion: "${env.BUILD_ID}",    // Docker tag version
  chart: "deployment/helm-chart",   // Location of helm chart
  kubeverify: "sre-go-helloworld",  // Deploy verification name
  namespace: 'helloworld',          // k8s namespace

  prepare: 1,                       // GIT Clone
  unittest: 1,                      // Unit-test
  build: 1,                         // Build container
  executeCC: 1,                     // Generate Code Coverage report
  lint: 1,                          // GO Lint
  sonarQube: 1,                     // SonarQube scan
  publishContainer: 1,              // Publish container
  ecr: 1,                           // Publish container to Private ECR
  ciscoContainer: 1,                // Publish container to containers.cisco.com
  dockerHub: 1,                     // Publish container to dockerhub.cisco.com
  pushPublicRegistryOnTag: 1,       // Publish container to Public ECR on tag
  blackduck: [
    imgId: "73243",
    coronaRelease: "1.0.0",
    email: "eti-sre-admins@cisco.com",
  ],                                // Blackduck Open Source Scan
  // forceBlackduck: 1,             // Force Blackduck Scan on any branch
  publishHelm: 1,                   // Stage HELM CREATE
  deployHelm: 1,                    // Stage DEPLOY k8s
  artifactory: 1,                   // Use Artifactory creds
  stricterCCThreshold: 90.0,        // Fail builds for Code Coverage below 90%
]

srePipeline( pipelinesettings )

