@Library(['nyotaPipeline@master']) _  

// --------------------------------------------
// see Nyota/pipeline/README.md file for all 
// options used in mysettings
// --------------------------------------------

def mysettings = [
  deploy: [
    [name: "nyota-gotemplate" ]
  ],
  tagversion: "${env.BUILD_ID}",
  chart: "deployments",
  kubeyaml: "staging/go-deploythis.yaml",
  kubeverify: "gotemplate",
]

nyotaPipeline( mysettings )
