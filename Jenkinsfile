@Library(['nyotaPipeline@master']) _  

// --------------------------------------------
// see Nyota/pipeline/README.md file for all 
// options used in mysettings
// --------------------------------------------

def mysettings = [
  deploy: [
    [name: "nyota-rest-go-template" ]
  ],
  tagversion: "${env.BUILD_ID}",
  chart: "deployments/helm-chart",
  kubeyaml: "deployments/staging/helm-chart.yaml",
  kubeverify: "rest-go-template",
  noCleanWs: 1,
]

nyotaPipeline( mysettings )
