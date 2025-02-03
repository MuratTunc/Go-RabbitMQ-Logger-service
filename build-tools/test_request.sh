#!/bin/bash

# Set the logger service URL
LOGGER_URL="http://localhost:8081/log"
GET_LOGS_URL="http://localhost:8081/logs"
SERVICE_NAME="LoggerService"
LOG_MESSAGE="Test log message"

# Step 1: Send a POST request to the logger service
echo "Sending POST request to logger service to insert log..."
echo "Request Body: {\"service_name\": \"$SERVICE_NAME\", \"log_message\": \"$LOG_MESSAGE\"}"  # Use double quotes
response=$(curl -s -X POST $LOGGER_URL -H "Content-Type: application/json" -d "{\"service_name\": \"$SERVICE_NAME\", \"log_message\": \"$LOG_MESSAGE\"}")

# Check if the POST request was successful
if [[ "$response" == "Log saved successfully" ]]; then
  echo "Log entry successfully created!"
else
  echo "Failed to create log entry: $response"
  exit 1
fi

# Step 2: Wait a few seconds to let the log be inserted into MongoDB
echo "Waiting for a few seconds to ensure log is inserted into the database..."
sleep 5

# Step 3: Send a GET request to check the inserted log
echo "Sending GET request to fetch logs for service '$SERVICE_NAME'..."
logs=$(curl -s "$GET_LOGS_URL?service_name=$SERVICE_NAME")

# Check if the response contains the expected log message
if echo "$logs" | grep -q "$LOG_MESSAGE"; then
  echo "Log entry found successfully!"
else
  echo "Log entry not found. Response from server:"
  echo "$logs"
  exit 1
fi

echo "Test completed successfully!"
