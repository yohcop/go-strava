This was generated with the Swagger jar file from https://github.com/swagger-api/swagger-codegen, with:

    java -jar ~/Downloads/swagger-codegen-cli-2.3.1.jar generate -i https://developers.strava.com/swagger/swagger.json -l go --output .


Then, I replaced all the imports for "golang.org/x/net/context" by "context".

I also generated the Bazel BUILD files with Gazelle.
