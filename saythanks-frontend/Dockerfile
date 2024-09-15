# syntax=docker/dockerfile:1.3
# Use the Node.js 20 slim image
FROM node:20-slim

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json files
COPY package*.json ./

# Install dependencies with cache
RUN --mount=type=cache,target=/root/.npm \
    npm install

# Copy the rest of the application code
COPY . .

# Build the Next.js application with cache
RUN --mount=type=cache,target=/root/.npm \
    npm run build

# Expose the port the app runs on
EXPOSE 3000

# Command to run the Next.js application
CMD ["npm", "start"]