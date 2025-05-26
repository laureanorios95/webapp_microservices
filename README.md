# WebApp Microservices

A modern, scalable web application framework built with microservices architecture.

## Overview

WebApp Microservices provides a comprehensive starter template for building distributed web applications using a microservices approach. This architecture enables independent development, deployment, and scaling of various application components while maintaining a cohesive overall system.

## Architecture

The application consists of the following microservices:

- **API Gateway**: Routes requests to appropriate services, handles authentication and API composition
- **User Service**: Manages user accounts, authentication, and profiles
- **Content Service**: Handles content creation, storage, and retrieval
- **Notification Service**: Manages email, push, and in-app notifications
- **Analytics Service**: Collects and processes application metrics and user behavior
- **Payment Service**: Processes payments and manages subscriptions

```
                                 ┌─────────────────┐
                                 │                 │
                                 │  API Gateway    │
                                 │                 │
                                 └────────┬────────┘
                                          │
                 ┌──────────┬─────────────┼────────────┬─────────┐
                 │          │             │            │         │
        ┌────────▼───┐ ┌────▼─────┐ ┌─────▼───────┐ ┌──▼──────┐ ┌▼────────┐
        │            │ │          │ │             │ │         │ │          │
        │ User       │ │ Content  │ │ Notification│ │ Payment │ │ Analytics│
        │ Service    │ │ Service  │ │ Service     │ │ Service │ │ Service  │
        │            │ │          │ │             │ │         │ │          │
        └────────────┘ └──────────┘ └─────────────┘ └─────────┘ └──────────┘
```

## Technology Stack

- **Backend Services**: Go (primary), Node.js (selected services)
- **API Gateway**: Envoy/Traefik
- **Frontend**: React with TypeScript
- **Message Broker**: RabbitMQ/Kafka
- **Databases**:
  - PostgreSQL for relational data
  - MongoDB for content and analytics
  - Redis for caching and session management
- **Deployment**: Docker, Kubernetes
- **Service Mesh**: Istio
- **Monitoring**: Prometheus, Grafana
- **Logging**: ELK Stack (Elasticsearch, Logstash, Kibana)

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Kubernetes (for production deployment)
- Go 1.16+
- Node.js 14+

### Local Development Setup

1. Clone the repository:

```bash
git clone https://github.com/laureanorios95/webapp_microservices.git
cd webapp_microservices
```

2. Start the development environment:

```bash
make dev-up
```

This will start all microservices in development mode with hot reloading enabled.

3. Access the application:

- Web UI: http://localhost:3000
- API Gateway: http://localhost:8080
- Swagger Documentation: http://localhost:8080/swagger

### Running Tests

```bash
# Run all tests
make test

# Run tests for a specific service
make test-user-service
make test-content-service
# etc...
```

## Service Details

### API Gateway

- Routes requests to appropriate services
- Handles authentication and authorization
- Implements rate limiting and request throttling
- Provides API composition for frontend clients

### User Service

- User registration and authentication
- Profile management
- Role-based access control
- OAuth integration

### Content Service

- Content CRUD operations
- Media upload and processing
- Content search and filtering
- Caching frequently accessed content

### Notification Service

- Email notifications
- Push notifications
- In-app notifications
- Notification preferences

### Payment Service

- Payment processing
- Subscription management
- Invoicing
- Payment history

### Analytics Service

- User behavior tracking
- Performance metrics
- Business insights
- Reporting

## Deployment

### Development Environment

```bash
make dev-up
```

### Staging Environment

```bash
make staging-deploy
```

### Production Environment

```bash
make prod-deploy
```

## Kubernetes Deployment

The repository includes Kubernetes manifests for deploying the entire application stack:

```bash
# Apply all manifests
kubectl apply -f k8s/

# Apply specific service manifests
kubectl apply -f k8s/user-service/
```

## Configuration

Each service can be configured via environment variables or configuration files. See the README in each service directory for specific configuration options.

Common configuration variables:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
LOG_LEVEL=info
KAFKA_BROKERS=kafka:9092
```

## API Documentation

API documentation is available via Swagger UI when the application is running:

- http://localhost:8080/swagger

## Monitoring and Logging

- Prometheus metrics: http://localhost:9090
- Grafana dashboards: http://localhost:3000
- Kibana logs interface: http://localhost:5601

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Development Workflow

1. Create a new branch for your feature or bugfix
2. Make your changes
3. Write or update tests
4. Run tests and ensure they pass
5. Submit a pull request

## Project Structure

```
webapp_microservices/
├── api-gateway/
├── user-service/
├── content-service/
├── notification-service/
├── payment-service/
├── analytics-service/
├── frontend/
├── k8s/             # Kubernetes manifests
├── docker/          # Docker-related files
├── scripts/         # Utility scripts
├── docs/            # Documentation
└── README.md
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

Unknown
