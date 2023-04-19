### Step 1: Update BUILD Files

Run the complete build pipeline using:

```sh
bazel run //:gazelle
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies 
```

### Step 2: Build it

Run the complete build pipeline using:

```sh
bazel build //...
```

### Step 2: Test it

Run the Test

```sh
bazel test //...
```

