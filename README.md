## aws_batch_exporter
AWS Batch metrics exporter.

### How to run
#### local
```bash
go build -o aws_batch_exporter cmd/aws_batch_exporter.go
# required
export REGION= #your region
./aws_batch_exporter
```

#### docker
```bash
docker build . -t aws_batch_exporter
docker run -d -p 8080:8080 -e REGION=<your region> aws_batch_exporter
```

### Requirements
Please set the access key and secret key of AWS or authentication by IAM role in advance.

### Environment Variables
|variable|description|default|required|
|---|---|---|---|
|SERVER_ADDR|The domain and port that runs the application|`:8080`|`false`|
|REGION|Target's AWS region|`none`|`true`|

### Metrics
|metric|labels|
|---|---|
|aws_batch_submitted_job|`region`, `id`, `queue`, `name`|
|aws_batch_pending_job|`region`, `id`, `queue`, `name`|
|aws_batch_runnable_job|`region`, `id`, `queue`, `name`|
|aws_batch_starting_job|`region`, `id`, `queue`, `name`|
|aws_batch_running_job|`region`, `id`, `queue`, `name`|
|aws_batch_failed_job|`region`, `id`, `queue`, `name`|
|aws_batch_succeeded_job|`region`, `id`, `queue`, `name`|