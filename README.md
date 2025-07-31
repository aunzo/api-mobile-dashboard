# API Mobile Dashboard

## Overview
This is a Go application that provides an API for a mobile dashboard. It uses Firebase Firestore for data storage and Gin as the web framework.

## Deployment

### Frontend (Firebase Hosting)
The frontend has been deployed to Firebase Hosting and is accessible at: https://mobile-dashboard-69c8b.web.app

### Backend (Cloud Run)
To deploy the backend to Cloud Run, follow these steps:

1. Enable the Cloud Run Admin API for your project:
   - Visit: https://console.developers.google.com/apis/api/run.googleapis.com/overview?project=163142454702
   - Click "Enable"

2. Build and deploy the Docker container to Cloud Run:
   ```bash
   # Build the container
   docker build -t gcr.io/mobile-dashboard-69c8b/api-mobile-dashboard .
   
   # Configure Docker to use Google Cloud credentials
   gcloud auth configure-docker
   
   # Push the container to Google Container Registry
   docker push gcr.io/mobile-dashboard-69c8b/api-mobile-dashboard
   
   # Deploy to Cloud Run
   gcloud run deploy api-mobile-dashboard \
     --image gcr.io/mobile-dashboard-69c8b/api-mobile-dashboard \
     --platform managed \
     --region us-central1 \
     --allow-unauthenticated
   ```

3. Update Firebase Hosting configuration:
   After deploying to Cloud Run, update the `firebase.json` file to include the Cloud Run service:
   ```json
   {
     "hosting": {
       "public": "public",
       "ignore": [
         "firebase.json",
         "**/.*",
         "**/node_modules/**"
       ],
       "rewrites": [
         {
           "source": "/build-info/**",
           "run": {
             "serviceId": "api-mobile-dashboard",
             "region": "us-central1"
           }
         },
         {
           "source": "/swagger/**",
           "run": {
             "serviceId": "api-mobile-dashboard",
             "region": "us-central1"
           }
         },
         {
           "source": "**",
           "destination": "/index.html"
         }
       ]
     }
   }
   ```

4. Deploy the updated Firebase configuration:
   ```bash
   firebase deploy
   ```

## Local Development

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

3. The server will start at http://localhost:3000

## API Documentation
Swagger documentation is available at `/swagger/index.html` when the application is running.