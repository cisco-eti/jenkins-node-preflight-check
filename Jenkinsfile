@Library(['nyotaPipeline@master']) _  

// --------------------------------------------
// see Nyota/pipeline/README.md file for all 
// options used in mysettings
// --------------------------------------------

def mysettings = [
  deploy: [
    [name: "go-template" ]
  ],
  tagversion: "${env.BUILD_ID}",
  chart: "deployments/helm-chart",
  kubeyaml: "deployments/staging/helm-chart.yaml",
  kubeverify: "go-template",
  noCleanWs: 1,
  experimental: 1,
  goldenPromote: 1,
  sa: [
    [lang: "go", find: "*.go"]
  ],
  //executeCC: 1,
]

nyotaPipeline( mysettings )
