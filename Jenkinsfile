@Library(['srePipeline']) _  

// --------------------------------------------
// see Nyota/pipeline/README.md file for all 
// options used in mysettings
// --------------------------------------------

def mysettings = [
  deploy: [
    [name: "go-template" ]
  ],
  tagversion: "${env.BUILD_ID}",
  chart: "deployment/helm-chart",
  pipelineLibraryBranch: 'master',
  microK8sDeploymentsBranch: 'master',
  kubeyaml: "deployment/staging/go-template-deploy.yaml",
  kubeverify: "go-template",
  noCleanWs: 1,
  // not yet experimental: 1,
  goldenPromote: 1,
  sa: [
    [lang: "go", find: "*.go"]
  ],
  executeCC: 1,
  stricterCCThreshold: 90.0,
  runPreE2EonMaster: 1
]

nyotaPipeline( mysettings )
