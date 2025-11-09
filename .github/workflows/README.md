# GitHub Actions Workflows

## CI/CD Pipeline

The `ci-cd.yml` workflow handles continuous integration and deployment.

### Workflow Triggers

- **Pull Request**: Runs lint and tests only
- **Push to main**: Full pipeline with deployment to AWS
- **Manual**: Can be triggered via GitHub Actions UI

### Jobs Overview

#### 1. `backend-lint-test`
- Runs golangci-lint
- Executes Go tests with race detection
- Uploads coverage to Codecov

#### 2. `frontend-lint-check`
- Runs ESLint
- Runs Svelte type checking
- Validates TypeScript compilation

#### 3. `build-backend`
- Compiles Go binary for Lambda (ARM64)
- Creates deployment package (zip)
- Uploads artifact for deployment job

#### 4. `build-frontend`
- Installs npm dependencies
- Builds static assets with SvelteKit
- Uploads artifact for deployment job

#### 5. `deploy`
- Downloads build artifacts
- Configures AWS credentials (OIDC)
- Runs Terraform apply
- Updates Lambda function code
- Syncs frontend to S3
- Outputs deployment URLs

### Required Secrets

Add these in `Settings` → `Secrets and variables` → `Actions`:

- `AWS_ROLE_ARN`: IAM role ARN for OIDC authentication

Alternative (less secure):
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`

### Environment

The workflow uses a GitHub Environment named `demo` for:
- Deployment approvals (optional)
- Environment-specific secrets
- Deployment history

To configure: `Settings` → `Environments` → `New environment` → `demo`

### Manual Deployment

To manually trigger deployment:

1. Go to `Actions` tab
2. Select `CI/CD Pipeline` workflow
3. Click `Run workflow` → `Run workflow`

### Customization

Edit `.github/workflows/ci-cd.yml` to:
- Change Node.js/Go versions
- Adjust AWS region
- Modify build commands
- Add additional steps
