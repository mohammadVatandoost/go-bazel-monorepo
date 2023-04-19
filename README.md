### Step 1: Update BUILD Files

Run the complete build pipeline using:

```sh
bazel run //:gazelle
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies 
bazel run //:gazelle -- update
```

### Step 2: Build it

Run the complete build pipeline using:

```sh
bazel build //...
```

### Step 3: Test it

Run the Test

```sh
bazel test //...
```

### Step 3: Run it

Run the services

```sh
 bazel run //services/servicea
 bazel run //services/serviceb
```
