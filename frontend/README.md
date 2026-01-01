# Zarish HIS UI

This is the standalone frontend for the Zarish Hospital Information System (HIS) module.

## Development

1. Install dependencies:
   ```bash
   npm install
   ```

2. Start the development server:
   ```bash
   npm run dev
   ```
   The app will run on `http://localhost:5173` by default.

## Configuration

- **API URL**: The app expects the backend API URL to be set in `VITE_API_URL`.
  - Default: `http://localhost:8083/api/v1`
  - Create a `.env` file to override:
    ```
    VITE_API_URL=http://localhost:8083/api/v1
    ```

## Deployment

### Docker

Build and run using Docker:

```bash
docker build -t zarish-his-ui .
docker run -p 8084:80 zarish-his-ui
```

### Docker Compose

This service is included in the platform `docker-compose.yml` as `zarish-his-ui`.

```bash
docker compose up -d zarish-his-ui
```
