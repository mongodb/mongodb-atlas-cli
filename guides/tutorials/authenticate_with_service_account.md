# Authenticate Using a Service Account in Atlas CLI

This tutorial provides step-by-step instructions for authenticating your MongoDB Atlas CLI profile using Service Accounts, allowing you to automate and and manage MongoDB resources. Service Accounts, also referred to as OAuth Applications, allows programmatic access through a secure Client ID and Secret—ideal for scripting use cases and CI/CD workflows.

__Note:__ This process is interactive. For programmatic authentication, set enviroment variables `MONGODB_ATLAS_CLIENT_ID` and `MONGODB_ATLAS_CLIENT_SECRET` before using the AtlasCLI. See [Atlas CLI Environment Variables](https://www.mongodb.com/docs/atlas/cli/stable/atlas-cli-env-variables/) for all supported environment variables.

---

### Prerequisites
Before starting, ensure you have:

* MongoDB Atlas CLI installed.
* An active MongoDB Atlas Organization.
* A Service Account created and configured. See [Grant Programmatic Access to an Organization](https://www.mongodb.com/docs/atlas/configure-api-access/#grant-programmatic-access-to-an-organization) for further details.
* The Client ID and Client Secret of your Service Account.

---

### Step 1: Start the Authentication Process
Open your terminal and run the following command:

```
atlas auth login --profile myProfile
```

This command initiates the interactive authentication setup guide for a Atlas CLI profile called 'myProfile'.

---
### Step 2: Select the Authentication Method
After running `atlas auth login --profile myProfile`, you will see the following prompt:

```
? Select authentication type:  [Use arrows to move, type to filter]  
  UserAccount - (best for getting started)  
> ServiceAccount - (best for automation)  
  APIKeys - (for existing automations)  
```

Use the arrow keys to select ServiceAccount and press **Enter**.

---
### Step 3: Enter Service Account Credentials
Once you've selected ServiceAccount, the CLI will prompt you to provide the Client ID and Client Secret of the Service Account you created:

```
? Select authentication type: ServiceAccount
You are configuring a profile for atlas.

All values are optional and you can use environment variables (MONGODB_ATLAS_*) instead.

Enter [?] on any option to get help.

? Client ID: <your-service-account-client-id>  
? Client Secret: <your-service-account-client-secret>  
```

Input these credentials accurately and press **Enter**.

---
### Step 4: Select Default Project and Output Format
If your Service Account is at organization level, the CLI will prompt you to select a default project. This will be the project to which commands will be executed by default.

Next, the CLI will prompt you to select the output format.

```
? Choose a default project:  [Use arrows to move, type to filter]
> a0123456789abcdef012345a - myProject
  b0123456789abcdef012345b - myOtherProject

? Default Output Format:  [Use arrows to move, type to filter]
> plaintext
  json
```

---
### Step 5: Authentication Complete
At this point, your Atlas CLI profile is authenticated using your Service Account. You can now execute authenticated commands in Atlas CLI within the scope of your Service Account permissions by adding `--profile myProfile` flag to any command you run.

For example, `atlas cluster list --profile myProfile`
