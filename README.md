# URL Shortener
This is a URL shortener built with the system design discussed in [this video](https://youtu.be/Cg3XIqs_-4c?si=0HU9Zkl6vRVtIF2y) in mind. **The main purpose** of this project is for me to practice deploying a simple project in GCP.

## The plan
My end goal for this project's deployment is as follows:

1) Have 3 deployed environments: *development*, *staging*, and *produdction*. While it would be interesting to extend this with QA & Load environments, I'd like to keep the scope down to a simple practice with GCP for now.

2) Have all pre-prod deployments accessible via [gyani.dev](https://gyani.dev) sub-domains.

3) Have the prod deployment accessible via [gyani.dev/url](http://gyani.dev).

## The project
Since the project focus is on GCP, I'm not spending time on the system design portion, the linked video above seems to be pretty thorough. Honestly, this might be biting off more than I can chew with the load balancer & building for large traffic, but it'll be a fun challenge anyways.

With that said, the rest of this document is just going to be various notes on the project design. All sourced from [this video](https://youtu.be/Cg3XIqs_-4c?si=nIQXJT7qVhyWZoxi).

### Requirements
| Functional Requirements | Non-Functional Requirements |
|-------------------------|-----------------------------|
| Given a long url, create a short url.  | Very low latency |
| Given a short url, redirect to a long url. | Very high availability |

### API Design
We'll use a REST API since the data is expected in the same format everytime, GraphQL would be overkill here.

1. *POST*: `/create-url`
    - Parameters: `long-url`
    - Status Code: `201 Created`
2. *GET*: `/{short-url}`
    - Status Code: `301 Permanent Redirect`

### Schema
Okay, here we're defining the database schema. It's pretty straightforward.

| Column | Type|
---------|------
| `long-url ` | string |
| `short-url` | string |
| `created-at` | timestamp |