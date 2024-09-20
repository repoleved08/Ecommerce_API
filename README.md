# Techxtrasol Commerce
Techxtrasol Commerce is an api created from Golang and Echo framework. It is a simple api that allows you to create, read, update and delete products. It also allows you to upload images for the products. It also allows the users to add items to cart and make payment. The api is created using the following technologies:
- Golang (Programming Language) 
- Echo (Web Framework)
- Gorm (ORM) 
- Postgres (Database)
- Docker (Containerization)
- AWS S3 (Storage)
- Stripe (Payment Gateway)
- mailgun (Email Service)
- Mpesa (Payment Gateway)

## Installation
To install the api, you need to have the following installed:
- Golang (https://golang.org/) 
- Docker (https://www.docker.com/)
- Postgres (https://www.postgresql.org/)
- AWS S3 (https://aws.amazon.com/s3/)
- Stripe (https://stripe.com/)
- mailgun (https://www.mailgun.com/)
- Mpesa (https://developer.safaricom.co.ke/)
- Postman (https://www.postman.com/)
- Git (https://git-scm.com/)
- Any code editor of your choice

## Usage
To use the api, you need to follow the following steps:
- Clone the repository
- Create a .env file in the root directory and add the following environment variables:
```
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=techxtrasol_commerce
AWS_ACCESS_KEY_ID=your_aws_access_key_id
AWS_SECRET_ACCESS_KEY=your_aws_secret_access_key
AWS_REGION=your_aws_region
AWS_BUCKET_NAME=your_aws_bucket_name
STRIPE_SECRET_KEY=your_stripe_secret_key
MAILGUN_API_KEY=your_mailgun_api_key
MAILGUN_DOMAIN=your_mailgun_domain
MPESA_CONSUMER_KEY=your_mpesa_consumer_key
MPESA_CONSUMER_SECRET=your_mpesa_consumer_secret
MPESA_SHORTCODE=your_mpesa_shortcode
MPESA_PASSKEY=your_mpesa_passkey
```
- Run the following command to start the api:
```
go run main.go
```
- Open Postman and test the api

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

