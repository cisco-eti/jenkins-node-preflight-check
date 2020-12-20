@Library(['srePipeline@master']) _

// --------------------------------------------
// see Nyota/pipeline/README.md file for all 
// options used in mysettings
// --------------------------------------------

def mysettings = [
  deploy: [
    [name: "sre-go-helloworld" ]
  ],
  tagversion: "${env.BUILD_ID}",
  chart: "deployment/helm-chart",
  pipelineLibraryBranch: 'master',
  microK8sDeploymentsBranch: 'master',
  kubeyaml: "deployment/staging/sre-go-helloworld-deploy.yaml",
  kubeverify: "sre-go-helloworld",
  artifactory: 1,
  noPreE2E: 1,
  noPII: 1,
  noE2E: 1,
  // not yet experimental: 1,
  //goldenPromote: 1,

  //   sa: [
  //     [lang: "go", find: "*.go"]
  //   ],
  // executeCC: 1,
  namespace: 'helloworld',
  stricterCCThreshold: 90.0,
  runPreE2EonMaster: 1
]

srePipeline( mysettings )
