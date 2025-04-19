
# Security Setup for Local Development and AWS Remote Environment

## Table of Contents
- [Local Development Setup](#local-development-setup)
- [AWS Remote Environment Setup](#aws-remote-environment-setup)
- [Storing Other Sensitive Data (e.g., S3 Bucket, Account ID)](#storing-other-sensitive-data-eg-s3-bucket-account-id)
- [Best Practices](#best-practices)

---

## Local Development Setup

### 1. **Use AWS CLI Profiles**
- Set up separate AWS CLI profiles for **local development**.
- Run `aws configure --profile <profile-name>` to configure your profile with **IAM user credentials** that have **limited permissions** (e.g., only access to necessary resources like S3 for testing).

Example:
   ```bash
   aws configure --profile dev-profile
   ```

### 2. **Use AWS IAM with Least Privilege**
- Create an IAM user with only the **necessary permissions** for local dev (e.g., access to S3 buckets, Lambda functions).
- Assign this IAM user to a specific **dev role**.
- Avoid using the **root account** for any development activity.

### 3. **Secure Credentials**
- Credentials are stored in `~/.aws/credentials` and are associated with your profile.
- Ensure that this file is **not committed** to version control (add it to `.gitignore`).
- Optionally, use **MFA (Multi-Factor Authentication)** for added security.

### 4. **Environment Variables for Development**
- Store other sensitive data such as API keys, S3 bucket name, AWS region, and AWS account ID using **environment variables** in your local dev environment.
- You can define them in `.env` files (excluded from version control) or export them manually in your shell.

Example `.env` (do not commit):
   ```dotenv
   AWS_REGION=us-east-1
   S3_BUCKET_NAME=my-dev-bucket
   AWS_ACCOUNT_ID=123456789012
   ```

---

## AWS Remote Environment Setup

### 1. **Use GitHub Actions for CI/CD**
- Use **GitHub Secrets** to securely store sensitive data such as:
    - `AWS_ACCESS_KEY_ID`
    - `AWS_SECRET_ACCESS_KEY`
    - `AWS_ACCOUNT_ID`
    - `S3_BUCKET_NAME`
    - `AWS_REGION`

Example usage in your GitHub Actions workflow:
   ```yaml
   env:
     AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
     AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
     AWS_ACCOUNT_ID: ${{ secrets.AWS_ACCOUNT_ID }}
     S3_BUCKET_NAME: ${{ secrets.S3_BUCKET_NAME }}
     AWS_REGION: ${{ secrets.AWS_REGION }}
   ```

### 2. **IAM Role for Remote Environment**
- In your **AWS account**, create an IAM **role** for your GitHub Actions runner with only the permissions required for deployment.
- Never use the root account or store its credentials.
- Use **assume role** setups for scoped, temporary access if needed.

### 3. **Secure AWS CLI in GitHub Actions**
- Use the [aws-actions/configure-aws-credentials](https://github.com/aws-actions/configure-aws-credentials) action to set up temporary credentials from GitHub Secrets.

---

## Storing Other Sensitive Data (e.g., S3 Bucket, Account ID)

Besides credentials, you’ll also need to securely store environment-specific values like:

- `AWS_ACCOUNT_ID`
- `S3_BUCKET_NAME`
- `AWS_REGION`
- `LOG_LEVEL`, `MAX_IMAGE_SIZE`, `SERVER_PORT` (if sensitive or variable)

### 🔒 Local Development
Store them in:
- Local environment variables
- A `.env` file excluded via `.gitignore`
- Your shell profile (e.g., `.bashrc`, `.zshrc`)

Example `.env` file:
```dotenv
AWS_ACCOUNT_ID=123456789012
S3_BUCKET_NAME=my-bucket
AWS_REGION=us-east-1
```

Load with:
```bash
source .env
```

### 🔐 GitHub Actions
Use GitHub Secrets:
1. Go to **Repo → Settings → Secrets and Variables → Actions**.
2. Click **New Repository Secret** for each key-value pair:
    - `AWS_ACCOUNT_ID`
    - `S3_BUCKET_NAME`
    - `AWS_REGION`

Then reference them in your workflow file:
```yaml
env:
  AWS_ACCOUNT_ID: ${{ secrets.AWS_ACCOUNT_ID }}
  S3_BUCKET_NAME: ${{ secrets.S3_BUCKET_NAME }}
  AWS_REGION: ${{ secrets.AWS_REGION }}
```

---

## Best Practices

### ✅ IAM with Least Privilege
- Assign only the permissions required.
- Use roles for service-to-service interactions.

### 🔁 Rotate Credentials Regularly
- Rotate IAM user access keys every 90 days or sooner.
- Update GitHub Secrets accordingly.

### 🔐 Use MFA
- Enable MFA for IAM users with console access.

### ❌ Never Hardcode Credentials
- Use environment variables and GitHub Secrets.
- Exclude `.env`, `.aws/credentials`, etc., from version control.

### 📜 Audit and Monitor Access
- Enable CloudTrail and monitor IAM usage.
- Use Access Analyzer to catch overly permissive policies.

### 🔐 Secure S3 Buckets
- Apply least privilege bucket policies.
- Enable server-side encryption and access logging.

---

## Summary

| Location         | Store AWS Keys | Store Other Sensitive Values | Notes                            |
|------------------|----------------|-------------------------------|----------------------------------|
| Local Dev        | AWS CLI Profile (`~/.aws/credentials`) | `.env` file or shell env vars   | Keep out of version control      |
| GitHub Actions   | GitHub Secrets | GitHub Secrets                | Secure and encrypted, CI/CD safe |

This setup provides a secure, cost-effective way to manage credentials and sensitive values for both local and cloud environments without relying on paid services like AWS Secrets Manager.
