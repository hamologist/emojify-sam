# Emojify Sam
**Emojify Sam** is an AWS SAM project that exposes emojification from [Emojify](https://github.com/hamologist/emojify) on AWS Lambdas.
The project currently initiates two seperate lambdas,
one for emojifying via an API endpoint,
and another that emojifies Discord messages via slash commands.

## Installation
Users interested in running their own lambdas can use SAM to get their infrastructure online.
Assuming SAM has been installed and their AWS account has been connected to their local setup,
users can use the guided setup via:
```bash
$ sam build
$ sam deploy --guided
```
**Note:** Running the above will require values for the following template parameters:
* **Hosted Zone Name**
    * The Route 53 Hosted zone domain name that should be used for resolving traffic to the lambdas spun up.
    * Example: example.com.
    * **Note:** Make sure to include the trailing "."
* **Domain Name**
    * The subdomain to use on the Hosted zone provided above.
    * Example: api.example.com
* **Cerificate ARN**
    * The ARN for a valid AWS certificate from the "AWS Certicate Manager".
    * Example: arn:aws:acm:us-east-1:...:certificate/...
    * **Note:** Certificate should match value provided for "Domain Name" above.
* **Enable Discord Resources**
    * String boolean (valid values are "true" or "false") that determines if Emojify Discord function resources are deployed.
* **Discord Public Key**
    * The public key for the bot that will hit the Emojify Discord endpoint.
        * Value can be left as default if Discord resources are disabled.

## Usage
Once the deploy completes the following endpoints should be available on the domain name provided during installation:
* {provided-domain-name}.{hosted-zone-name}/emojify
    * POST endpoint
    * Expects request using the following payload structure:
        ```json
        {
            "message": "This is emojification on AWS using SAM!"
        }
        ```
    * Curl request:
        ```bash
        curl --location --request POST 'https://{provided-domain-name}.{hosted-zone-name}/emojify' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "message": "This is emojification on AWS using SAM!"
        }'
        ```
* {provided-domain-name}.{hosted-zone-name}/discord
    * Available assuming Discord resources were enabled for deploy.
    * POST endpoint
    * Expects request following Discord's interaction payload structure:
        ```json
        {
            "id":"{id-value}",
            "token":"{token-value}",
            "type":{type-value},
            "user":{
                "avatar":"{avatar-value}",
                "discriminator":"{discriminator-value}",
                "id":"{user-id-value}",
                "public_flags":{public_flags-value},
                "username":"{username-value}"
            },
            "version":{version-value}
        }
        ```
        Check [here](https://discord.com/developers/docs/interactions/slash-commands) for more information.
* **Note:**
    * Current template.yaml for the project will ensure CORS are globally enabled for all endpoints.

## Public Resource
The `emojify` resource is available to everyone via https://emojify.hamologist.com/emojify.
Users can see the endpoint in action using the following:
```bash
curl --location --request POST 'https://emojify.hamologist.com/emojify' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "Public emojification resources? Sign me up!"
}'
```

## Related
* [Emojify](https://github.com/hamologist/emojify)
    * The library that handles application logic.
* [Meme Machine](https://github.com/hamologist/meme-machine)
    * Python discord bot that supports emojifying messages via `!emojify` [command](https://github.com/hamologist/meme-machine/blob/master/meme_machine/emojify.py)