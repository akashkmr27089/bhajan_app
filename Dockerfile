# Use the official MongoDB image from Docker Hub
FROM mongo:latest

# Set environment variables
ENV MONGO_INITDB_ROOT_USERNAME=admin
ENV MONGO_INITDB_ROOT_PASSWORD=admin123
ENV MONGO_INITDB_DATABASE=mydatabase

# Expose the default MongoDB port
EXPOSE 27017

# Copy any initialization scripts (optional)
# ADD ./init-db.d /docker-entrypoint-initdb.d

# The command to run MongoDB
CMD ["mongod"]
