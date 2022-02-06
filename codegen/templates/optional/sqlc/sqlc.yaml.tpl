version: 1
packages:
  - path: "sqlc"          # name of the dir / package that holds the generated code
    schema: "./sql/"      # path to the the dir that holds all of the tables
    queries: "./sql/"     # path to the the dir that holds all of the queries 
    engine: "postgresql"  # db type that will be used
    emit_json_tags: true