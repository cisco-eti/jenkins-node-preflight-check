@Library(['srePipeline@master']) _

// --------------------------------------------
// see Nyota/pipeline/README.md file for all
// options used in pipelinesettings
// --------------------------------------------

def pipelinesettings = [
  deploy: [
    [name: "sre-go-helloworld" ]
  ],
  tagversion: "${env.BUILD_ID}",
  chart: "deployment/helm-chart",
  kubeverify: "sre-go-helloworld",
  namespace: 'helloworld',

  // Knobs to turn on pipeline stages
  prepare: 1,
  unittest: 1,
  build: 1,
  // TODO: Disable and fix after break
  // blackduck:1,
  // sonarQube: [[ name: "sonar-sjc" ]],
  // sa: [[lang: "go", find: "*.go"]],
  helm: 1,
  // preDeployE2E: 1,
  deploy: 1,
  e2e: 1,
  apiDocs: 1,

  // use artifactory credentials for go modules
  artifactory: 1,

  // Push to ECR public repos on tags
  pushPublicRegistryOnTag: 1,

  // Code coverage threshold to fail builds
  stricterCCThreshold: 90.0,
  runPreE2EonMaster: 1
]

srePipeline( pipelinesettings )
